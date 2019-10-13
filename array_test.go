package schema

import "testing"

func TestArray_Success(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue", 12, "something_we_dont_care_about", true]`)

	err := Array("red", IsString, 12, IsPresent, true).Match(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestArray_WrongOrder(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue"]`)

	err := Array("blue", "red").Match(data)

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Errors["[0]"] != `"red" != "blue"` {
		t.Fatalf("wrong error msg: %s", err.Errors["[0]"])
	}
	if err.Errors["[1]"] != `"blue" != "red"` {
		t.Fatalf("wrong error msg: %s", err.Errors["[1]"])
	}
}

func TestArray_ExtraElements(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue", "yellow"]`)
	err := Array("red", "blue").Match(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors["."] != `length does not match 3 != 2` {
		t.Fatalf("wrong error msg: %s", err.Errors["."])
	}
}

func TestArray_MissingElements(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue"]`)
	err := Array("red", "blue", "yellow").Match(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors["."] != `length does not match 2 != 3` {
		t.Fatalf("wrong error msg: %s", err.Errors[selfField])
	}
}

func TestArrayEach_Success(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue"]`)

	err := ArrayEach(IsString).Match(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestArrayEach_Failure(t *testing.T) {
	data := dataFromJSON(t, `["red", 1]`)

	err := ArrayEach(IsString).Match(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors["[1]"] != `is no string but float64` {
		t.Fatalf("wrong error msg: %s", err.Errors["[1]"])
	}
}

func TestArrayIncluding_Success(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue", 12, {"a": 1}]`)

	err := ArrayIncluding(IsInteger, Map{"a": IsInteger}, "red").Match(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestArrayIncluding_Failure(t *testing.T) {
	data := dataFromJSON(t, `["red", "blue", {"a": 1}]`)

	err := ArrayIncluding(IsInteger, "green", Map{"a": IsInteger}, "red").Match(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors[selfField] != `green:string<1> not included, IsInteger<0> did not match` {
		t.Fatalf("wrong error msg: %s", err.Errors[selfField])
	}
}

func TestArrayIncluding_MatcherPrio(t *testing.T) {
	data := dataFromJSON(t, `[ 1,2,3,4,5,6 ]`)

	err := ArrayIncluding(IsPresent, IsInteger, 1, 2).Match(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestArrayUnordered_Success(t *testing.T) {
	data := dataFromJSON(t, `["red", 12, {"a": 1}]`)

	err := ArrayUnordered(IsInteger, Map{"a": IsInteger}, "red").Match(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestArrayUnordered_ExtraElements(t *testing.T) {
	data := dataFromJSON(t, `["red", 12, 42, {"a": 1}]`)

	err := ArrayUnordered(IsInteger, Map{"a": IsInteger}, "red").Match(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors[selfField] != `length does not match 4 != 3` {
		t.Fatalf("wrong error msg: %s", err.Errors[selfField])
	}
	if err.Errors["[2]"] != `unmatched` {
		t.Fatalf("wrong error msg: %s", err.Errors["[2]"])
	}
}

func TestArrayUnordered_MissingElement(t *testing.T) {
	data := dataFromJSON(t, `[12, {"a": 1}]`)

	err := ArrayUnordered(IsInteger, Map{"a": IsInteger}, "red").Match(data)

	if err == nil {
		t.Fatal("expected error")
	}
	if err.Errors[selfField] != `length does not match 2 != 3, red:string<2> not included` {
		t.Fatalf("wrong error msg: %s", err.Errors[selfField])
	}
}
