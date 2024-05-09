package view

import (
	"encoding/json"
	"testing"
)

var v *View

func TestRenderJSON(t *testing.T) {
	type Result struct {
		Test struct {
			Name    string
			Version int
		}
	}

	type TestData struct {
		Name    string
		Version int
	}

	testData := TestData{
		Name:    "Unit",
		Version: 1,
	}

	jsonResponse, err := v.RenderJSON(Envelope{"Test": testData})

	if err != nil {
		t.Errorf("RenderJSON failed: %v", err)
	}

	var result Result
	err = json.Unmarshal(jsonResponse, &result)

	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if testData.Name != result.Test.Name {
		t.Errorf("Rendered JSON name doesn't match. Expected: %s, Got: %s", testData.Name, result.Test.Name)
	}

	if testData.Version != result.Test.Version {
		t.Errorf("Rendered JSON email doesn't match. Expected: %d, Got: %d", testData.Version, result.Test.Version)
	}
}
