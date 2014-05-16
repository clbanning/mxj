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
