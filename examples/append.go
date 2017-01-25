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

	b, err = appendElement(b, 2)
	if err != nil {
		fmt.Println("append err:", err)
		return
	}

	// create the new value for 'b' as a map
	// and update 'm'
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

func appendElement(v interface{}, n int) (interface{}, error) {
	switch v.(type) {
	case string:
		v = []interface{}{v.(string), strconv.Itoa(n)}
	case []interface{}:
		v = append(v.([]interface{}), interface{}(strconv.Itoa(n)))
	default:
		// capture map[string]interface{} value, simple element, etc.
		return v, fmt.Errorf("invalid type")
	}
	return v, nil
}
