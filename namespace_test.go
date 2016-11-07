package mxj

import (
	"fmt"
	"testing"
)

func TestNamespaceHeader(t *testing.T) {
	fmt.Println("\n---------------- namespace_test.go ...\n")
}

func TestBeautifyXml(t *testing.T) {
	fmt.Println("\n----------------  TestBeautifyXml ...")
	v, err := BeautifyXml([]byte(flatxml), "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(flatxml)
	fmt.Println(string(v))
}
