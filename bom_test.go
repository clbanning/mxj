// bomxml.go - test handling Byte-Order-Mark headers

package mxj

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestBom(t *testing.T) {
	fmt.Println("\n--------------- bom_test.go \n")
	fmt.Println("TestBom ...")
	x := boms // just grab the array

	for _, v := range x {
		if !isBOM(v) {
			t.Fatalf("isBOM returned 'false'; %x", v)
		}
	}

	// use just UTF-8 BOM ... no alternative CharSetReader
	if _, err := NewMapXml(x[0]); err != io.EOF {
		t.Fatalf("NewMapXml err;", err)
	}

	if _, err := NewMapXmlSeq(x[0]); err != io.EOF {
		t.Fatalf("NewMapXmlSeq err:", err)
	}
}

var bomdata = append(boms[0], []byte(`<Allitems>
	<Item>
	</Item>
   <Item>
        <link>http://www.something.com</link>
        <description>Some description goes here.</description>
   </Item>
</Allitems>`)...)

func TestBomData(t *testing.T) {
	fmt.Println("TestBomData ...")
	m, err := NewMapXml(bomdata)
	if err != nil {
		t.Fatalf("err: didn't find xml.StartElement")
	}
	fmt.Printf("m: %v\n", m)
	j, _ := m.Xml()
	fmt.Println("m:", string(j))
}

func TestBomDataSeq(t *testing.T) {
	fmt.Println("TestBomDataSeq ...")
	m, err := NewMapXmlSeq(bomdata)
	if err != nil {
		t.Fatalf("err: didn't find xml.StartElement")
	}
	fmt.Printf("m: %v\n", m)
	j, _ := m.XmlSeq()
	fmt.Println("m:", string(j))
}

func TestBomDataReader(t *testing.T) {
	fmt.Println("TestBomDataReader ...")
	r := bytes.NewReader(bomdata)
	m, err := NewMapXmlReader(r)
	if err != nil {
		t.Fatalf("err: didn't find xml.StartElement")
	}
	fmt.Printf("m: %v\n", m)
	j, _ := m.Xml()
	fmt.Println("m:", string(j))
}

func TestBomDataSeqReader(t *testing.T) {
	fmt.Println("TestBomDataSeqReader ...")
	r := bytes.NewReader(bomdata)
	m, err := NewMapXmlSeqReader(r)
	if err != nil {
		t.Fatalf("err: didn't find xml.StartElement")
	}
	fmt.Printf("m: %v\n", m)
	j, _ := m.XmlSeq()
	fmt.Println("m:", string(j))
}
