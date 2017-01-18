package mxj

import (
	"fmt"
	"testing"
)

func TestStakeCase(t *testing.T) {
	PrependAttrWithHyphen(true)
	fmt.Println("\n----------- TestSnakeCase")
	CoerceKeysToSnakeCase()
	defer CoerceKeysToSnakeCase()

	data1 := `<xml-rpc><element-one attr-1="an attribute">something</element-one></xml-rpc>`
	data2 := `<xml_rpc><element_one attr_1="an attribute">something</element_one></xml_rpc>`

	m, err := NewMapXml([]byte(data1))
	if err != nil {
		t.Fatal(err)
	}

	x, err := m.Xml()
	if err != nil {
		t.Fatal(err)
	}
	if string(x) != data2 {
		t.Fatal(string(x), "!=", data2)
	}
}

