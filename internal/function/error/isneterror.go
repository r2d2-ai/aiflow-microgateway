package error

import (
	"net"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

func init() {
	function.Register(&fnIsNetError{})
}

type fnIsNetError struct {
}

// Name returns the name of the function
func (fnIsNetError) Name() string {
	return "isneterror"
}

// Sig returns the function signature
func (fnIsNetError) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

// Eval executes the function
func (fnIsNetError) Eval(params ...interface{}) (interface{}, error) {
	_, ok := params[0].(net.Error)
	return ok, nil
}
