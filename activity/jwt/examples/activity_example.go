package examples

import (
	"github.com/r2d2-ai/aiflow/api"
	"github.com/r2d2-ai/aiflow/engine"
	microgateway "github.com/r2d2-ai/aiflow/microgateway"
	"github.com/r2d2-ai/aiflow/microgateway/activity/jwt"
	microapi "github.com/r2d2-ai/aiflow/microgateway/api"
	"github.com/r2d2-ai/contrib/activity/rest"
	trigger "github.com/r2d2-ai/contrib/trigger/rest"
)

// Example returns an API example
func Example() (engine.Engine, error) {
	app := api.NewApp()
	gateway := microapi.New("JWT")

	jwtService := gateway.NewService("jwtService", &jwt.Activity{})
	jwtService.SetDescription("Validate JWT")
	jwtService.AddSetting("signingMethod", "HMAC")
	jwtService.AddSetting("key", "qwertyuiopasdfghjklzxcvbnm789101")
	jwtService.AddSetting("aud", "www.mashling.io")
	jwtService.AddSetting("iss", "Mashling")
	jwtService.AddSetting("sub", "tempuser@mail.com")

	serviceStore := gateway.NewService("PetStorePets", &rest.Activity{})
	serviceStore.SetDescription("Get pets by ID from the petstore")
	serviceStore.AddSetting("uri", "https://petstore.swagger.io/v2/pet/:petId")
	serviceStore.AddSetting("method", "GET")
	serviceStore.AddSetting("headers", map[string]string{
		"Accept": "application/json",
	})

	step := gateway.NewStep(jwtService)
	step.AddInput("token", "=$.payload.headers.Authorization")
	step = gateway.NewStep(serviceStore)
	step.AddInput("pathParams.petId", "=$.jwtService.outputs.token.claims.id")

	response := gateway.NewResponse(false)
	response.SetIf("$.jwtService.outputs.valid == true")
	response.SetCode(200)
	response.SetData(map[string]interface{}{
		"error": "JWT token is valid",
		"pet":   "=$.PetStorePets.outputs.data",
	})
	response = gateway.NewResponse(true)
	response.SetIf("$.jwtService.outputs.valid == false")
	response.SetCode(401)
	response.SetData(map[string]interface{}{
		"error": "=$.jwtService.outputs",
		"pet":   nil,
	})

	settings, err := gateway.AddResource(app)
	if err != nil {
		return nil, err
	}

	trg := app.NewTrigger(&trigger.Trigger{}, &trigger.Settings{Port: 9096})
	handler, err := trg.NewHandler(&trigger.HandlerSettings{
		Method: "GET",
		Path:   "/pets",
	})
	if err != nil {
		return nil, err
	}

	_, err = handler.NewAction(&microgateway.Action{}, settings)
	if err != nil {
		return nil, err
	}

	return api.NewEngine(app)
}
