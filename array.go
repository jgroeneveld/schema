package schema

import (
	"fmt"
	"sort"
	"strings"
)

func Array(exps ...interface{}) Matcher {
	return MatcherFunc("Array", func(data interface{}) *Error {
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
			matchValue(fieldError, errorIdxField(i), exp, actual)
		}

		if fieldError.Any() {
			return fieldError
		}
		return nil
	})
}

func ArrayEach(exp interface{}) Matcher {
	return MatcherFunc("ArrayEach", func(data interface{}) *Error {
		fieldError := &Error{}

		dataArray, ok := data.([]interface{})
		if !ok {
			return SelfError(fmt.Sprintf("is no array: %t", data))
		}

		for i := 0; i < len(dataArray); i++ {
			actual := dataArray[i]
			matchValue(fieldError, errorIdxField(i), exp, actual)
		}

		if fieldError.Any() {
			return fieldError
		}
		return nil
	})
}

func ArrayUnordered(exps ...interface{}) Matcher {
	return MatcherFunc("ArrayUnordered", func(data interface{}) *Error {
		fieldError := &Error{}

		dataArray, ok := data.([]interface{})
		if !ok {
			return SelfError(fmt.Sprintf("is no array: %t", data))
		}

		if len(exps) != len(dataArray) {
			fieldError.Add(selfField, fmt.Sprintf("length does not match %d != %d", len(dataArray), len(exps)))
		}

		matchedIndices, err := arrayIncludingMatchedIndices(exps, dataArray)
		if err != nil {
			fieldError.Merge(selfField, err)
		}

		if len(matchedIndices) != len(dataArray) {
			for i := 0; i < len(dataArray); i++ {
				if !matchedIndices[i] {
					fieldError.Add(errorIdxField(i), "unmatched")
				}
			}
		}

		if fieldError.Any() {
			return fieldError
		}
		return nil
	})
}

func ArrayIncluding(exps ...interface{}) Matcher {
	return MatcherFunc("ArrayIncluding", func(data interface{}) *Error {
		dataArray, ok := data.([]interface{})
		if !ok {
			return SelfError(fmt.Sprintf("is no array: %t", data))
		}

		_, err := arrayIncludingMatchedIndices(exps, dataArray)
		return err
	})
}

func arrayIncludingMatchedIndices(exps []interface{}, dataArray []interface{}) (matchedIndices map[int]bool, err *Error) {
	fieldError := &Error{}

	sortableExps := make([]*origExp, len(exps))
	for i, exp := range exps {
		sortableExps[i] = &origExp{OriginalIndex: i, Exp: exp}
	}
	sortExps(sortableExps)

	matchedIndices = map[int]bool{}
	for _, exp := range sortableExps {
		foundMatching := false

		for i, v := range dataArray {
			if matchedIndices[i] {
				continue
			}
			e := &Error{}
			matchValue(e, errorIdxField(i), exp.Exp, v)
			if !e.Any() {
				matchedIndices[i] = true
				foundMatching = true
				break
			}
		}

		if !foundMatching {
			switch t := exp.Exp.(type) {
			case Matcher:
				fieldError.Add(selfField, fmt.Sprintf("%s<%d> did not match", t, exp.OriginalIndex))
			default:
				fieldError.Add(selfField, fmt.Sprintf("%v:%T<%d> not included", t, t, exp.OriginalIndex))
			}
		}
	}

	if fieldError.Any() {
		return matchedIndices, fieldError
	}
	return matchedIndices, nil
}

func sortExps(exps sortableExps) {
	sort.Sort(exps)
}

type origExp struct {
	OriginalIndex int
	Exp           interface{}
}

type sortableExps []*origExp

func (exps sortableExps) Len() int {
	return len(exps)
}

func (exps sortableExps) Swap(i, j int) {
	exps[i], exps[j] = exps[j], exps[i]
}

func (exps sortableExps) Less(i, j int) bool {
	a, b := exps[i], exps[j]
	return expPrio(a.Exp) < expPrio(b.Exp)
}

func expPrio(a interface{}) int {
	_, isGenericMatcher := a.(Matcher)
	_, isCaptureMatcher := a.(CaptureMatcher)

	// order is relevant
	switch {
	case a == IsPresent:
		return 9
	case isCaptureMatcher:
		return 7
	case isGenericMatcher:
		return 8
	}
	return 0
}

func errorIdxField(i int) string {
	return fmt.Sprintf("[%d]", i)
}

func isErrorIdxField(f string) bool {
	return strings.HasPrefix(f, "[") && strings.HasSuffix(f, "]")
}
