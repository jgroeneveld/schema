package schema

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestFullSuccess(t *testing.T) {
	data := dataFromJSON(t, `
	{
		"id": 12,
		"name": "Max Mustermann",
		"age": 42,
		"footsize": "unknown",
		"address": {
			"street": "Musterstrasse 12",
			"zip": "12345",
			"country": "Germany"
		},
		"tags": ["blue","red"]
	}`)

	err := Map{
		"id":       IsInteger,
		"name":     "Max Mustermann",
		"age":      42,
		"footsize": IsPresent,
		"address": Map{
			"street":  IsString,
			"zip":     IsString,
			"country": IsString,
		},
		"tags": Array{"blue", "red"},
	}.Check(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMap_ExtraKeys(t *testing.T) {
	data := dataFromJSON(t, `{"hans":true, "wurst": "def"}`)

	err := Map{}.Check(data)

	if err == nil {
		t.Fatal("Expected error got none")
	}
	if err.Errors["."] != `Found extra keys: "hans, wurst"` {
		t.Fatalf("wrong error msg: %s", err.Error())
	}
}

func TestMap_MissingKeys(t *testing.T) {
	data := dataFromJSON(t, `{}`)

	err := Map{"id": 12}.Check(data)

	if err == nil {
		t.Fatal("Expected error got none")
	}
	if err.Errors["."] != `Missing keys: "id"` {
		t.Fatalf("wrong error msg: %s", err.Error())
	}
}

func TestMap_WrongValue(t *testing.T) {
	data := dataFromJSON(t, `{"id": 47, "name": "hans", "footsize": "12 inches"}`)

	err := Map{"id": 12, "name": "wurst", "footsize": 12}.Check(data)

	if err == nil {
		t.Fatal("Expected error got none")
	}
	if err.Errors["id"] != `47 != 12` {
		t.Fatalf("wrong error msg: %s", err.Errors["id"])
	}
	if err.Errors["name"] != `"hans" != "wurst"` {
		t.Fatalf("wrong error msg: %s", err.Errors["name"])
	}
	if err.Errors["footsize"] != `"12 inches" != 12` {
		t.Fatalf("wrong error msg: %s", err.Errors["footsize"])
	}
}

func TestArray_Success(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue", 12, "something_inbetween_we_dont_care_about", true]`)

	err := Array{"red", IsString, 12, IsPresent, true}.Check(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestArray_WrongOrder(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue"]`)

	err := Array{"blue", "red"}.Check(data)

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Errors["0"] != `"red" != "blue"` {
		t.Fatalf("wrong error msg: %s", err.Errors["0"])
	}
	if err.Errors["1"] != `"blue" != "red"` {
		t.Fatalf("wrong error msg: %s", err.Errors["1"])
	}
}

func TestArray_ExtraElements(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue", "yellow"]`)
	err := Array{"red", "blue"}.Check(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors["."] != `length does not match 3 != 2` {
		t.Fatalf("wrong error msg: %s", err.Errors["."])
	}
}

func TestArray_MissingElements(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue"]`)
	err := Array{"red", "blue", "yellow"}.Check(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors["."] != `length does not match 2 != 3` {
		t.Fatalf("wrong error msg: %s", err.Errors["."])
	}
}

func TestIsInteger_Success(t *testing.T) {
	err := IsInteger.Check(12)

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsInteger_Fail(t *testing.T) {
	err := IsInteger.Check("12")

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func TestIsPresent_Success(t *testing.T) {
	err := IsPresent.Check("12")

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsPresent_Fail(t *testing.T) {
	err := IsPresent.Check(nil)

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsPresent_MapSuccess(t *testing.T) {
	data := dataFromJSON(t, `
	{
		"footsize": "unknown"
	}`)

	err := Map{
		"footsize": IsPresent,
	}.Check(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsPresent_MapFailure(t *testing.T) {
	data := dataFromJSON(t, `{}`)

	err := Map{
		"footsize": IsPresent,
	}.Check(data)

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func dataFromJSON(t *testing.T, jsonStr string) interface{} {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("can not get runtime caller")
	}
	location := fmt.Sprintf("%s:%d", path.Base(file), line)

	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		t.Fatalf("Fail to unmarshal JSON: %s\nsource: %s", err, location)
	}
	return data
}
