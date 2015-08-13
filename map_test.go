package schema

import "testing"

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
