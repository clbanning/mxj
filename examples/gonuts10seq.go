/* gonuts10.go - https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/tf4aDQ1Hn_c
change:
<author>
    <first-name effect_range="1999-2011">Sam</first-name>
    <first-name effect_range="2012-">Kevin</first-name>
    <last-name>Smith</last-name>
   <full-name></full-name>
</author>

to:
<author>
     <first-name effect_range="1999-2011">Sam</first-name>
    <first-name effect_range="2012-">Kevin</first-name>
    <last-name>Smith</last-name>
   <full-name>Kevin Smith</full-name>
</author>

NOTE: use NewMapXmlSeq() and mv.XmlSeqIndent() to preserve structure.

Here we build the "full-name" element value from other values in the doc by selecting the
"first-name" value with the latest dates.
*/

package main

import (
	"fmt"
	"github.com/clbanning/mxj"
)

var data = []byte(`
<author>
    <first-name effect_range="1999-2011">Sam</first-name>
    <first-name effect_range="2012-">Kevin</first-name>
    <last-name>Smith</last-name>
   <full-name></full-name>
</author>
`)

func main() {
	fmt.Println(string(data))
	m, err := mxj.NewMapXmlSeq(data)
	if err != nil {
		fmt.Println("NewMapXml err:", err)
		return
	}
	vals, err := m.ValuesForPath("author.first-name") // full-path option
	if err != nil {
		fmt.Println("ValuesForPath err:", err)
		return
	} else if len(vals) == 0 {
		fmt.Println("no first-name vals")
		return
	}
	var fname, date string
	for _, v := range vals {
		vm, ok := v.(map[string]interface{})
		if !ok {
			fmt.Println("assertion failed")
			return
		}
		fn, ok := vm["#text"].(string)
		if !ok {
			fmt.Println("no #text tag")
			return
		}
		// dt, ok := vm["#attr"].([]interface{})[0].(map[string]interface{})["effect_range"].(string)
		attri, ok := vm["#attr"].([]interface{})
		if !ok {
			fmt.Println("no #attr")
			return
		}
		/* <!-- the general case where there might be more than one attribute -->
			var dt string
			var ok bool
			for _, v := range attri {
				if dt, ok = v.(map[string]interface{})["effect_range"].(string); ok {
					break
				}
			}
			if dt == "" {
				fmt.Println("no effect_range attr k:v pair")
				return
			}
		*/
		dt, ok := attri[0].(map[string]interface{})["effect_range"].(string)
		if !ok {
			fmt.Println("no effect_range attr")
			return
		}
		if dt > date {
			date = dt
			fname = fn
		}
	}

	vals, err = m.ValuesForPath("author.last-name.#text") // full-path option
	if err != nil {
		fmt.Println("ValuesForPath err:", err)
		return
	} else if len(vals) == 0 {
		fmt.Println("no last-name vals")
		return
	}
	lname := vals[0].(string)
	if err = m.SetValueForPath(fname+" "+lname, "author.full-name.#text"); err != nil {
		fmt.Println("SetValueForPath err:", err)
		return
	}
	b, err := m.XmlSeqIndent("", "  ")
	if err != nil {
		fmt.Println("XmlIndent err:", err)
		return
	}
	fmt.Println(string(b))
}
