package schema_test

import (
	"io"
	"strings"
	"testing"

	"github.com/jgroeneveld/schema"
)

func TestJSON(t *testing.T) {
	reader := getJSONResponse()

	err := schema.MatchJSON(
		schema.Map{
			"id":       schema.IsInteger,
			"name":     "Max Mustermann",
			"age":      42,
			"height":   schema.IsFloat,
			"footsize": schema.IsPresent,
			"address": schema.Map{
				"street": schema.IsString,
				"zip":    schema.IsString,
			},
			"tags": schema.ArrayIncluding("red"),
		},
		reader,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func getJSONResponse() io.Reader {
	return strings.NewReader(`{
		"id": 12,
		"name": "Max Mustermann",
		"age": 42,
		"height": 1.91,
		"footsize": 12,
		"address": {
			"street": "Musterstrasse 12",
			"zip": "12345"
		},
		"tags": ["blue", "red", "green"]
	}`)
}

func getJSONResponseFailure() io.Reader {
	return strings.NewReader(`{
		"id": 12,
		"name": "Hans Meier",
		"age": 42,
		"height": 1.91,
		"address": {
			"street": 12
		},
		"tags": ["blue", "green"]
	}`)
}
