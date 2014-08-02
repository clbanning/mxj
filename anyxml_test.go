package mxj

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAnyXmlHeader(t *testing.T) {
	fmt.Println("\n----------------  anyxml_test.go ...\n")
}

var anydata = []byte(`[
    {
        "somekey": "somevalue"
    },
    {
        "somekey": "somevalue"
    },
    {
        "somekey": "somevalue"
    }
]`)

func TestAnyXml(t *testing.T) {
	a := make([]interface{},3)
	err := json.Unmarshal(anydata,&a)
	x, err := AnyXml(a)
	if err != nil {
		t.Fatal(err)
	}
   fmt.Println("[]->x:", string(x))

	f := interface{}(3.14159625)
	x, err = AnyXml(f)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("f->x:", string(x))
}

func TestAnyXmlIndent(t *testing.T) {
	a := make([]interface{},3)
	err := json.Unmarshal(anydata,&a)
	x, err := AnyXmlIndent(a, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
   fmt.Println("[]->x:\n", string(x))

	f := interface{}(3.14159625)
	x, err = AnyXmlIndent(f, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("f->x:\n", string(x))
}
