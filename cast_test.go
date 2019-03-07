package mxj

import (
	"fmt"
	"testing"
)

var castdata = []byte(`<doc>
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
	m, err := NewMapXml(castdata)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%#v\n", m)
}

func TestCastTrue(t *testing.T) {
	fmt.Println("------------ TestCastTrue ...")
	m, _ := NewMapXml(castdata, true)
	fmt.Printf("%#v\n", m)
}

func TestSetCheckTagToSkipFunc(t *testing.T) {
	fmt.Println("------------ TestSetCheckTagToSkipFunc ...")
	fn := func(tag string) bool {
		list := []string{"int","false"}
		for _, v := range list {
			if v == tag {
				return true
			}
		}
		return false
	}
	SetCheckTagToSkipFunc(fn)

	m, err := NewMapXml(castdata, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%#v\n", m)
}

func TestCastValuesToFloat(t *testing.T) {
	fmt.Println("------------ TestCastValuesToFloat(false) ...")
	CastValuesToFloat(false)
	defer CastValuesToFloat(true)

	m, err := NewMapXml(castdata, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%#v\n", m)
}

func TestCastValuesToBool(t *testing.T) {
	fmt.Println("------------ TestCastValuesToBool(false) ...")
	CastValuesToBool(false)
	defer CastValuesToBool(true)

	m, err := NewMapXml(castdata, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%#v\n", m)
}
