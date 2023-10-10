package main

import (
	"fmt"

	"github.com/clbanning/mxj/v2"
)

func main() {
	data := map[interface{}]interface{}{
		"hello": "out there",
		1:       "number one",
		3.12:    "pi",
		"five":  5,
	}

	m, err := mxj.AnyXmlIndent(data,"", "   ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(m))
}
