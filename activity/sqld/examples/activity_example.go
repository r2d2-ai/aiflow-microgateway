package examples

import (
	"github.com/r2d2-ai/aiflow/api"
	"github.com/r2d2-ai/aiflow/engine"
	microgateway "github.com/r2d2-ai/aiflow/microgateway"
	"github.com/r2d2-ai/aiflow/microgateway/activity/sqld"
	microapi "github.com/r2d2-ai/aiflow/microgateway/api"
	"github.com/r2d2-ai/contrib/activity/rest"
	trigger "github.com/r2d2-ai/contrib/trigger/rest"
)

// Example returns an API example
func Example() (engine.Engine, error) {
	app := api.NewApp()

	gateway := microapi.New("Update")
	serviceSQLD := gateway.NewService("SQLSecurity", &sqld.Activity{})
	serviceSQLD.SetDescription("Look for sql injection attacks")

	serviceUpdate := gateway.NewService("PetStorePetsUpdate", &rest.Activity{})
	serviceUpdate.SetDescription("Update pets")
	serviceUpdate.AddSetting("uri", "http://petstore.swagger.io/v2/pet")
	serviceUpdate.AddSetting("method", "PUT")
	serviceUpdate.AddSetting("headers", map[string]string{
		"Accept": "application/json",
	})

	step := gateway.NewStep(serviceSQLD)
	step.AddInput("payload", "=$.payload")
	step = gateway.NewStep(serviceUpdate)
	step.SetIf("$.SQLSecurity.outputs.attack < 80")
	step.AddInput("content", "=$.payload.content")

	response := gateway.NewResponse(false)
	response.SetIf("$.SQLSecurity.outputs.attack < 80")
	response.SetCode(200)
	response.SetData("=$.PetStorePetsUpdate.outputs.data")
	response = gateway.NewResponse(true)
	response.SetIf("$.SQLSecurity.outputs.attack > 80")
	response.SetCode(403)
	response.SetData(map[string]interface{}{
		"error":        "hack attack!",
		"attackValues": "=$.SQLSecurity.outputs.attackValues",
	})

	settings, err := gateway.AddResource(app)
	if err != nil {
		return nil, err
	}

	trg := app.NewTrigger(&trigger.Trigger{}, &trigger.Settings{Port: 9096})
	handler, err := trg.NewHandler(&trigger.HandlerSettings{
		Method: "PUT",
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
