package schema

import (
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	Errors map[string]string
}

func (e *Error) Add(field, message string) {
	if e.Errors == nil {
		e.Errors = map[string]string{}
	}
	if msg, exists := e.Errors[field]; exists {
		message = msg + ", " + message
	}
	e.Errors[field] = message
}

func (e *Error) Any() bool {
	return len(e.Errors) > 0
}

func (e *Error) Error() string {
	msgs := []string{}
	for field, message := range e.Errors {
		msgs = append(msgs, fmt.Sprintf("%q: %s", field, message))
	}
	return strings.Join(msgs, "\n")
}

func SelfError(msg string) *Error {
	return &Error{
		Errors: map[string]string{
			".": msg,
		},
	}
}

type Checker interface {
	Check(data interface{}) *Error
}

type Map map[string]interface{}

func (m Map) Check(data interface{}) *Error {
	fieldError := &Error{}

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return SelfError(fmt.Sprintf("is no map: %t", data))
	}

	extraKeys := []string{}
	for k := range dataMap {
		if _, found := m[k]; !found {
			extraKeys = append(extraKeys, k)
		}
	}
	if len(extraKeys) > 0 {
		fieldError.Add(".", fmt.Sprintf("Found extra keys: %q", strings.Join(extraKeys, ", ")))
	}

	missingKeys := []string{}
	for k := range m {
		if _, found := dataMap[k]; !found {
			missingKeys = append(missingKeys, k)
		}
	}
	if len(missingKeys) > 0 {
		fieldError.Add(".", fmt.Sprintf("Missing keys: %q", strings.Join(missingKeys, ", ")))
	}

	for k, exp := range m {
		actual, found := dataMap[k]
		if !found {
			continue
		}
		compareValue(fieldError, k, exp, actual)
	}

	if fieldError.Any() {
		return fieldError
	}
	return nil
}

type Array []interface{}

func (a Array) Check(data interface{}) *Error {
	fieldError := &Error{}

	dataArray, ok := data.([]interface{})
	if !ok {
		return SelfError(fmt.Sprintf("is no array: %t", data))
	}

	if len(a) != len(dataArray) {
		fieldError.Add(".", fmt.Sprintf("length does not match %d != %d", len(dataArray), len(a)))
	}

	for i := 0; i < len(a) && i < len(dataArray); i++ {
		actual := dataArray[i]
		exp := a[i]
		compareValue(fieldError, strconv.Itoa(i), exp, actual)
	}

	if fieldError.Any() {
		return fieldError
	}
	return nil
}

var IsInteger = PredicateFunc(isInteger)

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

var IsString = PredicateFunc(isString)

func isString(data interface{}) *Error {
	if _, ok := data.(string); !ok {
		return SelfError(fmt.Sprintf("is no string %q", data))
	}

	return nil
}

var IsPresent = PredicateFunc(isPresent)

func isPresent(data interface{}) *Error {
	// Map is checking this implicitly, we only need to be called
	return nil
}
