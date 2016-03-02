package schema

import (
	"testing"
	"time"
)

func TestIsInteger_Success(t *testing.T) {
	err := IsInteger.Match(12)

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsInteger_Fail(t *testing.T) {
	err := IsInteger.Match("12")

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func TestIsString_Success(t *testing.T) {
	err := IsString.Match("harald")

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsString_Fail(t *testing.T) {
	err := IsString.Match(12)

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func TestIsBool_Success(t *testing.T) {
	err := IsBool.Match(true)
	if err != nil {
		t.Fatal(err)
	}

	err = IsBool.Match(false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsBool_Fail(t *testing.T) {
	err := IsBool.Match("1")

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func TestIsFloat_Success(t *testing.T) {
	err := IsFloat.Match(12.42)
	if err != nil {
		t.Fatal(err)
	}

	err = IsFloat.Match(12)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsFloat_Fail(t *testing.T) {
	err := IsFloat.Match("12")

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func TestIsPresent_Success(t *testing.T) {
	err := IsPresent.Match("12")

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsPresent_Fail(t *testing.T) {
	err := IsPresent.Match(nil)

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
	}.Match(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsPresent_MapFailure(t *testing.T) {
	data := dataFromJSON(t, `{}`)

	err := Map{
		"footsize": IsPresent,
	}.Match(data)

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func TestIsTime_Success(t *testing.T) {
	err := IsTime(time.RFC3339).Match("2016-02-28T12:42:00Z")

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsTime_Fail(t *testing.T) {
	err := IsTime(time.RFC3339).Match("12")

	if err == nil {
		t.Fatal("Expected error got none")
	}
}
