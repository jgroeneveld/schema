# schema

**schema** makes it easier to check if map/array structures match a certain schema. Great for testing JSON API's or validating the format of incoming requests and providing error messages to the api user.  Also see [trial](https://github.com/jgroeneveld/trial) for simple test assertions.

[godoc](https://godoc.org/github.com/jgroeneveld/schema)

The initial version was built during a [Dynport](https://github.com/dynport) hackday by [gfrey](https://github.com/gfrey) and [jgroenveld](https://github.com/jgroeneveld). It was inspired by [chancancode/json_expressions](https://github.com/chancancode/json_expressions)

## Example

[example_test.go](example_test.go)
```go
func TestJSON(t *testing.T) {
    reader := getJSONResponse()

    err := schema.MatchJSON(
        schema.Map{
            "id":       schema.IsInteger,
            "name":     "Max Mustermann",
            "age":      42,
            "height":   schema.IsFloat,
            "footsize": schema.IsPresent,
            "address": schema.Map{
                "street": schema.IsString,
                "zip":    schema.IsString,
            },
            "tags": schema.ArrayIncluding("red"),
        },
        reader,
    )

    if err != nil {
        t.Fatal(err)
    }
}
```

**JSON Input**

```json
{
    "id": 12,
    "name": "Hans Meier",
    "age": 42,
    "height": 1.91,
    "address": {
        "street": 12
    },
    "tags": ["blue", "green"]
}
```
    
**err.Error() Output**

```
"address": Missing keys: "zip"
"address.street": is no string but float64
"name": "Hans Meier" != "Max Mustermann"
"tags": red:string(0) not included
Missing keys: "footsize"
```

## Entry Points

```go
schema.Match(schema.Matcher, interface{}) error
schema.MatchJSON(schema.Matcher, io.Reader) error
```


## Matchers

    "ConcreteValue"
        Any concrete value like: "Name", 12, true, false, nil

    IsPresent
        Is the value given (empty string counts as given).
        This is essentially a wildcard in map values or array elements.

    Types
        - IsString
        - IsInt
        - IsFloat
        - IsBool
        - IsTime(format)

    Map{"key":Matcher, ...}
        Matches maps where all given keys and values have to match. 
        No extra or missing keys allowed.

    MapIncluding{"key":Matcher, ...}
        Matches maps but only checks the given keys and values and ignores extra ones.

    Array(Matcher...)
        Matches all array elements in order.

    ArrayUnordered(Matcher...)
        Matches all array elements but order is ignored.

    ArrayIncluding(Matcher...)
        Reports elements that can not be matched.

    ArrayEach(Matcher)
        Each element of the array has to match the given matcher.
        
    Capture(name)
        Can be used once or more to capture values and to make sure a value stays the same 
        if it occurs multiple times in a schema. See [capture_test.go](capture_test.go).
        Can also be used to use the captured value in future operations (e.g. multiple requests with the same id).
        
    StringEnum(...values)

## How to write matchers

To use custom or more specialized matchers, the [schema.Matcher](schema.go#L12) interface needs to be implemented.
Either via struct or by using `schema.MatcherFunc`

To report errors, `schema.SelfError(message)` needs to be used if the `data` itself is the problem.

`schema.Error.Add(field, message)` if a subelement of the data is the problem (see `Map` and `Array`).
```go
var IsTime = schema.MatcherFunc("IsTime",
    func(data interface{}) *schema.Error {
        s, ok := data.(string)
        if !ok {
            return schema.SelfError("is no valid time: not a string")
        }

        _, err := time.Parse(time.RFC3339, s)
        if err != nil {
            return schema.SelfError("is no valid time: " + err.Error())
        }
        return nil
    },
)
```

To be more generic with regard to the time format the following pattern can be used:

```go
func IsTime (format string) schema.Matcher
	return schema.MatcherFunc("IsTime",
		func(data interface{}) *schema.Error {
			s, ok := data.(string)
			if !ok {
				return schema.SelfError("is no valid time: not a string")
			}

			_, err := time.Parse(format, s)
			if err != nil {
				return schema.SelfError("is no valid time: " + err.Error())
			}
			return nil
		},
	)
}
```

## Ideas

- write Optional(Matcher) that matches if key is missing or given matcher is satisfied
- write Combine(...Matcher) that matches if all given matchers are satisfied


    
## Issues

**Numbers**

JSON does not differ between integers and floats, ie. there are only numbers. This is why the go JSON library will always return a `float64` value if no type was specified (unmarshalling into an `interface{}` type). This requires some magic internally and can result in false positives. For very large or small integers errors could occur due to rounding errors. If something has no fractional value it is assumed to be equal to an integer (42.0 == 42).


**Array Matchers**

For arrays there are matcher variants for including and unordered. They take the following steps:

* Order the given matchers, where concrete values are matched first, then all matchers except the most generic `IsPresent`, and finally all `IsPresent` matchers. This order guarantees that the most specific values are matched first.
* For each of the ordered matchers, verify one of the remaining values matches.
* Keep a log of all matched values.

This will work in most of the cases, but might fail for some weird nested structures where something like a backtracking approach would be required. 
