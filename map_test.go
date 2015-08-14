package schema

import "testing"

func TestMap_Success(t *testing.T) {
	data := dataFromJSON(t, `
	{
		"id": 12,
		"name": "Max Mustermann"
	}`)

	err := Map{
		"id":   IsInteger,
		"name": "Max Mustermann",
	}.Match(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMap_ExtraKeys(t *testing.T) {
	data := dataFromJSON(t, `{"hans":true, "wurst": "def"}`)

	err := Map{}.Match(data)

	if err == nil {
		t.Fatal("Expected error got none")
	}
	// check twice because of non-deterministic order of keys
	if err.Errors[selfField] != `Found extra keys: "hans, wurst"` {
		t.Fatalf("wrong error msg: %s", err.Error())
	}
}

func TestMap_MissingKeys(t *testing.T) {
	data := dataFromJSON(t, `{}`)

	err := Map{"id": 12}.Match(data)

	if err == nil {
		t.Fatal("Expected error got none")
	}
	if err.Errors["."] != `Missing keys: "id"` {
		t.Fatalf("wrong error msg: %s", err.Error())
	}
}

func TestMap_WrongValue(t *testing.T) {
	data := dataFromJSON(t, `{"id": 47, "name": "hans", "footsize": "12 inches"}`)

	err := Map{"id": 12, "name": "wurst", "footsize": 12}.Match(data)

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

func TestMapIncluding_Success(t *testing.T) {
	data := dataFromJSON(t, `{"id": 47, "name": "hans", "footsize": "12 inches"}`)

	err := MapIncluding{"id": 47, "name": "hans"}.Match(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMapIncluding_Failure(t *testing.T) {
	data := dataFromJSON(t, `{"id": 12}`)

	err := MapIncluding{"id": 47, "name": "hans"}.Match(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors[selfField] != `Missing keys: "name"` {
		t.Fatalf("wrong error msg: %s", err.Errors[selfField])
	}
	if err.Errors["id"] != `12 != 47` {
		t.Fatalf("wrong error msg: %s", err.Errors["id"])
	}
}
