package schema

import (
	"encoding/json"
	"io"
)

type Checker interface {
	Check(data interface{}) *Error
}

func CheckerFunc(name string, fun func(data interface{}) *Error) *checkerFunc {
	return &checkerFunc{
		Name: name,
		Fun:  fun,
	}
}

// Check wraps checker.Check for nil error handling.
func Check(checker Checker, data interface{}) error {
	if err := checker.Check(data); err != nil {
		return err
	}
	return nil
}

// CheckJSON wraps Check with a reader for JSON raw data.
func CheckJSON(checker Checker, r io.Reader) error {
	var data interface{}
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return err
	}

	return Check(checker, data)
}
