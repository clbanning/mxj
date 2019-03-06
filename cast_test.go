package mxj

import (
	"fmt"
	"testing"
)

var data = []byte(`<doc>
	<string>string</string>
	<float>3.14159625</float>
	<int>2019</int>
	<bool>
		<true>true</true>
		<false>FALSE</false>
		<T>T</T>
		<f>f</f>
	</bool></doc>`)

func TestHeader(t *testing.T) {
	fmt.Println("\ncast_test.go ----------")
}

func TestCastDefault(t *testing.T) {
	fmt.Println("------------ TestCastDefault ...")
	m, err := NewMapXml(data)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%#v\n", m)
}

func TestCastTrue(t *testing.T) {
	fmt.Println("------------ TestCastTrue ...")
	m, _ := NewMapXml(data, true)
	fmt.Printf("%#v\n", m)
}


