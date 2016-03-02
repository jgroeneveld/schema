package schema

import (
	"testing"
)

func TestStringEnum_Success(t *testing.T) {
	err := StringEnum("GET", "POST").Match("GET")
	if err != nil {
		t.Fatal(err)
	}

	err = StringEnum("GET", "POST").Match("POST")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStringEnum_Failure(t *testing.T) {
	err := StringEnum("GET", "POST").Match("DELETE")
	if err == nil {
		t.Fatal("Expected error got none")
	}
	if err.Error() != "\"DELETE\" not included in [\"GET\" \"POST\"]" {
		t.Fatalf("wrong error message %q", err.Error())
	}
}
