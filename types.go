package schema

import "fmt"

var IsInteger = CheckerFunc(isInteger)

func isInteger(data interface{}) *Error {
	switch i := data.(type) {
	case int:
		return nil
	case float64:
		if float64(int(i)) == i {
			return nil
		}
	}

	return SelfError(fmt.Sprintf("is no Integer %q", data))
}

var IsString = CheckerFunc(isString)

func isString(data interface{}) *Error {
	if _, ok := data.(string); !ok {
		return SelfError(fmt.Sprintf("is no string %q", data))
	}

	return nil
}

var IsBool = CheckerFunc(isBool)

func isBool(data interface{}) *Error {
	if _, ok := data.(bool); !ok {
		return SelfError(fmt.Sprintf("is no bool %q", data))
	}

	return nil
}

var IsFloat = CheckerFunc(isFloat)

func isFloat(data interface{}) *Error {
	switch data.(type) {
	case float64, int:
		return nil
	}

	return SelfError(fmt.Sprintf("is no float %q", data))
}

var IsPresent = CheckerFunc(isPresent)

func isPresent(data interface{}) *Error {
	// Map is checking this implicitly, we only need to be called
	return nil
}
