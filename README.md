# schema

**schema** makes it easier to check if map/array structures match a certain schema. Great for testing JSON API's.

## ideas

"ConcreteValue"
IsPresent
IsString, IsInt, IsFloat, IsBool
Map, MapIncluding
Array, ArrayUnordered, ArrayIncluding, ArrayEach
 
## How to use

    // my_test.go
    
    func TestJSON(t *testing.T) {
        jsonData := getJSON()
        
        var data interface{}
        err := json.Unmarshal(jsonData, &data)
        if err != nil {
            t.Fatal(err)
        }
        
        err := schema.Map{
            "id": schema.IsInteger,
            "name": "Max Mustermann",
            "age": 42,
            "footsize": schema.IsGiven,
        }.Check(data)
        
        if err != nil {
            t.Fatal(err)
        }
    }
    
    
# Issues

- TODO document number issue