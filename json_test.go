package mxj

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

var jdata = []byte(`{ "key1":"string", "key2":34, "key3":true, "key4":"unsafe: <>" }`)
var jdata2 = []byte(`{ "key1":"string", "key2":34, "key3":true, "key4":"unsafe: <>" },
	{ "key":"value in new JSON string" }`)

func TestJsonHeader(t *testing.T) {
	fmt.Println("\n----------------  json_test.go ...\n")
}

func TestNewMapJson(t *testing.T) {

	m, merr := NewMapJson(jdata)
	if merr != nil {
		t.Fatal("NewMapJson, merr:", merr.Error())
	}

	fmt.Println("NewMapJson, jdata:", string(jdata))
	fmt.Println("NewMapJson, m    :", m)
}

func TestNewMapJsonError(t *testing.T) {

	m, merr := NewMapJson(jdata[:len(jdata)-2])
	if merr == nil {
		t.Fatal("NewMapJsonError, m:", m)
	}

	fmt.Println("NewMapJsonError, jdata :", string(jdata[:len(jdata)-2]))
	fmt.Println("NewMapJsonError, merror:", merr.Error())

	newData := []byte(`{ "this":"is", "in":error }`)
	m, merr = NewMapJson(newData)
	if merr == nil {
		t.Fatal("NewMapJsonError, m:", m)
	}

	fmt.Println("NewMapJsonError, newData :", string(newData))
	fmt.Println("NewMapJsonError, merror  :", merr.Error())
}

func TestNewMapJsonReader(t *testing.T) {

	rdr := bytes.NewBuffer(jdata2)

	for {
		m, jb, merr := NewMapJsonReaderRaw(rdr)
		if merr != nil && merr != io.EOF {
			t.Fatal("NewMapJsonReader, merr:", merr.Error())
		}
		if merr == io.EOF {
			break
		}

		fmt.Println("NewMapJsonReader, jb:", string(*jb))
		fmt.Println("NewMapJsonReader, m :", m)
	}
}

func TestJson(t *testing.T) {

	m, _ := NewMapJson(jdata)

	j, jerr := m.Json()
	if jerr != nil {
		t.Fatal("Json, jerr:", jerr.Error())
	}

	fmt.Println("Json, jdata:", string(jdata))
	fmt.Println("Json, j    :", string(j))
}

func TestJsonWriter(t *testing.T) {
	mv := Map( map[string]interface{}{ "this":"is a", "float":3.14159, "and":"a", "bool":true } )

	w := new(bytes.Buffer)
	raw, err := mv.JsonWriterRaw(w)
	if err != nil {
		t.Fatal("err:",err.Error())
	}

	b := make([]byte,w.Len())
	_, err = w.Read(b)
	if err != nil {
		t.Fatal("err:",err.Error())
	}

	fmt.Println("JsonWriter, raw:", string(*raw))
	fmt.Println("JsonWriter, b  :", string(b))
}

// --------------------------  JSON Handler test cases -------------------------

/* tested in bulk_test.go ...
var jhdata = []byte(`{ "string":"this is it", "number":4, "boolean":true },
	{ "some":"thing", "that":[ "is", "more", "complex", { "like":"this", "value":"here" } ] },
	{ "this":"has", "an":error }`)

func TestHandleJsonReader(t *testing.T) {
	fmt.Println("HandleJsonReader:", string(jhdata))

	rdr := bytes.NewReader(jhdata)
	err := HandleJsonReader(rdr, jmhandler, jehandler)
	if err != nil {
		t.Fatal("err:", err.Error())
	}
}

var jt *testing.T

func jmhandler(m Map, raw *[]byte) bool {
	j, jerr := m.Json()
	if jerr != nil {
		jt.Fatal("... jmhandler:", jerr.Error())
		return false
	}

	fmt.Println("... jmhandler, raw:", string(*raw))
	fmt.Println("... jmhandler, j  :", string(j))
	return true
}

func jehandler(err error, raw *[]byte) bool {
	if err == nil {
		// shouldn't be here
		jt.Fatal("... jehandler: <nil>")
		return false
	}

	fmt.Println("... jehandler, err:", err.Error())
	fmt.Println("... jehandler: raw", string(*raw))
	return true
}
*/
