package main

import (
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"

	"github.com/r2d2-ai/aiflow/engine"
	_ "github.com/r2d2-ai/aiflow/microgateway/activity/circuitbreaker"
	_ "github.com/r2d2-ai/aiflow/microgateway/activity/jwt"
	_ "github.com/r2d2-ai/aiflow/microgateway/activity/ratelimiter"
	"github.com/r2d2-ai/aiflow/microgateway/examples"
	_ "github.com/r2d2-ai/contrib/activity/rest"
)

var (
	server = flag.Bool("server", false, "run the test server")
)

const reply = `{
	"id": 1,
	"category": {
		"id": 0,
		"name": "string"
	},
	"name": "sally",
	"photoUrls": ["string"],
	"tags": [{ "id": 0,"name": "string" }],
	"status":"available"
}`

func main() {
	flag.Parse()

	if *server {
		http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("url: %q\n", html.EscapeString(r.URL.Path))
			defer r.Body.Close()
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))
			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write([]byte(reply))
			if err != nil {
				panic(err)
			}
		})

		err := http.ListenAndServe(":1234", nil)
		if err != nil {
			panic(err)
		}

		return
	}

	e, err := examples.DefaultHTTPPattern()
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
