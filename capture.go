package schema

import "fmt"

func Capture(name string) CaptureMatcher {
	return &captureMatcher{
		name:  name,
		value: nil,
	}
}

type CaptureMatcher interface {
	Matcher
	Equals(interface{}) bool
	CapturedValue() interface{}
}

type captureMatcher struct {
	name  string
	value interface{}
}

func (m *captureMatcher) CapturedValue() interface{} {
	return m.value
}

func (m *captureMatcher) Equals(expected interface{}) bool {
	if i, isInt := expected.(int); isInt {
		valueAsFloat, valueIsFloat := m.value.(float64)
		if !valueIsFloat {
			// actual is no number
			return false
		}
		if float64(int(valueAsFloat)) != valueAsFloat {
			// actual can not be used as int
			return false
		}
		// int values are distinct?
		return int(valueAsFloat) == i
	}

	return m.value == expected
}

func (m *captureMatcher) Match(data interface{}) *Error {
	if m.value == nil {
		m.value = data
		return nil
	}

	if m.value != data {
		return SelfError(fmt.Sprintf("%s: %v != %v", m.String(), data, m.value))
	}

	return nil
}

func (m *captureMatcher) String() string {
	return fmt.Sprintf("Capture(%s)", m.name)
}
