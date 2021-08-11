package main

import (
	"github.com/r2d2-ai/aiflow/engine"
	"github.com/r2d2-ai/aiflow/microgateway/activity/circuitbreaker/examples"
)

func main() {
	e, err := examples.Example("a", 2, 0, 0)
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
