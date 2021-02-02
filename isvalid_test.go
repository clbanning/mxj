package mxj

import (
	"fmt"
	"testing"
)

func TestXmlCheckIsValid(t *testing.T) {
	fmt.Println("================== TestXmlCheckIsValid")

	data := []byte(`{"":"empty", "$invalid":"hex$", "entities":"<>&", "nil": null}`)
	m, err := NewMapJson(data)
	if err != nil {
		t.Fatal("NewMapJson err;", err)
	}
	fmt.Printf("%v\n", m)

	XmlCheckIsValid()
	defer XmlCheckIsValid()
	if _, err = m.Xml(); err == nil {
		t.Fatal("Xml err: nil")
	}

	if _, err = m.XmlIndent("", "   "); err == nil {
		t.Fatal("XmlIndent err: nil")
	}

	ms, err := NewMapXmlSeq([]byte(`<doc><elem>element</elem></doc>`))
	fmt.Printf("%v\n", ms)
	if _, err = ms.Xml(); err != nil {
		t.Fatal("Xml err:", err)
	}

	if _, err = ms.XmlIndent("", "   "); err != nil {
		t.Fatal("XmlIndent err:", err)
	}
}
