// from: https://www.reddit.com/r/golang/comments/99lnxd/help_needed_encodingxml_parsing_malformed_xml/
//

package main

import (
	"encoding/xml"
	"fmt"

	"github.com/clbanning/mxj/v2"
)

func main() {
	type decodedXml struct {
		XMLName  xml.Name `xml:"tag"          binding:"required"`
		SomeAttr string   `xml:"someattr,attr"`
	}

	data := []byte(`
-==sddsfdqsdqsdsqd
A BUNCH OF INVALID DATA IN THE XML FILE HAHAH
--SEPARATERUFCIFHEIR
Content-Type: text/plain

<tag someattr="myvalue"></tag>
--SEPARATERUFCIFHEIR--
{"hello": "world"}
		`)

	// don't prepend attributes with '-'
	mxj.SetAttrPrefix("")

	// parse the data as a map[string]interface{} value
	m, err := mxj.NewMapXml(data)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// check that we got a 'tag' tagged doc
	if _, ok := m["tag"]; !ok {
		fmt.Println("no tag doc ...")
		return
	}
	fmt.Printf("%v\n", m)

	// extract the attribute value
	attrval, err := m.ValueForPath("tag.someattr")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// create decodeXml value
	val := decodedXml{
		XMLName:  xml.Name{"", "tag"},
		SomeAttr: attrval.(string)}
	fmt.Printf("%v\n", val)
}
