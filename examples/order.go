// Preserve list order with intermixed sub-elements.
// from: https://groups.google.com/forum/#!topic/golang-nuts/8KvlKsdh84k

package main

import (
	"fmt"
	"sort"

	"github.com/clbanning/mxj"
)

var data = `<node>
  <a>sadasd</a>
  <b>gfdfg</b>
  <a>hihihi</a>
  <b>jkjkjk</b>
</node>`

func main() {
	m, err := mxj.NewMapXmlSeq([]byte(data))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// Merge a and b members into a single list
	// that we can work with.
	list := make([]*listval, 0)
	for k, v := range m["node"].(map[string]interface{}) {
		for _, vv := range v.([]interface{}) {
			mem := vv.(map[string]interface{})
			lval := &listval{k, mem["#text"].(string), mem["#seq"].(int)}
			list = append(list, lval)
		}
	}

	// sort the list into orignal DOC sequence
	sort.Sort(Lval(list))

	// do some work with the list members - let's swap values
	for i := 0; i < 2; i++ {
		list[i].val, list[3-i].val = list[3-i].val, list[i].val
	}

	// rebuild map[string]interface{} value for "node"
	a := make([]interface{}, 0)
	b := make([]interface{}, 0)
	for _, v := range list {
		val := map[string]interface{}{"#text": v.val, "#seq": v.seq}
		switch v.list {
		case "a":
			a = append(a, interface{}(val))
		case "b":
			b = append(b, interface{}(val))
		}
	}
	val := map[string]interface{}{"a": a, "b": b}
	m["node"] = interface{}(val)

	x, err := m.XmlSeqIndent("", "  ")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(data)      // original
	fmt.Println(string(x)) // modified
}

// ======== sort interface implementation ========
type listval struct {
	list string
	val  string
	seq  int
}

type Lval []*listval

func (l Lval) Len() int {
	return len(l)
}

func (l Lval) Less(i, j int) bool {
	if l[i].seq < l[j].seq {
		return true
	}
	return false
}

func (l Lval) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
