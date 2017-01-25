// Per https://github.com/clbanning/mxj/issues/34

package main

import (
	"fmt"
	"strconv"

	"github.com/clbanning/mxj"
)

var data = []byte(`
<a>
  <b>1</b>
</a>`)

func main() {
	m, err := mxj.NewMapXml(data)
	if err != nil {
		fmt.Println("new  err:", err)
		return
	}
	b, err := m.ValueForPath("a.b")
	if err != nil {
		fmt.Println("value err:", err)
		return
	}

	switch b.(type) {
	case string:
		b = []interface{}{b.(string), strconv.Itoa(2)}
	case []interface{}:
		b = append(b.([]interface{}), interface{}(strconv.Itoa(2)))
	default:
		// capture map[string]interface{} value
		fmt.Println("err: invalid type for a.b")
		return
	}

	val := map[string]interface{}{"b": b}
	n, err := m.UpdateValuesForPath(val, "a.b")
	if err != nil {
		fmt.Println("update err:", err)
		return
	}
	if n != 1 {
		fmt.Println("err: a.b not updated, n =", n)
		return
	}

	x, err := m.XmlIndent("", "  ")
	fmt.Println(string(x))
}
