# schema

**schema** makes it easier to check if map/array structures match a certain schema. Great for testing JSON API's.

## ideas

"ConcreteValue"
IsString, IsInt, IsFloat, IsBool
Map, MapIncluding
Array, ArrayUnordered, ArrayIncluding, ArrayEach
 
## How to use

    // my_test.go
    
    func TestJSON(t *testing.T) {
        bodyReader, err := makeRequest()
        if err != nil {
            t.Fatal(err)
        }
        defer bodyReader.Close()
        
        
        err := schema.Map{
            "id": schema.IsInteger,
            "name": "Max Mustermann",
            "age": 42,
            "footsize": schema.IsGiven,
        }.Check(bodyReader)
        
        if err != nil {
            t.Fatal(err)
        }
    }
    
    
# Issues

- TODO document number issue