# schema

**schema** makes it easier to check if map/array structures match a certain schema. Great for testing JSON API's.

The initial version was built during a [Dynport](https://github.com/dynport) hackday by [gfrey](https://github.com/gfrey) and [jgroenveld](https://github.com/jgroeneveld). It was inspired by [chancancode/json_expressions](https://github.com/chancancode/json_expressions)

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

    Map
        Matches maps where all given keys and values have to match. No extra or missing keys allowed.

    MapIncluding
        Matches maps but only checks the given keys and values and ignores extra ones.

    Array
        Matches all array elements in order.

    ArrayUnordered
        Matches all array elements but order is ignored.

    ArrayIncluding
        Reports elements that can not be matched.

    ArrayEach
        Each element of the array has to match the given matcher.

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

### Input

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
    
### err.Error()

```
"address.street": is no string but float64
"tags": red:string(0) not included
Missing keys: "footsize"
"name": "Hans Meier" != "Max Mustermann"
"address": Missing keys: "zip"
```

## How to write matchers

To use custom or more specialized matchers, the [schema.Matcher](schema.go#L9) interface needs to be implemented.
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
    
## Issues

- TODO document number issue
- TODO document ArrayIncluding/ArrayUnordered issue with complex matchers

- TODO document writing own matchers
- TODO document all existing matchers
