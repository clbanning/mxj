// xml4_test.go - new features KeepNamespece / ForceList

package mxj

import (
	"fmt"
	"testing"
)

func TestKeepNamespace(t *testing.T) {
	src := `<ns:r a="2" ns:b="3"><e>1</e></ns:r>`
	FixRoot(true)
	KeepNamespace(true)
	dom, err := NewMapXml([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	var domv map[string]interface{}
	domv = dom
	xml, err := AnyXml(domv)
	if err != nil {
		t.Fatal(err)
	}
	if string(xml) != src {
		fmt.Println(string(xml), "!=", src)
		t.Fatal()
	}
}

func TestForList(t *testing.T) {
	src := `<r><e>1</e></r>`
	FixRoot(true)
	ForceList("r", "e")
	dom, err := NewMapXml([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	var domv map[string]interface{}
	domv = dom
	if r, ok := domv["r"]; ok {
		if rr, ok := r.(map[string]interface{}); ok {
			if e, ok := rr["e"]; ok {
				if ee, ok := e.([]interface{}); ok {
					if len(ee) == 1 && ee[0] == "1" {
						return
					}
				}
			}
		}
	}
	fmt.Printf("%v\n", domv)
	t.Fatalf("Invalid, must be an array")
}
