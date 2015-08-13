package schema

import (
	"fmt"
	"strconv"
)

func Array(exps ...interface{}) Checker {
	return CheckerFunc(func(data interface{}) *Error {
		fieldError := &Error{}

		dataArray, ok := data.([]interface{})
		if !ok {
			return SelfError(fmt.Sprintf("is no array: %t", data))
		}

		if len(exps) != len(dataArray) {
			fieldError.Add(selfField, fmt.Sprintf("length does not match %d != %d", len(dataArray), len(exps)))
		}

		for i := 0; i < len(exps) && i < len(dataArray); i++ {
			actual := dataArray[i]
			exp := exps[i]
			compareValue(fieldError, strconv.Itoa(i), exp, actual)
		}

		if fieldError.Any() {
			return fieldError
		}
		return nil
	})
}

func ArrayEach(exp interface{}) Checker {
	return CheckerFunc(func(data interface{}) *Error {
		fieldError := &Error{}

		dataArray, ok := data.([]interface{})
		if !ok {
			return SelfError(fmt.Sprintf("is no array: %t", data))
		}

		for i := 0; i < len(dataArray); i++ {
			actual := dataArray[i]
			compareValue(fieldError, strconv.Itoa(i), exp, actual)
		}

		if fieldError.Any() {
			return fieldError
		}
		return nil
	})
}
