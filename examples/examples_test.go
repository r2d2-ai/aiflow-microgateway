package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/r2d2-ai/aiflow/engine"
	_ "github.com/r2d2-ai/aiflow/microgateway/activity/circuitbreaker"
	_ "github.com/r2d2-ai/aiflow/microgateway/activity/jwt"
	_ "github.com/r2d2-ai/aiflow/microgateway/activity/ratelimiter"
	"github.com/r2d2-ai/aiflow/microgateway/api"
	test "github.com/r2d2-ai/aiflow/microgateway/internal/testing"
	_ "github.com/r2d2-ai/contrib/activity/channel"
	_ "github.com/r2d2-ai/contrib/activity/rest"
	"github.com/stretchr/testify/assert"
)

// Response is a reply form the server
type Response struct {
	Error string `json:"error"`
}

type handler struct {
	Slow bool
}

type resourceHandler struct {
	Slow bool
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if h.Slow {
		time.Sleep(10 * time.Second)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(reply))
	if err != nil {
		panic(err)
	}
}

func (h *resourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if h.Slow {
		time.Sleep(10 * time.Second)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(resource))
	if err != nil {
		panic(err)
	}
}

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

func testBasicGatewayApplication(t *testing.T, e engine.Engine) {
	defer api.ClearResources()
	test.Drain("9096")
	err := e.Start()
	assert.Nil(t, err)
	defer func() {
		err := e.Stop()
		assert.Nil(t, err)
	}()
	test.Pour("9096")

	transport := &http.Transport{
		MaxIdleConns: 1,
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{
		Transport: transport,
	}
	request := func() string {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/pets/1", nil)
		assert.Nil(t, err)
		response, err := client.Do(req)
		assert.Nil(t, err)
		body, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		response.Body.Close()
		return string(body)
	}

	body := request()
	assert.NotEqual(t, 0, string(body))
}

func TestBasicGatewayIntegrationAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway API integration test in short mode")
	}

	e, err := BasicGatewayExample()
	assert.Nil(t, err)
	testBasicGatewayApplication(t, e)
}

func TestBasicGatewayIntegrationJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway JSON integration test in short mode")
	}

	data, err := ioutil.ReadFile(filepath.FromSlash("./json/basic-gateway/flogo.json"))
	assert.Nil(t, err)
	cfg, err := engine.LoadAppConfig(string(data), false)
	assert.Nil(t, err)
	e, err := engine.New(cfg)
	assert.Nil(t, err)
	testBasicGatewayApplication(t, e)
}

func testHandlerRoutingApplication(t *testing.T, e engine.Engine) {
	defer api.ClearResources()
	test.Drain("9096")
	err := e.Start()
	assert.Nil(t, err)
	defer func() {
		err := e.Stop()
		assert.Nil(t, err)
	}()
	test.Pour("9096")

	transport := &http.Transport{
		MaxIdleConns: 1,
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{
		Transport: transport,
	}
	request := func(auth string, id int) (string, Response) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:9096/pets/%d", id), nil)
		assert.Nil(t, err)
		if auth != "" {
			req.Header.Add("Auth", auth)
		}
		response, err := client.Do(req)
		assert.Nil(t, err)
		body, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		response.Body.Close()
		var rsp Response
		err = json.Unmarshal(body, &rsp)
		assert.Nil(t, err)
		return string(body), rsp
	}

	body, response := request("", 1)
	assert.Equal(t, "", response.Error)
	assert.NotEqual(t, 0, len(body))

	_, response = request("", 8)
	assert.Equal(t, "id must be less than 8", response.Error)

	body, _ = request("1337", 8)
	assert.NotEqual(t, 0, len(body))
}

func TestHandlerRoutingIntegrationAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Handler Routing API integration test in short mode")
	}

	e, err := HandlerRoutingExample()
	assert.Nil(t, err)
	testHandlerRoutingApplication(t, e)
}

func TestHandlerRoutingIntegrationJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Handler Routing JSON integration test in short mode")
	}

	data, err := ioutil.ReadFile(filepath.FromSlash("./json/handler-routing/flogo.json"))
	assert.Nil(t, err)
	cfg, err := engine.LoadAppConfig(string(data), false)
	assert.Nil(t, err)
	e, err := engine.New(cfg)
	assert.Nil(t, err)
	testHandlerRoutingApplication(t, e)
}

