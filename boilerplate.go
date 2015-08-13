package schema

import "fmt"

type PredicateFunc func(data interface{}) *Error

func (p PredicateFunc) Check(data interface{}) *Error {
	return p(data)
}

func compareValue(fieldError *Error, k string, rawExp interface{}, rawActual interface{}) error {
	switch exp := rawExp.(type) {
	case int:
		switch actual := rawActual.(type) {
		case int:
			if actual != exp {
				fieldError.Add(k, fmt.Sprintf("%#v != %#v", rawActual, exp))
			}
		case float64:
			if float64(int(actual)) != actual || int(actual) != exp {
				fieldError.Add(k, fmt.Sprintf("%#v != %#v", rawActual, exp))
			}
		default:
			fieldError.Add(k, fmt.Sprintf("%#v != %#v", rawActual, exp))
		}
	case string:
		if actual, ok := rawActual.(string); !ok || actual != exp {
			fieldError.Add(k, fmt.Sprintf("%#v != %#v", rawActual, exp))
		}
	case bool:
		if actual, ok := rawActual.(bool); !ok || actual != exp {
			fieldError.Add(k, fmt.Sprintf("%#v != %#v", rawActual, exp))
		}
	case Checker:
		if err := exp.Check(rawActual); err != nil {
			fieldError.Add(k, err.Error())
		}
	default:
		panic("unknown type to check")
	}

	return nil
}
