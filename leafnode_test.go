package mxj

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestLNHeader(t *testing.T) {
	fmt.Println("\n----------------  leafnode_test.go ...")
}

func TestLeafNodes(t *testing.T) {
	json1 := []byte(`{
		"friends": [
			{
				"skills": [
					44, 12
				]
			}
		]
	}`)

	json2 := []byte(`{
		"friends":
			{
				"skills": [
					44, 12
				]
			}

	}`)

	m, _ := NewMapJson(json1)
	ln := m.LeafNodes()
	fmt.Println("\njson1-LeafNodes:")
	for _, v := range ln {
		fmt.Printf("%#v\n", v)
	}
	p := m.LeafPaths()
	fmt.Println("\njson1-LeafPaths:")
	for _, v := range p {
		fmt.Printf("%#v\n", v)
	}

	m, _ = NewMapJson(json2)
	ln = m.LeafNodes()
	fmt.Println("\njson2-LeafNodes:")
	for _, v := range ln {
		fmt.Printf("%#v\n", v)
	}
	v := m.LeafValues()
	fmt.Println("\njson2-LeafValues:")
	for _, v := range v {
		fmt.Printf("%#v\n", v)
	}

	fmt.Println("\nxmldata-LeafNodes:")
	r := bytes.NewReader(xmldata)
	for {
		m, err := NewMapXmlReader(r)
		if err == io.EOF {
			break
		}
		ln = m.LeafNodes()
		for _, v := range ln {
			fmt.Printf("%#v\n", v)
		}
	}
}
