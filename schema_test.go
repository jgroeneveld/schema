package schema

import (
	"testing"
)

func TestFullSuccess(t *testing.T) {
	data := dataFromJSON(t, `
	{
		"id": 12,
		"name": "Max Mustermann",
		"age": 42,
		"footsize": "unknown",
		"admin": false,
		"height": 1.91,
		"address": {
			"street": "Musterstrasse 12",
			"zip": "12345",
			"country": "Germany"
		},
		"tags": ["blue","red"],
		"friends": ["hans", "peter", "harald", "gundel"],
		"pi": [3,1,4,1,5,9,2,6,5,3,5,9]
	}`)

	err := Map{
		"id":       IsInteger,
		"name":     "Max Mustermann",
		"age":      42,
		"admin":    IsBool,
		"height":   IsFloat,
		"footsize": IsPresent,
		"address": Map{
			"street":  IsString,
			"zip":     IsString,
			"country": IsString,
		},
		"tags":    Array("blue", "red"),
		"friends": ArrayIncluding("hans"),
		"pi":      ArrayEach(IsInteger),
	}.Check(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestNestingFailures(t *testing.T) {
	data := dataFromJSON(t, `
		{
			"name": "-",
			"address": {
				"street": "-",
				"geo": {"lat":"-", "extra": "-"}
			},
			"tags": ["-"]
		}`)

	err := Map{
		"name": "Max Mustermann",
		"address": Map{
			"street": "Bauernhof",
			"geo": Map{
				"lat": "12",
			},
		},
		"tags": Array("blue"),
	}.Check(data)

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Errors["name"] != `"-" != "Max Mustermann"` {
		t.Errorf(`wrong error on "name": %s`, err.Errors["name"])
	}
	if err.Errors["address.street"] != `"-" != "Bauernhof"` {
		t.Errorf(`wrong error on "address.street": %s`, err.Errors["address.street"])
	}
	if err.Errors["address.geo"] != `Found extra keys: "extra"` {
		t.Errorf(`wrong error on "address.geo": %s`, err.Errors["address.geo"])
	}
	if err.Errors["address.geo.lat"] != `"-" != "12"` {
		t.Errorf(`wrong error on "address.geo.lat": %s`, err.Errors["address.geo.lat"])
	}
	if err.Errors["tags.0"] != `"-" != "blue"` {
		t.Errorf(`wrong error on "tags.0": %s`, err.Errors["tags.0"])
	}
}
