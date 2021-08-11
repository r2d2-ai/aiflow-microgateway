package main

import (
	"github.com/r2d2-ai/aiflow/engine"
	"github.com/r2d2-ai/aiflow/microgateway/activity/jwt/examples"
)

func main() {
	e, err := examples.Example()
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