func testDefaultHTTPPattern(t *testing.T, e engine.Engine) {
	defer api.ClearResources()
	test.Drain("1234")
	testHandler := handler{}
	s := &http.Server{
		Addr:           ":1234",
		Handler:        &testHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		s.ListenAndServe()
	}()
	test.Pour("1234")
	defer s.Shutdown(context.Background())

	test.Drain("9096")
	err := e.Start()
	assert.Nil(t, err)
	defer func() {
		err := e.Stop()
		assert.Nil(t, err)
	}()
	test.Pour("9096")

	transport := &http.Transport{
		MaxIdleConns: 1,
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{
		Transport: transport,
	}
	request := func() string {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/pets/1", nil)
		assert.Nil(t, err)
		response, err := client.Do(req)
		assert.Nil(t, err)
		body, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		response.Body.Close()
		return string(body)
	}

	body := request()
	assert.NotEqual(t, 0, len(body))
}

func TestDefaultHttpPatternAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway API integration test in short mode")
	}

	e, err := DefaultHTTPPattern()
	assert.Nil(t, err)
	testDefaultHTTPPattern(t, e)
}

func TestDefaultHttpPatternJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway JSON integration test in short mode")
	}
	data, err := ioutil.ReadFile(filepath.FromSlash("./json/default-http-pattern/flogo.json"))
	assert.Nil(t, err)
	cfg, err := engine.LoadAppConfig(string(data), false)
	assert.Nil(t, err)
	e, err := engine.New(cfg)
	assert.Nil(t, err)
	testDefaultHTTPPattern(t, e)
}

func testDefaultChannelPattern(t *testing.T, e engine.Engine) {
	defer api.ClearResources()

	test.Drain("9096")
	err := e.Start()
	assert.Nil(t, err)
	defer func() {
		err := e.Stop()
		assert.Nil(t, err)
	}()
	test.Pour("9096")

	transport := &http.Transport{
		MaxIdleConns: 1,
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{
		Transport: transport,
	}
	request := func() string {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/endpoint", nil)
		assert.Nil(t, err)
		response, err := client.Do(req)
		assert.Nil(t, err)
		body, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		response.Body.Close()
		return string(body)
	}

	body := request()
	assert.NotEqual(t, 0, len(body))
}

func TestDefaultChannelPatternAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway API integration test in short mode")
	}

	e, err := DefaultChannelPattern()
	assert.Nil(t, err)
	testDefaultChannelPattern(t, e)
}

func TestDefaultChannelPatternJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway JSON integration test in short mode")
	}
	data, err := ioutil.ReadFile(filepath.FromSlash("./json/default-channel-pattern/flogo.json"))
	assert.Nil(t, err)
	cfg, err := engine.LoadAppConfig(string(data), false)
	assert.Nil(t, err)
	e, err := engine.New(cfg)
	assert.Nil(t, err)
	testDefaultChannelPattern(t, e)
}

func testAsyncGatewayExample(t *testing.T, e engine.Engine) {
	defer api.ClearResources()

	test.Drain("9096")
	err := e.Start()
	assert.Nil(t, err)
	defer func() {
		err := e.Stop()
		assert.Nil(t, err)
	}()
	test.Pour("9096")

	transport := &http.Transport{
		MaxIdleConns: 1,
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{
		Transport: transport,
	}
	request := func() string {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/endpoint", nil)
		assert.Nil(t, err)
		response, err := client.Do(req)
		assert.Nil(t, err)
		body, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		response.Body.Close()
		return string(body)
	}

	body := request()
	assert.NotEqual(t, 0, len(body))
}

func TestAsyncGatewayExampleAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway API integration test in short mode")
	}

	e, err := AsyncGatewayExample()
	assert.Nil(t, err)
	testAsyncGatewayExample(t, e)
}

func TestAsyncGatewayExampleJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway JSON integration test in short mode")
	}
	data, err := ioutil.ReadFile(filepath.FromSlash("./json/async-gateway/flogo.json"))
	assert.Nil(t, err)
	cfg, err := engine.LoadAppConfig(string(data), false)
	assert.Nil(t, err)
	e, err := engine.New(cfg)
	assert.Nil(t, err)
	testAsyncGatewayExample(t, e)
}

