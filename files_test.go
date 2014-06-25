package mxj

import (
	"fmt"
	"testing"
)

func TestReadFileHeader(t *testing.T) {
	fmt.Println("\n----------------  files_test.go ...\n")
}

func TestReadJsonFile(t *testing.T) {
	fmt.Println("ReadMapsFromJsonFile()")
	am, err := ReadMapsFromJsonFile("files_test.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range am {
		fmt.Printf("%v\n", v)
	}

	am, err = ReadMapsFromJsonFile("nil")
	if err == nil {
		t.Fatal("no error returned for read of nil file")
	}
	fmt.Println("caught error: ", err.Error())

	am, err = ReadMapsFromJsonFile("files_test.badjson")
	if err == nil {
		t.Fatal("no error returned for read of badjson file")
	}
	fmt.Println("caught error: ", err.Error())
}

func TestReadJsonFileRaw(t *testing.T) {
	fmt.Println("ReadMapsFromJsonFileRaw()")
	mr, err := ReadMapsFromJsonFileRaw("files_test.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range mr {
		fmt.Printf("%v\n", v)
	}

	mr, err = ReadMapsFromJsonFileRaw("nil")
	if err == nil {
		t.Fatal("no error returned for read of nil file")
	}
	fmt.Println("caught error: ", err.Error())

	mr, err = ReadMapsFromJsonFileRaw("files_test.badjson")
	if err == nil {
		t.Fatal("no error returned for read of badjson file")
	}
	fmt.Println("caught error: ", err.Error())
}

func TestReadXmFile(t *testing.T) {
	fmt.Println("ReadMapsFromXmlFile()")
	am, err := ReadMapsFromXmlFile("files_test.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range am {
		fmt.Printf("%v\n", v)
	}

	am, err = ReadMapsFromXmlFile("nil")
	if err == nil {
		t.Fatal("no error returned for read of nil file")
	}
	fmt.Println("caught error: ", err.Error())

	am, err = ReadMapsFromXmlFile("files_test.badxml")
	if err == nil {
		t.Fatal("no error returned for read of badjson file")
	}
	fmt.Println("caught error: ", err.Error())
}

func TestReadXmFileRaw(t *testing.T) {
	fmt.Println("ReadMapsFromXmlFileRaw()")
	mr, err := ReadMapsFromXmlFileRaw("files_test.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range mr {
		fmt.Printf("%v\n", v)
	}

	mr, err = ReadMapsFromXmlFileRaw("nil")
	if err == nil {
		t.Fatal("no error returned for read of nil file")
	}
	fmt.Println("caught error: ", err.Error())

	mr, err = ReadMapsFromXmlFileRaw("files_test.badxml")
	if err == nil {
		t.Fatal("no error returned for read of badjson file")
	}
	fmt.Println("caught error: ", err.Error())
}

func TestMaps(t *testing.T) {
	fmt.Println("TestMaps()")
	mvs := NewMaps()
	for i := 0 ; i < 2 ; i++ {
		m, _ := NewMapJson([]byte(`{ "this":"is", "a":"test" }`))
		mvs = append(mvs, m)
	}
	fmt.Println("mvs:", mvs)

	s, _ := mvs.JsonString()
	fmt.Println("JsonString():", s)

	s, _ = mvs.JsonStringIndent("", "  ")
	fmt.Println("JsonStringIndent():", s)

	s, _ = mvs.XmlString()
	fmt.Println("XmlString():", s)

	s, _ = mvs.XmlStringIndent("", "  ")
	fmt.Println("XmlStringIndent():", s)
}
