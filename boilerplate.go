package schema

import "fmt"

const selfField = "."

// MatcherFuncImpl is a wrapper for a function to obey the Matcher interface
type MatcherFuncImpl struct {
	Name string
	Fun  func(data interface{}) *Error
}

// Match is the actual matching function
func (f *MatcherFuncImpl) Match(data interface{}) *Error {
	return f.Fun(data)
}

func (f *MatcherFuncImpl) String() string {
	return f.Name
}

func matchValue(fieldError *Error, k string, rawExp interface{}, rawActual interface{}) error {
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
	case Matcher:
		if err := exp.Match(rawActual); err != nil {
			fieldError.Merge(k, err)
		}
	default:
		if rawExp != rawActual {
			fieldError.Add(k, fmt.Sprintf("%#v != %#v", rawActual, exp))
		}
	}

	return nil
}
