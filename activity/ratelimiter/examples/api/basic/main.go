package main

import (
	"github.com/r2d2-ai/aiflow/engine"
	"github.com/r2d2-ai/aiflow/microgateway/activity/ratelimiter/examples"
)

func main() {
	e, err := examples.Example("3-M", 0)
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
