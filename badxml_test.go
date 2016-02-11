// trying to recreate a panic

package mxj

import (
	"fmt"
	"testing"
)

var baddata = []byte(`
	something strange
<Allitems>
	<Item>
	</Item>
   <Item>
        <link>http://www.something.com</link>
        <description>Some description goes here.</description>
   </Item>
</Allitems>
`)

func TestBadXml(t *testing.T) {
	fmt.Println("\n---------------- badxml_test.go ...\n")
	fmt.Println("TestBadXml ...")
	_, err := NewMapXml(baddata)
	if err == nil {
		t.Fatalf("no err on baddata")
	}
	fmt.Println("ok:", err)
}

func TestBadXmlSeq(t *testing.T) {
	fmt.Println("TestBadXmlSeq ...")
	_, err := NewMapXmlSeq(baddata)
	if err == nil {
		t.Fatalf("no err on baddata")
	}
	fmt.Println("ok:", err)
}
