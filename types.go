package schema

import "fmt"

var IsInteger = CheckerFunc("IsInteger", isInteger)

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

var IsString = CheckerFunc("IsString", isString)

func isString(data interface{}) *Error {
	if _, ok := data.(string); !ok {
		return SelfError(fmt.Sprintf("is no string but %T", data))
	}

	return nil
}

var IsBool = CheckerFunc("IsBool", isBool)

func isBool(data interface{}) *Error {
	if _, ok := data.(bool); !ok {
		return SelfError(fmt.Sprintf("is no bool but %T", data))
	}

	return nil
}

var IsFloat = CheckerFunc("IsFloat", isFloat)

func isFloat(data interface{}) *Error {
	switch data.(type) {
	case float64, int:
		return nil
	}

	return SelfError(fmt.Sprintf("is no float but %T", data))
}

var IsPresent = CheckerFunc("IsPresent", isPresent)

func isPresent(data interface{}) *Error {
	// Map is checking this implicitly, we only need to be called
	return nil
}
