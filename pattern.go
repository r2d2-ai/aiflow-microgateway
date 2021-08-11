package microgateway

import (
	"encoding/json"
	"fmt"

	"github.com/r2d2-ai/aiflow/microgateway/api"
	"github.com/r2d2-ai/aiflow/microgateway/internal/pattern"
)

var patternMap = make(map[string][]byte)

func init() {
	patterns := []string{"DefaultChannelPattern", "DefaultHttpPattern"}
	for i := range patterns {
		patternName := patterns[i] + ".json"
		JSON, err := pattern.Asset(patternName)
		if err != nil {
			fmt.Println("Error from Asset function")
		}
		patternMap[patternName] = JSON
	}
}

// Load loads a pattern
func Load(pattern string) (*api.Microgateway, error) {
	patternJSON := []byte{}
	patternJSON = getPattern(pattern)
	pDef := &api.Microgateway{}
	err := json.Unmarshal(patternJSON, pDef)
	if err != nil {
		return nil, err
	}
	return pDef, nil
}

//Registers a pattern
func Register(patternName string, pattern string) error {
	patternFileName := patternName + ".json"
	if _, ok := patternMap[patternFileName]; !ok {
		patternMap[patternFileName] = []byte(pattern)
	}
	return nil
}

//Returns a registered pattern
func getPattern(pattern string) []byte {
	patternFileName := pattern + ".json"
	if _, ok := patternMap[patternFileName]; ok {
		return patternMap[patternFileName]
	}
	return nil
}
