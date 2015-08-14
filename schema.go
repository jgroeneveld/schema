// schema makes it easier to check if map/array structures match a certain schema. Great for testing JSON API's.
//
// For examples see https://github.com/jgroeneveld/schema/blob/master/README.md
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
