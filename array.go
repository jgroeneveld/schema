package schema

import (
	"fmt"
	"strconv"
)

type Array []interface{}

func (a Array) Check(data interface{}) *Error {
	fieldError := &Error{}

	dataArray, ok := data.([]interface{})
	if !ok {
		return SelfError(fmt.Sprintf("is no array: %t", data))
	}

	if len(a) != len(dataArray) {
		fieldError.Add(selfField, fmt.Sprintf("length does not match %d != %d", len(dataArray), len(a)))
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
