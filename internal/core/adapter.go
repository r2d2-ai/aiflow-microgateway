package core

import (
	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/microgateway/api"
)

// Adapter is an adapter activity for ServiceFunc
type Adapter struct {
	Handler api.ServiceFunc
}

// Metadata returns the metadata for the adapter activity
func (a *Adapter) Metadata() *activity.Metadata {
	return nil
}

// Eval evaluates the adapter activity
func (a *Adapter) Eval(ctx activity.Context) (done bool, err error) {
	return a.Handler(ctx)
}
