package schema

import "fmt"

func StringEnum(values ...string) Matcher {
	return MatcherFunc("StringEnum", func(data interface{}) *Error {
		actual, ok := data.(string)
		if !ok {
			return SelfError(fmt.Sprintf("is no string: %t", data))
		}

		for i := 0; i < len(values); i++ {
			enum := values[i]
			if actual == enum {
				return nil
			}
		}
		return SelfError(fmt.Sprintf("%q not included in %q", actual, values))
	})
}
