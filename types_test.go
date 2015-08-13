package schema

import "testing"

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

func TestIsString_Success(t *testing.T) {
	err := IsString.Check("harald")

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsString_Fail(t *testing.T) {
	err := IsString.Check(12)

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func TestIsBool_Success(t *testing.T) {
	err := IsBool.Check(true)
	if err != nil {
		t.Fatal(err)
	}

	err = IsBool.Check(false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsBool_Fail(t *testing.T) {
	err := IsBool.Check("1")

	if err == nil {
		t.Fatal("Expected error got none")
	}
}

func TestIsFloat_Success(t *testing.T) {
	err := IsFloat.Check(12.42)
	if err != nil {
		t.Fatal(err)
	}

	err = IsFloat.Check(12)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsFloat_Fail(t *testing.T) {
	err := IsFloat.Check("12")

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
