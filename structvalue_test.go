package mxj

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type result struct {
	XMLName xml.Name `xml:"xml"`
	Content string   `xml:"content"`
}

func TestStructValue(t *testing.T) {
	fmt.Println("----------------- structvalue_test.go ...")

	data, err := Map(map[string]interface{}{
		"data": result{Content: "content"},
	}).Xml()
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "<data><xml><content>content</content></xml></data>" {
		t.Fatal("encoding error:", string(data))
	}
}
