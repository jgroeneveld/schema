package schema

import "testing"

func TestCapture_Simple(t *testing.T) {
	data := dataFromJSON(t, `{
			"id": 12,
			"messages": [
				{
					"sender_id": 12,
					"text": "hi"
				},
				{
					"sender_id": 12,
					"text": "Hello World"
				},
				{
					"sender_id": 12,
					"text": "Moin Moin"
				}
			]
		}`)

	userID := Capture("user_id")
	err := Match(Map{
		"id": userID,
		"messages": ArrayEach(Map{
			"sender_id": userID,
			"text":      IsString,
		}),
	}, data)

	if err != nil {
		t.Fatal(err)
	}

	if !userID.Equals(12) {
		t.Fatalf("userID is not 12 but %v:%T", userID.CapturedValue(), userID.CapturedValue())
	}
}

func TestCapture(t *testing.T) {
	data := dataFromJSON(t, `{
			"id": 12,
			"relationship": [42, 12],
			"messages": [{"sender_id": 12, "text": "hi"}]
		}`)

	userID := Capture("user_id")
	err := Match(Map{
		"id":           userID,
		"relationship": ArrayUnordered(userID, IsInteger),
		"messages": Array(Map{
			"sender_id": userID,
			"text":      IsString,
		}),
	}, data)

	if err != nil {
		t.Fatal(err)
	}

	if !userID.Equals(12) {
		t.Fatalf("userID is not 12 but %v:%T", userID.CapturedValue(), userID.CapturedValue())
	}
}
