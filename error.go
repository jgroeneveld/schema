package schema

import (
	"fmt"
	"sort"
	"strings"
)

// Error is used to record errors that happen during a schema check
type Error struct {
	Errors map[string]string
}

// Add adds an error
func (e *Error) Add(field, message string) {
	if e.Errors == nil {
		e.Errors = map[string]string{}
	}
	if msg, exists := e.Errors[field]; exists {
		message = msg + ", " + message
	}
	e.Errors[field] = message
}

// Any checks if there are any errors
func (e *Error) Any() bool {
	return len(e.Errors) > 0
}

func (e *Error) Error() string {
	msgs := []string{}
	for field, message := range e.Errors {
		if field == selfField {
			msgs = append(msgs, message)
		} else {
			msgs = append(msgs, fmt.Sprintf("%q: %s", field, message))
		}

	}
	sort.Strings(msgs)
	return strings.Join(msgs, "\n")
}

// Merge merges another error into this error to have a error tree.
func (e *Error) Merge(otherField string, other *Error) {
	for field, msg := range other.Errors {
		f := otherField
		if field != selfField {
			if isErrorIdxField(field) {
				f = fmt.Sprintf("%s%s", otherField, field)
			} else {
				f = fmt.Sprintf("%s.%s", otherField, field)
			}
		}
		e.Add(f, msg)
	}
}

// SelfError is an error without any field it describes
func SelfError(msg string) *Error {
	return &Error{
		Errors: map[string]string{
			selfField: msg,
		},
	}
}
