package main

import (
	"flag"
	"io/ioutil"

	"github.com/r2d2-ai/aiflow/engine"
	microgateway "github.com/r2d2-ai/aiflow/microgateway"
)

var (
	input  = flag.String("input", "flogo.json", "the input file")
	output = flag.String("output", "app.go", "the output file")
)

func main() {
	flag.Parse()

	flogo, err := ioutil.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	app, err := engine.LoadAppConfig(string(flogo), false)
	if err != nil {
		panic(err)
	}
	microgateway.Generate(app, *output, "go.mod")
}
