package web

import "testing"

func TestGetResult(t *testing.T) {

	sample := `{
    "fname":"John",
    "sname":"Alice",
    "percentage":"46",
    "result":"Can choose someone better."
    }
    `
	w, err := GetResult(sample)

	if err != nil {
		t.Error(err)
	}

	t.Log(w)
	if w.Name1 != "John" {
		t.Log("Name one is not as expected! >:( )")
		t.Fail()
	}

}
