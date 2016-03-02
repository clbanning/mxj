// nan_test.go

package mxj

import (
	"fmt"
	"testing"
)

func TestNan(t *testing.T) {
	fmt.Println("\n------------ TestNan\n")

	data := []byte("<foo><bar>NAN</bar></foo>")
	m, err := NewMapXml(data)
	if err != nil {
		t.Fatal("err:", err)
	}
	v, err := m.ValueForPath("foo.bar")
	if err != nil {
		t.Fatal("err:", err)
	}
	if _, ok := v.(string); !ok {
		t.Fatal("v not string")
	}
	fmt.Println("foo.bar:", v)
}
