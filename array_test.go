package schema

import "testing"

func TestArray_Success(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue", 12, "something_inbetween_we_dont_care_about", true]`)

	err := Array("red", IsString, 12, IsPresent, true).Check(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestArray_WrongOrder(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue"]`)

	err := Array("blue", "red").Check(data)

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
	err := Array("red", "blue").Check(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors["."] != `length does not match 3 != 2` {
		t.Fatalf("wrong error msg: %s", err.Errors["."])
	}
}

func TestArray_MissingElements(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue"]`)
	err := Array("red", "blue", "yellow").Check(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors["."] != `length does not match 2 != 3` {
		t.Fatalf("wrong error msg: %s", err.Errors[selfField])
	}
}

func TestArrayEach_Success(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue"]`)

	err := ArrayEach(IsString).Check(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestArrayEach_Failure(t *testing.T) {
	data := dataFromJSON(t, `["red", 1]`)

	err := ArrayEach(IsString).Check(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors["1"] != `is no string but float64` {
		t.Fatalf("wrong error msg: %s", err.Errors["1"])
	}
}
