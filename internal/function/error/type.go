package error

import (
	"reflect"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

func init() {
	function.Register(&fnType{})
}

type fnType struct {
}

// Name returns the name of the function
func (fnType) Name() string {
	return "type"
}

// Sig returns the function signature
func (fnType) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

// Eval executes the function
func (fnType) Eval(params ...interface{}) (interface{}, error) {
	return reflect.TypeOf(params[0]).String(), nil
}
