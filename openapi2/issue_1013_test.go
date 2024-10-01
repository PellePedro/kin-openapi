package openapi2

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/invopop/yaml"
)

func TestCPQ(t *testing.T) {
	specFile := "testdata/cpq.yaml"
	input, err := os.ReadFile(specFile)
	if err != nil {
	    t.Fatalf("Failed to read json file [%s]: %v",  specFile, err)
	}
	var doc2 T
	err = json.Unmarshal(input, &doc2)
	if err != nil {
		// If JSON parsing fails, try YAML
		err = yaml.Unmarshal(input, &doc2)
		if err != nil {
			t.Fatalf("Failed to parse as JSON and YAML: %v", err)
		}
	}
	t.Logf("%+v", doc2)
}
