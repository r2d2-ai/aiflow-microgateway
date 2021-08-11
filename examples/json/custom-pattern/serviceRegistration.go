package custom_pattern

import microgateway "github.com/r2d2-ai/aiflow/microgateway"

const pattern = `{
  "name": "CustomPattern",
  "steps": [
    {
      "service": "HttpBackend",
      "halt": "($.HttpBackend.error != nil) && !error.isneterror($.HttpBackend.error)"
    },
    {
      "if": "$.HttpBackend.error == nil",
      "service": "SuccessCounter"
    },
    {
      "if": "$.HttpBackend.error != nil",
      "service": "ErrorCounter"
    },
    {
      "service": "GetCounterSuccess"
    },
    {
      "service": "GetCounterError"
    }
  ],
  "responses": [
    {
      "if" : "$.GetCounterSuccess.error == nil",
      "error": false,
      "output": {
        "code": 200,
        "data": {
          "Success-Calls": "=$.GetCounterSuccess.outputs.value",
          "Error-Calls": "=$.GetCounterError.outputs.value"
        }
      }
    },
    {
      "error": true,
      "output": {
        "code": 400,
        "data": "Error"
      }
    }
  ],
  "services": [
    {
      "name": "HttpBackend",
      "description": "Make an http call to your backend",
      "ref": "github.com/r2d2-ai/contrib/activity/rest",
      "settings": {
        "method": "GET",
        "uri": "http://localhost:1234/pets"
      }
    },
    {
      "name": "SuccessCounter",
      "description": "Increment counter on successful call",
      "ref": "github.com/r2d2-ai/contrib/activity/counter",
      "settings": {
        "counterName": "SuccessCounter",
        "op": "increment"
      }
    },
    {
      "name": "ErrorCounter",
      "description": "Increment counter on error call",
      "ref": "github.com/r2d2-ai/contrib/activity/counter",
      "settings": {
        "counterName": "ErrorCounter",
        "op": "increment"
      }
    },
    {
      "name": "GetCounterSuccess",
      "description": "Get success counter",
      "ref": "github.com/r2d2-ai/contrib/activity/counter",
      "settings": {
        "counterName": "SuccessCounter"
      }
    },
    {
      "name": "GetCounterError",
      "description": "Get error counter",
      "ref": "github.com/r2d2-ai/contrib/activity/counter",
      "settings": {
        "counterName": "ErrorCounter"
      }
    }
  ]
}`

func init() {
	err := microgateway.Register("CustomPattern", pattern)
	if err != nil {
		panic(err)
	}
}
