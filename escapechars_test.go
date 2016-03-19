package mxj

import (
	"fmt"
	"testing"
)

var s = `"'<>&`

func TestEscapeChars(t *testing.T) {
	fmt.Println("\n================== TestEscapeChars")

	ss := escapeChars(s)

	if ss != `&quot;&apos;&lt;&gt;&amp;` {
		t.Fatal(s, ":", ss)
	}

	fmt.Println(" s:", s)
	fmt.Println("ss:", ss)
}

func TestXMLEscapeChars(t *testing.T) {
	fmt.Println("================== TestXMLEscapeChars")

	XMLEscapeChars(true)

	m := map[string]interface{}{"mychars":s}

	x, err := AnyXmlIndent(s, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("s:", string(x))

	x, err = AnyXmlIndent(m, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("m:", string(x))

	XMLEscapeChars(false)
}

