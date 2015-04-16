package mxj

import (
	"testing"
)

func TestSetValueForPath(t *testing.T) {
	m := map[string]interface{}{
		"Div": map[string]interface{}{
			"Colour": "blue",
			"Font": map[string]interface{}{
				"Family": "sans",
			},
		},
	}
	mv := Map(m)
	err := mv.SetValueForPath("Big", "Div.Font.Size")

	if err != nil {
		t.Fatal(err)
	}

	values, err := mv.ValuesForPath("Div.Font.Size")
	if len(values) == 0 {
		t.Fatal("err: didn't add the new key")
	}
	if values[0] != "Big" {
		t.Fatal("err: value is different")
	}

	err = mv.SetValueForPath("red", "Div.Colour")
	if err != nil {
		t.Fatal(err)
	}
	values, err = mv.ValuesForPath("Div.Colour")
	if len(values) == 0 {
		t.Fatal("err: existing key is deleted")
	}
	if values[0] != "red" {
		t.Fatal("err: existig key value is different")
	}
}
