package main

import (
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

var (
	server = flag.Bool("server", false, "run the test server")
)

const resource = `{
  "name": "Pets",
  "steps": [{
    "service": "PetStorePets",
    "input": {
      "pathParams": "=$.payload.pathParams"
    }
  }],
  "responses": [{
    "error": false,
    "output": {
      "code": 200,
      "data": "=$.PetStorePets.outputs.data"
    }
  }],
  "services": [{
    "name": "PetStorePets",
    "description": "Get pets by ID from the petstore",
    "ref": "github.com/r2d2-ai/contrib/activity/rest",
    "settings": {
      "uri": "http://petstore.swagger.io/v2/pet/:petId",
      "method": "GET",
      "headers": {
        "Accept": "application/json"
      }
    }
  }]
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
			_, err = w.Write([]byte(resource))
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
}
