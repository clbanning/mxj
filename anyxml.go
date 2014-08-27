package mxj

import (
	"encoding/xml"
	"reflect"
)

// Encode arbitrary value as XML. 
// 
// Note: unmarshaling the resultant
// XML may not return the original value, since tag labels may have been injected
// to create the XML representation of the value.
/*
 Encode an arbitrary JSON object.
	package main
	
	import (
		"encoding/json"
		"fmt"
		"github/clbanning/mxj"
	)
	
	func main() {
		jasondata := []byte(`[
			{ "somekey":"somevalue" },
			"string",
			3.14159265,
			true
		]`)
		var i interface{}
		err := json.Unmarshal(jsondata, &i)
		if err != nil {
			// do something
		}
		x, err := anyxml.XmlIndent(i, "", "  ", "mydoc")
		if err != nil {
			// do something else
		}
		fmt.Println(string(x))
	}
	
	output:
		<mydoc>
		  <somekey>somevalue</somekey>
		  <element>string</element>
		  <element>3.14159265</element>
		  <element>true</element>
		</mydoc>
*/
func AnyXml(v interface{}, rootTag ...string) ([]byte, error) {
	if reflect.TypeOf(v).Kind() == reflect.Struct {
		return xml.Marshal(v)
	}

	var err error
	s := new(string)
	p := new(pretty)

	var rt string
	if len(rootTag) == 1 {
		rt = rootTag[0]
	} else {
		rt = DefaultRootTag
	}

	var ss string
	var b []byte
	switch v.(type) {
	case []interface{}:
		ss = "<" + rt + ">"
		for _, vv := range v.([]interface{}) {
			switch vv.(type) {
			case map[string]interface{}:
				m := vv.(map[string]interface{})
				if len(m) == 1 {
					for tag, val := range m {
						err = mapToXmlIndent(false, s, tag, val, p)
					}
				} else {
					err = mapToXmlIndent(false, s, "element", vv, p)
				}
			default:
				err = mapToXmlIndent(false, s, "element", vv, p)
			}
			if err != nil {
				break
			}
		}
		ss += *s + "</" + rt + ">"
		b = []byte(ss)
	case map[string]interface{}:
		m := Map(v.(map[string]interface{}))
		b, err = m.Xml(rootTag...)
	default:
		err = mapToXmlIndent(false, s, rt, v, p)
		b = []byte(*s)
	}

	return b, err
}


// Encode an arbitrary value as a pretty XML string.
func AnyXmlIndent(v interface{}, prefix, indent string, rootTag ...string) ([]byte, error) {
	if reflect.TypeOf(v).Kind() == reflect.Struct {
		return xml.MarshalIndent(v, prefix, indent)
	}

	var err error
	s := new(string)
	p := new(pretty)
	p.indent = indent
	p.padding = prefix

	var rt string
	if len(rootTag) == 1 {
		rt = rootTag[0]
	} else {
		rt = DefaultRootTag
	}

	var ss string
	var b []byte
	switch v.(type) {
	case []interface{}:
		ss = "<" + rt + ">\n"
		p.Indent()
		for _, vv := range v.([]interface{}) {
			switch vv.(type) {
			case map[string]interface{}:
				m := vv.(map[string]interface{})
				if len(m) == 1 {
					for tag, val := range m {
						err = mapToXmlIndent(true, s, tag, val, p)
					}
				} else {
					p.start = 1 // we 1 tag in
					err = mapToXmlIndent(true, s, "element", vv, p)
					*s += "\n"
				}
			default:
				p.start = 0 // in case trailing p.start = 1
				err = mapToXmlIndent(true, s, "element", vv, p)
			}
			if err != nil {
				break
			}
		}
		ss += *s + "</" + rt + ">"
		b = []byte(ss)
	case map[string]interface{}:
		m := Map(v.(map[string]interface{}))
		b, err = m.XmlIndent(prefix, indent, rootTag...)
	default:
		err = mapToXmlIndent(true, s, rt, v, p)
		b = []byte(*s)
	}

	return b, err
}
