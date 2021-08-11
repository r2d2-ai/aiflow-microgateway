package anomaly

import (
	"encoding/json"
	"math"
	"math/bits"
	"sync"

	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/data/metadata"
)

const (
	// CDF16Fixed is the shift for 16 bit coders
	CDF16Fixed = 16 - 3
	// CDF16Scale is the scale for 16 bit coder
	CDF16Scale = 1 << CDF16Fixed
	// CDF16Rate is the damping factor for 16 bit coder
	CDF16Rate = 5
	// CDF16Size is the size of the cdf
	CDF16Size = 256
	// CDF16Depth is the depth of the context tree
	CDF16Depth = 2
)

func init() {
	activity.Register(&Activity{}, New)
}

var (
	activityMetadata = activity.ToMetadata(&Settings{}, &Input{}, &Output{})
)

// Node16 is a context node
type Node16 struct {
	Model    []uint16
	Children map[uint16]*Node16
}

// NewNode16 creates a new context node
func NewNode16() *Node16 {
	model, children, sum := make([]uint16, CDF16Size+1), make(map[uint16]*Node16), 0
	for i := range model {
		model[i] = uint16(sum)
		sum += 32
	}
	return &Node16{
		Model:    model,
		Children: children,
	}
}

// CDF16 is a context based cumulative distributive function model
// https://fgiesen.wordpress.com/2015/05/26/models-for-adaptive-arithmetic-coding/
type CDF16 struct {
	Root  *Node16
	Mixin [][]uint16
}

// NewCDF16 creates a new CDF16 with a given context depth
func NewCDF16() *CDF16 {
	root, mixin := NewNode16(), make([][]uint16, CDF16Size)

	for i := range mixin {
		sum, m := 0, make([]uint16, CDF16Size+1)
		for j := range m {
			m[j] = uint16(sum)
			sum++
			if j == i {
				sum += CDF16Scale - CDF16Size
			}
		}
		mixin[i] = m
	}

	return &CDF16{
		Root:  root,
		Mixin: mixin,
	}
}

// Context16 is a 16 bit context
type Context16 struct {
	Context []uint16
	First   int
}

// NewContext16 creates a new context
func NewContext16(depth int) *Context16 {
	return &Context16{
		Context: make([]uint16, depth),
	}
}

// ResetContext resets the context
func (c *Context16) ResetContext() {
	c.First = 0
	for i := range c.Context {
		c.Context[i] = 0
	}
}

// AddContext adds a symbol to the context
func (c *Context16) AddContext(s uint16) {
	context, first := c.Context, c.First
	length := len(context)
	if length > 0 {
		context[first], c.First = s, (first+1)%length
	}
}

// Model gets the model for the current context
func (c *CDF16) Model(ctxt *Context16) []uint16 {
	context := ctxt.Context
	length := len(context)
	var lookUp func(n *Node16, current, depth int) *Node16
	lookUp = func(n *Node16, current, depth int) *Node16 {
		if depth >= length {
			return n
		}

		node := n.Children[context[current]]
		if node == nil {
			return n
		}
		child := lookUp(node, (current+1)%length, depth+1)
		if child == nil {
			return n
		}
		return child
	}

	return lookUp(c.Root, ctxt.First, 0).Model
}

// Update updates the model
func (c *CDF16) Update(s uint16, ctxt *Context16) {
	context, first, mixin := ctxt.Context, ctxt.First, c.Mixin[s]
	length := len(context)
	var update func(n *Node16, current, depth int)
	update = func(n *Node16, current, depth int) {
		model := n.Model
		size := len(model) - 1

		for i := 1; i < size; i++ {
			a, b := int(model[i]), int(mixin[i])
			model[i] = uint16(a + ((b - a) >> CDF16Rate))
		}

		if depth >= length {
			return
		}

		node := n.Children[context[current]]
		if node == nil {
			node = NewNode16()
			n.Children[context[current]] = node
		}
		update(node, (current+1)%length, depth+1)
	}

	update(c.Root, first, 0)
	ctxt.AddContext(s)
}

// Complexity is an entorpy based anomaly detector
type Complexity struct {
	*CDF16
	depth          int
	count          int
	mean, dSquared float32
	sync.RWMutex
}

// NewComplexity creates a new entorpy based model
func NewComplexity(depth int) *Complexity {
	return &Complexity{
		CDF16: NewCDF16(),
		depth: depth,
	}
}

// Complexity outputs the complexity
func (c *Complexity) Complexity(input []byte) (float32, int) {
	var total uint64
	ctxt := NewContext16(c.depth)
	c.RLock()
	for _, s := range input {
		model := c.Model(ctxt)
		total += uint64(bits.Len16(model[s+1] - model[s]))
		ctxt.AddContext(uint16(s))
	}
	c.RUnlock()

	ctxt.ResetContext()
	c.Lock()
	for _, s := range input {
		c.Update(uint16(s), ctxt)
	}

	complexity := float32(CDF16Fixed+1) - (float32(total) / float32(len(input)))
	// https://dev.to/nestedsoftware/calculating-standard-deviation-on-streaming-data-253l
	c.count++
	count := c.count
	mean, n := c.mean, float32(count)
	meanDifferential := (complexity - mean) / n
	newMean := mean + meanDifferential
	dSquaredIncrement := (complexity - newMean) * (complexity - mean)
	newDSquared := c.dSquared + dSquaredIncrement
	c.mean, c.dSquared = newMean, newDSquared
	c.Unlock()

	stddev := float32(math.Sqrt(float64(newDSquared / n)))
	normalized := (complexity - newMean) / stddev
	if normalized < 0 {
		normalized = -normalized
	}
	if math.IsNaN(float64(normalized)) {
		normalized = 0
	}

	return normalized, count
}

// Activity is an anomaly detector
type Activity struct {
	complexity *Complexity
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := Settings{
		Depth: CDF16Depth,
	}
	err := metadata.MapToStruct(ctx.Settings(), &settings, true)
	if err != nil {
		return nil, err
	}

	logger := ctx.Logger()
	logger.Debugf("Setting: %b", settings)

	act := &Activity{
		complexity: NewComplexity(settings.Depth),
	}

	return act, nil
}

// Metadata return the metadata for the activity
func (a *Activity) Metadata() *activity.Metadata {
	return activityMetadata
}

// Eval executes the activity
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	input := Input{}
	err = ctx.GetInputObject(&input)
	if err != nil {
		return false, err
	}

	data, err := json.Marshal(input.Payload)
	if err != nil {
		return
	}
	complexity, count := a.complexity.Complexity(data)

	output := Output{
		Complexity: complexity,
		Count:      count,
	}
	err = ctx.SetOutputObject(&output)
	if err != nil {
		return false, err
	}

	return true, nil
}
