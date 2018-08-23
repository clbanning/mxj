// from: https://www.reddit.com/r/golang/comments/99lnxd/help_needed_encodingxml_parsing_malformed_xml/
//

package main

import (
	"encoding/xml"
	"fmt"

	"github.com/clbanning/mxj"
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
	mxj.SetAttrPrefix("")
	m, err := mxj.NewMapXml(data)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("%v\n", m)

	val := decodedXml{
		XMLName:  xml.Name{"","tag"},
		SomeAttr: m["tag"].(map[string]interface{})["someattr"].(string),
	}
	fmt.Printf("%v\n", val)
}