func testResourceHandlerExampleFile(t *testing.T, e engine.Engine) {
	defer api.ClearResources()

	test.Drain("9096")
	err := e.Start()
	assert.Nil(t, err)
	defer func() {
		err := e.Stop()
		assert.Nil(t, err)
	}()
	test.Pour("9096")

	transport := &http.Transport{
		MaxIdleConns: 1,
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{
		Transport: transport,
	}
	request := func() string {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/pets/4", nil)
		assert.Nil(t, err)
		response, err := client.Do(req)
		assert.Nil(t, err)
		body, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		response.Body.Close()
		return string(body)
	}

	body := request()
	assert.NotEqual(t, 0, len(body))
}

func TestResourceHandlerExampleAPI_File(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway API integration test in short mode")
	}

	e, err := FileResourceHandlerExample("file://./json/resource-handler/fileResource/resource.json")
	assert.Nil(t, err)
	testResourceHandlerExampleFile(t, e)
}

func TestResourceHandlerExampleJSON_File(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway JSON integration test in short mode")
	}
	data, err := ioutil.ReadFile(filepath.FromSlash("./json/resource-handler/fileResource/flogo.json"))
	assert.Nil(t, err)
	flogoApp := FlogoJSON{}
	err = json.Unmarshal(data, &flogoApp)
	assert.Nil(t, err)
	flogoApp.Actions[0].Settings["uri"] = "file://./json/resource-handler/fileResource/resource.json"
	data, err = json.Marshal(&flogoApp)
	assert.Nil(t, err)
	cfg, err := engine.LoadAppConfig(string(data), false)
	assert.Nil(t, err)
	e, err := engine.New(cfg)
	assert.Nil(t, err)
	testResourceHandlerExampleFile(t, e)
}

func testResourceHandlerExampleHTTP(t *testing.T, e engine.Engine) {
	defer api.ClearResources()
	test.Drain("9096")
	err := e.Start()
	assert.Nil(t, err)
	defer func() {
		err := e.Stop()
		assert.Nil(t, err)
	}()
	test.Pour("9096")

	transport := &http.Transport{
		MaxIdleConns: 1,
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{
		Transport: transport,
	}
	request := func() string {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/pets/4", nil)
		assert.Nil(t, err)
		response, err := client.Do(req)
		assert.Nil(t, err)
		body, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		response.Body.Close()
		return string(body)
	}

	body := request()
	assert.NotEqual(t, 0, len(body))
}

func TestResourceHandlerExampleAPI_HTTP(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway API integration test in short mode")
	}
	test.Drain("1234")
	testHandler := resourceHandler{}
	s := &http.Server{
		Addr:           ":1234",
		Handler:        &testHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		s.ListenAndServe()
	}()
	test.Pour("1234")
	defer s.Shutdown(context.Background())
	e, err := HTTPResourceHandlerExample()
	assert.Nil(t, err)
	testResourceHandlerExampleHTTP(t, e)
}

func TestResourceHandlerExampleJSON_HTTP(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Basic Gateway JSON integration test in short mode")
	}
	data, err := ioutil.ReadFile(filepath.FromSlash("./json/resource-handler/httpResource/flogo.json"))
	assert.Nil(t, err)
	cfg, err := engine.LoadAppConfig(string(data), false)
	assert.Nil(t, err)
	test.Drain("1234")
	testHandler := resourceHandler{}
	s := &http.Server{
		Addr:           ":1234",
		Handler:        &testHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		s.ListenAndServe()
	}()
	test.Pour("1234")
	defer s.Shutdown(context.Background())

	e, err := engine.New(cfg)
	assert.Nil(t, err)
	testResourceHandlerExampleHTTP(t, e)
}

// FlogoJSON is the flogo JSON
type FlogoJSON struct {
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Version   string      `json:"version"`
	Desc      string      `json:"description"`
	Prop      interface{} `json:"properties"`
	Channels  interface{} `json:"channels"`
	Trig      interface{} `json:"triggers"`
	Resources []struct {
		ID       string `json:"id"`
		Compress bool   `json:"compressed"`
		Data     struct {
			Name      string                   `json:"name"`
			Steps     []interface{}            `json:"steps"`
			Responses []interface{}            `json:"responses"`
			Services  []map[string]interface{} `json:"services"`
		} `json:"data"`
	} `json:"resources"`
	Actions []struct {
		Ref      string                 `json:"ref"`
		Settings map[string]interface{} `json:"settings"`
		ID       string                 `json:"id"`
	} `json:"actions"`
}
