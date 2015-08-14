package schema

import (
	"fmt"
	"sort"
	"strings"
)

type Map map[string]interface{}

func (m Map) Match(data interface{}) *Error {
	fieldError := &Error{}

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return SelfError(fmt.Sprintf("is no map: %t", data))
	}

	checkExtraKeys(fieldError, m, dataMap)
	checkMissingKeys(fieldError, m, dataMap)
	matchValues(fieldError, m, dataMap)

	if fieldError.Any() {
		return fieldError
	}
	return nil
}

type MapIncluding map[string]interface{}

func (m MapIncluding) Match(data interface{}) *Error {
	fieldError := &Error{}

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return SelfError(fmt.Sprintf("is no map: %t", data))
	}

	checkMissingKeys(fieldError, m, dataMap)
	matchValues(fieldError, m, dataMap)

	if fieldError.Any() {
		return fieldError
	}
	return nil
}

func checkExtraKeys(fieldError *Error, m map[string]interface{}, dataMap map[string]interface{}) {
	extraKeys := []string{}
	for k := range dataMap {
		if _, found := m[k]; !found {
			extraKeys = append(extraKeys, k)
		}
	}
	if len(extraKeys) > 0 {
		sort.Strings(extraKeys)
		fieldError.Add(selfField, fmt.Sprintf("Found extra keys: %q", strings.Join(extraKeys, ", ")))
	}
}

func checkMissingKeys(fieldError *Error, m map[string]interface{}, dataMap map[string]interface{}) {
	missingKeys := []string{}
	for k := range m {
		if _, found := dataMap[k]; !found {
			missingKeys = append(missingKeys, k)
		}
	}
	if len(missingKeys) > 0 {
		sort.Strings(missingKeys)
		fieldError.Add(selfField, fmt.Sprintf("Missing keys: %q", strings.Join(missingKeys, ", ")))
	}
}

func matchValues(fieldError *Error, m map[string]interface{}, dataMap map[string]interface{}) {
	keys := []string{}
	exps := sortableExps{}

	i := 0
	for k, exp := range m {
		keys = append(keys, k)
		exps = append(exps, &origExp{OriginalIndex: i, Exp: exp})
		i++
	}

	sort.Sort(exps)

	for _, exp := range exps {
		k := keys[exp.OriginalIndex]

		actual, found := dataMap[k]
		if !found {
			continue
		}
		matchValue(fieldError, k, exp.Exp, actual)
	}
}
