# schema

**schema** makes it easier to check if map/array structures match a certain schema. Great for testing JSON API's.

The initial version was built during a [Dynport](https://github.com/dynport) hackday by [gfrey](https://github.com/gfrey) and [jgroenveld](https://github.com/jgroeneveld) and was inspired by [chancancode/json_expressions](https://github.com/chancancode/json_expressions)

## Matchers

- "ConcreteValue"
- IsPresent, IsString, IsInt, IsFloat, IsBool
- Map, MapIncluding
- Array, ArrayUnordered, ArrayIncluding, ArrayEach
 
## Example

    // example_test.go
    
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
    
### Input

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
    
    
### err.Error()

    "address.street": is no string but float64
    "tags": red:string(0) not included
    Missing keys: "footsize"
    "name": "Hans Meier" != "Max Mustermann"
    "address": Missing keys: "zip"
    
# Issues

- TODO document number issue
- TODO document ArrayIncluding/ArrayUnordered issue with complex matchers

- TODO document writing own matchers
- TODO document all existing matchers
