package schema

import (
	"fmt"
	"time"
)

// IsInteger checks if a value is an integer
var IsInteger = MatcherFunc("IsInteger", isInteger)

func isInteger(data interface{}) *Error {
	switch i := data.(type) {
	case int:
		return nil
	case float64:
		if float64(int(i)) == i {
			return nil
		}
	}

	return SelfError(fmt.Sprintf("is no integer but %T", data))
}

// IsString checks if a value is an string
var IsString = MatcherFunc("IsString", isString)

func isString(data interface{}) *Error {
	if _, ok := data.(string); !ok {
		return SelfError(fmt.Sprintf("is no string but %T", data))
	}

	return nil
}

// IsBool checks if a value is a bool
var IsBool = MatcherFunc("IsBool", isBool)

func isBool(data interface{}) *Error {
	if _, ok := data.(bool); !ok {
		return SelfError(fmt.Sprintf("is no bool but %T", data))
	}

	return nil
}

// IsFloat checks if a value is a float
var IsFloat = MatcherFunc("IsFloat", isFloat)

func isFloat(data interface{}) *Error {
	switch data.(type) {
	case float64, int:
		return nil
	}

	return SelfError(fmt.Sprintf("is no float but %T", data))
}

// IsPresent is a placeholder for Map to check if something "just exists"
var IsPresent = MatcherFunc("IsPresent", isPresent)

func isPresent(data interface{}) *Error {
	// Map is checking this implicitly, we only need to be called
	return nil
}

// IsTime checks if its a parsable time format.
func IsTime(format string) Matcher {
	return MatcherFunc("IsTime",
		func(data interface{}) *Error {
			s, ok := data.(string)
			if !ok {
				return SelfError("is no valid time: not a string")
			}

			_, err := time.Parse(format, s)
			if err != nil {
				return SelfError("is no valid time: " + err.Error())
			}
			return nil
		},
	)
}
