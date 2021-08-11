package microgateway

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/r2d2-ai/aiflow/engine"
	_ "github.com/r2d2-ai/aiflow/examples/action"
)

var testFlogoJSON = `{
  "name": "Test",
  "type": "flogo:app",
  "version": "1.0.0",
  "description": "This is a test application.",
	"imports": [
		"nil github.com/r2d2-ai/core@v0.9.5-0.20191107201139-c4e948d4cc8d",
		"github.com/r2d2-ai/aiflow/examples/action",
		"github.com/r2d2-ai/aiflow/examples/trigger",
		"_ github.com/r2d2-ai/aiflow/data/expression/script"
	],
  "properties": [
		{"name": "test0", "type": "string", "value": "test"},
		{"name": "test1", "type": "int", "value": 1},
		{"name": "test2", "type": "bool", "value": true}
	],
  "channels": [
    "test0:1",
		"test1:2",
		"test2:3"
  ],
  "triggers": [
    {
      "name": "flogo-test0",
      "id": "test0",
      "ref": "#trigger",
      "settings": {
        "aSetting": 123
      },
      "handlers": [
        {
          "settings": {
            "aSetting": 123
          },
          "actions": [
            {
              "id": "action:Test0"
            }
          ]
        }
      ]
    },
		{
      "name": "flogo-test1",
      "id": "test1",
      "ref": "#trigger",
      "settings": {
        "aSetting": 123
      },
      "handlers": [
        {
          "settings": {
            "aSetting": 123
          },
          "actions": [
            {
              "id": "action:Test1"
            }
          ]
        }
      ]
    },
		{
      "name": "flogo-test1",
      "id": "test1",
      "ref": "github.com/r2d2-ai/aiflow/examples/trigger",
      "settings": {
        "aSetting": 123
      },
      "handlers": [
        {
          "settings": {
            "aSetting": 123
          },
          "actions": [
            {
							"if": "1 == 1",
              "id": "action:Test1",
							"input": {
								"test0": "=1",
								"test1": "=2",
								"test2": "=3"
							},
							"output": {
								"test0": "=1",
								"test1": "=2",
								"test2": "=3"
							}
            },
						{
							"ref": "github.com/r2d2-ai/aiflow/examples/action",
				      "settings": {
				        "aSetting": "action:Test"
				      }
						}
          ]
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "action:Test",
      "compressed": false,
      "data": {
				"message": "hello world"
			}
    }
  ],
  "actions": [
    {
      "ref": "github.com/r2d2-ai/aiflow/examples/action",
      "settings": {
        "aSetting": "action:Test"
      },
      "id": "action:Test0",
      "metadata": null
    },
		{
      "ref": "github.com/r2d2-ai/aiflow/examples/action",
      "settings": {
        "aSetting": "action:Test"
      },
      "id": "action:Test1",
      "metadata": null
    }
  ]
}`

func TestGenerate(t *testing.T) {
	app, err := engine.LoadAppConfig(testFlogoJSON, false)
	if err != nil {
		t.Fatal(err)
	}
	current, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Chdir(current)
		if err != nil {
			t.Fatal(err)
		}
	}()
	tmp, err := ioutil.TempDir("", "generate")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tmp)
	err = os.Chdir(tmp)
	if err != nil {
		t.Fatal(err)
	}
	Generate(app, "./test.go", "./go.mod")
	cmd := exec.Command("go", "build")
	err = cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll(tmp)
	if err != nil {
		t.Fatal(err)
	}
}
