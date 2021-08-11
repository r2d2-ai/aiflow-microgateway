package ratelimiter

import (
	"testing"
	"time"

	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/mapper"
	"github.com/r2d2-ai/aiflow/data/metadata"
	logger "github.com/r2d2-ai/aiflow/support/log"
	"github.com/r2d2-ai/aiflow/support/trace"
	"github.com/stretchr/testify/assert"
)

type initContext struct {
	settings map[string]interface{}
}

func newInitContext(values map[string]interface{}) *initContext {
	if values == nil {
		values = make(map[string]interface{})
	}
	return &initContext{
		settings: values,
	}
}

func (i *initContext) Settings() map[string]interface{} {
	return i.settings
}

func (i *initContext) MapperFactory() mapper.Factory {
	return nil
}

func (i *initContext) Logger() logger.Logger {
	return logger.RootLogger()
}

type activityContext struct {
	input  map[string]interface{}
	output map[string]interface{}
}

func newActivityContext(values map[string]interface{}) *activityContext {
	if values == nil {
		values = make(map[string]interface{})
	}
	return &activityContext{
		input:  values,
		output: make(map[string]interface{}),
	}
}

func (a *activityContext) ActivityHost() activity.Host {
	return a
}

func (a *activityContext) Name() string {
	return "test"
}

func (a *activityContext) GetInput(name string) interface{} {
	return a.input[name]
}

func (a *activityContext) SetOutput(name string, value interface{}) error {
	a.output[name] = value
	return nil
}

func (a *activityContext) GetInputObject(input data.StructValue) error {
	return input.FromMap(a.input)
}

func (a *activityContext) SetOutputObject(output data.StructValue) error {
	a.output = output.ToMap()
	return nil
}

func (a *activityContext) GetSharedTempData() map[string]interface{} {
	return nil
}

func (a *activityContext) ID() string {
	return "test"
}

func (a *activityContext) IOMetadata() *metadata.IOMetadata {
	return nil
}

func (a *activityContext) Reply(replyData map[string]interface{}, err error) {

}

func (a *activityContext) Return(returnData map[string]interface{}, err error) {

}

func (a *activityContext) Scope() data.Scope {
	return nil
}

func (a *activityContext) Logger() logger.Logger {
	return logger.RootLogger()
}

func (a *activityContext) GetTracingContext() trace.TracingContext {
	return nil
}

func TestRatelimiter(t *testing.T) {
	activity, err := New(newInitContext(map[string]interface{}{
		"limit": "1-S",
	}))
	assert.Nil(t, err)

	ctx := newActivityContext(map[string]interface{}{
		"token": "abc123",
	})
	_, err = activity.Eval(ctx)
	assert.Nil(t, err)
	assert.False(t, ctx.output["limitReached"].(bool), "limit should not be reached")

	ctx = newActivityContext(map[string]interface{}{
		"token": "abc123",
	})
	_, err = activity.Eval(ctx)
	assert.Nil(t, err)
	assert.True(t, ctx.output["limitReached"].(bool), "limit should be reached")

	ctx = newActivityContext(map[string]interface{}{
		"token": "sally",
	})
	_, err = activity.Eval(ctx)
	assert.Nil(t, err)
	assert.False(t, ctx.output["limitReached"].(bool), "limit should not be reached")

	time.Sleep(time.Second)

	ctx = newActivityContext(map[string]interface{}{
		"token": "abc123",
	})
	_, err = activity.Eval(ctx)
	assert.Nil(t, err)
	assert.False(t, ctx.output["limitReached"].(bool), "limit should not be reached")
}

func TestSmartRatelimiter(t *testing.T) {
	activity, err := New(newInitContext(map[string]interface{}{
		"limit":          "1000-S",
		"spikeThreshold": "2",
	}))
	assert.Nil(t, err)

	for i := 0; i < 256; i++ {
		time.Sleep(50 * time.Millisecond)
		ctx := newActivityContext(map[string]interface{}{
			"token": "abc123",
		})
		_, err = activity.Eval(ctx)
		assert.Nil(t, err)
		assert.False(t, ctx.output["limitReached"].(bool), "limit should not be reached")
	}
	blocked, notBlocked := 0, 0
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		ctx := newActivityContext(map[string]interface{}{
			"token": "abc123",
		})
		_, err = activity.Eval(ctx)
		assert.Nil(t, err)
		if ctx.output["limitReached"].(bool) {
			blocked++
		} else {
			notBlocked++
		}
	}
	for i := 0; i < 256; i++ {
		time.Sleep(50 * time.Millisecond)
		ctx := newActivityContext(map[string]interface{}{
			"token": "abc123",
		})
		_, err = activity.Eval(ctx)
		assert.Nil(t, err)
		if ctx.output["limitReached"].(bool) {
			blocked++
		} else {
			notBlocked++
		}
	}
	assert.Condition(t, func() (success bool) {
		return blocked > 0
	}, "some requests should have been blocked")
	assert.Condition(t, func() (success bool) {
		return notBlocked > 0
	}, "some requests should not have been blocked")
}
