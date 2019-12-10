package mxj

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	fmt.Println("------------ exists_test.go")
	m := map[string]interface{}{
		"Div": map[string]interface{}{
			"Colour": "blue",
		},
	}
	mv := Map(m)

	if v, _ := mv.Exists("Div.Colour"); !v {
		t.Fatal("Haven't found an existing element")
	}

	if v, _ := mv.Exists("Div.Color"); v {
		t.Fatal("Have found a non existing element")
	}
}

/*
var existsDoc = []byte(`
<doc>
   <books>
      <book seq="1">
         <author>William T. Gaddis</author>
         <title>The Recognitions</title>
         <review>One of the great seminal American novels of the 20th century.</review>
      </book>
   </books>
	<book>Something else.</book>
</doc>
`)

func TestExistsWithSubKeys(t *testing.T) {
	mv, err := NewMapXml(existsDoc)
	if err != nil {
		t.Fatal("err:", err.Error())
	}

	if !mv.Exists("doc.books.book", "-seq:1") {
		t.Fatal("Did't find an existing element")
	}

	if mv.Exists("doc.books.book", "-seq:2") {
		t.Fatal("Found a non-existing element")
	}
}
*/
