package error

import (
	"errors"
	"testing"

	"github.com/r2d2-ai/aiflow/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnType_Eval(t *testing.T) {
	f := &fnType{}
	err1 := errors.New("test error")
	v, err := function.Eval(f, err1)
	assert.Nil(t, err)
	assert.Equal(t, "*errors.errorString", v)
}
