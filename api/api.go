package api

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/api"
	"github.com/r2d2-ai/aiflow/support"
)

var (
	resources      = make(map[string]*Microgateway)
	resourcesMutex = sync.RWMutex{}
)

// GetResource gets the resource
func GetResource(name string) *Microgateway {
	resourcesMutex.RLock()
	defer resourcesMutex.RUnlock()
	return resources[name]

}

// ClearResources clears the resources for testing
func ClearResources() {
	resources = make(map[string]*Microgateway)
}

// New creates a new microgateway action
func New(name string) *Microgateway {
	return &Microgateway{
		Name: name,
	}
}

// NewService adds a new service to the microgateway
func (m *Microgateway) NewService(name string, act interface{}) *Service {
	service := &Service{
		Name:     name,
		Settings: make(map[string]interface{}),
	}
	switch act := act.(type) {
	case string:
		service.Ref = act
	case activity.Activity:
		if hr, ok := act.(support.HasRef); ok {
			service.Ref = hr.Ref()
		} else {
			value := reflect.ValueOf(act)
			value = value.Elem()
			service.Ref = value.Type().PkgPath()
		}
	case func(ctx activity.Context) (done bool, err error):
		service.Handler = act
	case ServiceFunc:
		service.Handler = act
	default:
		panic("invalid type for act")
	}
	m.Services = append(m.Services, service)
	return service
}

// NewStep adds a new execution step to the microgateway
func (m *Microgateway) NewStep(service *Service) *Step {
	step := &Step{
		Service: service.Name,
		Input:   make(map[string]interface{}),
	}
	m.Steps = append(m.Steps, step)
	return step
}

// NewResponse adds a new response to the microgateway
func (m *Microgateway) NewResponse(isError bool) *Response {
	response := &Response{
		Error: isError,
	}
	m.Responses = append(m.Responses, response)
	return response
}

// SetDescription sets the description of the service
func (s *Service) SetDescription(description string) {
	s.Description = description
}

// AddSetting adds a setting to the service
func (s *Service) AddSetting(name string, value interface{}) {
	s.Settings[name] = value
}

// SetIf sets the execution condition of the step
func (s *Step) SetIf(condition string) {
	s.Condition = condition
}

// AddInput adds an input to the step
func (s *Step) AddInput(name string, value interface{}) {
	s.Input[name] = value
}

// SetHalt sets the halting condition for the step
func (s *Step) SetHalt(condition string) {
	s.HaltCondition = condition
}

// SetIf sets the condition for the response
func (r *Response) SetIf(condition string) {
	r.Condition = condition
}

// SetCode sets the status code for the response
func (r *Response) SetCode(code interface{}) {
	r.Output.Code = code
}

// SetData sets the return data for the response
func (r *Response) SetData(data interface{}) {
	r.Output.Data = data
}

// AddResource adds the microgateway resource to the app and returns the action settings
func (m *Microgateway) AddResource(app *api.App, inputSettings ...map[string]interface{}) (map[string]interface{}, error) {
	name, async := "microgateway:"+m.Name, false
	if len(inputSettings) == 1 && inputSettings[0] != nil {
		if value := inputSettings[0]["uri"]; value != nil {
			if uri, ok := value.(string); ok && uri != "" {
				name = uri
			}
		}
		if value := inputSettings[0]["async"]; value != nil {
			if a, ok := value.(bool); ok {
				async = a
			}
		}
	}
	resourcesMutex.RLock()
	_, ok := resources[name]
	resourcesMutex.RUnlock()
	if ok {
		return nil, fmt.Errorf("resource already exists: %s", name)
	}
	resourcesMutex.Lock()
	resources[name] = m
	resourcesMutex.Unlock()

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	app.AddResource(name, data)
	settings := map[string]interface{}{
		"uri":   name,
		"async": async,
	}
	return settings, nil
}
