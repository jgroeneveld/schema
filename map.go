package schema

import (
	"fmt"
	"strings"
)

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
		fieldError.Add(selfField, fmt.Sprintf("Found extra keys: %q", strings.Join(extraKeys, ", ")))
	}

	missingKeys := []string{}
	for k := range m {
		if _, found := dataMap[k]; !found {
			missingKeys = append(missingKeys, k)
		}
	}
	if len(missingKeys) > 0 {
		fieldError.Add(selfField, fmt.Sprintf("Missing keys: %q", strings.Join(missingKeys, ", ")))
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
