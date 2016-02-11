// bomxml.go - test handling Byte-Order-Mark headers

package mxj

import (
	"fmt"
	"testing"
)

func TestBom( t *testing.T) {
	fmt.Println("\n--------------- bom_test.go \n")
	fmt.Println("TestBom ... \n")
	x := boms // just grab the array

	for _, v := range x {
		if !isBOM(v) {
			t.Fatalf("isBOM returned 'false'; %x", v)
		}
	}

	if _, err := NewMapXml(x[0]); err != IsBOM {
		t.Fatalf("NewMapXml err;", err)
	}

	if _, err := NewMapXmlSeq(x[0]); err != IsBOM {
		t.Fatalf("NewMapXmlSeq err:", err)
	}
}

