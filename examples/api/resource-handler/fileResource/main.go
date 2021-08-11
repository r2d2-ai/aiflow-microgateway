package main

import (
	"github.com/r2d2-ai/aiflow/engine"
	"github.com/r2d2-ai/aiflow/microgateway/examples"
)

func main() {
	e, err := examples.FileResourceHandlerExample("file://../../../json/resource-handler/fileResource/resource.json")
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
