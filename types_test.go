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
