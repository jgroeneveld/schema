// schema makes it easier to check if map/array structures match a certain schema. Great for testing JSON API's.
//
// Example:
//
//  func TestJSON(t *testing.T) {
//      reader := getJSONResponse()
//
//      err := schema.MatchJSON(
//          schema.Map{
//              "id":       schema.IsInteger,
//              "name":     "Max Mustermann",
//              "age":      42,
//              "height":   schema.IsFloat,
//              "footsize": schema.IsPresent,
//              "address":  schema.Map{
//                  "street": schema.IsString,
//                  "zip":    schema.IsString,
//              },
//              "tags": schema.ArrayIncluding("red"),
//          },
//          reader,
//      )
//
//      if err != nil {
//          t.Fatal(err)
//      }
//  }
//
// JSON Input
//
//  {
//      "id": 12,
//      "name": "Hans Meier",
//      "age": 42,
//      "height": 1.91,
//      "address": {
//          "street": 12
//      },
//      "tags": ["blue", "green"]
//  }
//
// err.Error() Output
//
//  "address": Missing keys: "zip"
//  "address.street": is no string but float64
//  "name": "Hans Meier" != "Max Mustermann"
//  "tags": red:string(0) not included
//  Missing keys: "footsize"
//
// See https://github.com/jgroeneveld/schema for more examples.
//
// Also see https://github.com/jgroeneveld/trial for lightweight assertions.
package schema

import (
	"encoding/json"
	"fmt"
	"io"
)

type Matcher interface {
	Match(data interface{}) *Error
}

func MatcherFunc(name string, fun func(data interface{}) *Error) *matcherFunc {
	return &matcherFunc{
		Name: name,
		Fun:  fun,
	}
}

// Match wraps matcher.Match for nil error handling.
func Match(m Matcher, data interface{}) error {
	if err := m.Match(data); err != nil {
		return err
	}
	return nil
}

// MatchJSON wraps Match with a reader for JSON raw data.
func MatchJSON(m Matcher, r io.Reader) error {
	var data interface{}
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return fmt.Errorf("cant parse JSON: %s", err.Error())
	}

	return Match(m, data)
}
