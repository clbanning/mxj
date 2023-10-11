// https://github.com/clbanning/mxj/issues/53

package main

import (
	"fmt"

	"github.com/clbanning/mxj/v2"
)

var source = []byte(`<a>
        <b>
            <c>c</c>
            <d>d</d>
        </b>
        <e>
            <x>x</x>
            <y>y</y>
            <z>z</z>
        </e>
        <f>f</f>
        <g>g</g>
    </a>`)

var wants = map[string][]string{
	"one": {
		"a.b.c",
		"a.e.x",
	},
	"two": {
		"a.b.d",
		"a.e.y",
		"a.e.z",
	},
	"three": {
		"a.f",
		"a.g",
	},
}

func main() {
	m, err := mxj.NewMapXml(source)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	n, err := m.NewMap(wants["one"]...)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	target, err := n.XmlIndent("", "  ")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(string(target))

	// the rest of the wants
	n, _ = m.NewMap(wants["two"]...)
	target, _ = n.XmlIndent("", "  ")
	fmt.Println(string(target))

	n, _ = m.NewMap(wants["three"]...)
	target, _ = n.XmlIndent("", "  ")
	fmt.Println(string(target))
}
