package mxj

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestXmlHeader(t *testing.T) {
	fmt.Println("\n----------------  xml_test.go ...\n")
}

func TestNewMapXml(t *testing.T) {
	x := []byte(`<root2><newtag>something more</newtag><list><item>1</item><item>2</item></list></root2>`)

	mv, merr := NewMapXml(x)
	if merr != nil {
		t.Fatal("merr:", merr.Error())
	}

	fmt.Println("NewMapXml, x :", string(x))
	fmt.Println("NewMapXml, mv:", mv)
}

func TestNewMapXmlError(t *testing.T) {
	x := []byte(`<root2><newtag>something more</newtag><list><item>1</item><item>2</item></list>`)

	m, merr := NewMapJson(x)
	if merr == nil {
		t.Fatal("NewMapXmlError, m:", m)
	}

	fmt.Println("NewMapXmlError, x   :", string(x))
	fmt.Println("NewMapXmlError, merr:", merr.Error())

	x = []byte(`<root2><newtag>something more</newtag><list><item>1<item>2</item></list></root2>`)
	m, merr = NewMapJson(x)
	if merr == nil {
		t.Fatal("NewMapXmlError, m:", m)
	}

	fmt.Println("NewMapXmlError, x   :", string(x))
	fmt.Println("NewMapXmlError, merr:", merr.Error())
}

func TestNewMapXmlReader(t *testing.T) {
	x := []byte(`<root><this>is a test</this></root><root2><newtag>something more</newtag><list><item>1</item><item>2</item></list></root2>`)

	r := bytes.NewReader(x)

	for {
		m, raw, err := NewMapXmlReaderRaw(r)
		if err != nil && err != io.EOF {
			t.Fatal("err:", err.Error())
		}
		if err == io.EOF && len(m) == 0 {
			break
		}
		fmt.Println("NewMapXmlReader, raw:", string(*raw))
		fmt.Println("NewMapXmlReader, m  :", m)
	}
}

// ---------------------  Xml() and XmlWriter() test cases -------------------

func TestXml(t *testing.T) {
	mv := Map{"tag1": "some data", "tag2": "more data", "boolean": true, "float": 3.14159625}

	x, err := mv.Xml()
	if err != nil {
		t.Fatal("err:", err.Error())
	}

	fmt.Println("Xml, mv:", mv)
	fmt.Println("Xml, x :", string(x))
}

func TestXmlWriter(t *testing.T) {
	mv := Map{"tag1": "some data", "tag2": "more data", "boolean": true, "float": 3.14159625}
	w := new(bytes.Buffer)

	raw, err := mv.XmlWriterRaw(w, "myRootTag")
	if err != nil {
		t.Fatal("err:",err.Error())
	}

	b := make([]byte,w.Len())
	_, err = w.Read(b)
	if err != nil {
		t.Fatal("err:", err.Error())
	}

	fmt.Println("XmlWriter, raw:", string(*raw))
	fmt.Println("XmlWriter, b  :", string(b))
}

// --------------------------  XML Handler test cases -------------------------

/* tested in bulk_test.go ...
var xhdata = []byte(`<root><this>is a test</this></root><root2><newtag>something more</newtag><list><item>1</item><item>2</item></list></root2><root3><tag></root3>`)

func TestHandleXmlReader(t *testing.T) {
	fmt.Println("HandleXmlReader:", string(xhdata))

	rdr := bytes.NewReader(xhdata)
	err := HandleXmlReader(rdr, xmhandler, xehandler)
	if err != nil {
		t.Fatal("err:", err.Error())
	}
}

var xt *testing.T

func xmhandler(m Map, raw *[]byte) bool {
	x, xerr := m.Xml()
	if xerr != nil {
		xt.Fatal("... xmhandler:", xerr.Error())
		return false
	}

	fmt.Println("... xmhandler, raw:", string(*raw))
	fmt.Println("... xmhandler, x  :", string(x))
	return true
}

func xehandler(err error, raw *[]byte) bool {
	if err == nil {
		// shouldn't be here
		xt.Fatal("... xehandler: <nil>")
		return false
	}
	if err == io.EOF {
		return true
	}

	fmt.Println("... xehandler raw:", string(*raw))
	fmt.Println("... xehandler err:", err.Error())
	return true
}
*/
