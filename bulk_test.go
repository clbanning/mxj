// bulk_test.go - uses Handler and Writer functions to process some streams as a demo.

package mxj

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBulkHeader(t *testing.T) {
	fmt.Println("\n----------------  bulk_test.go ...\n")
}

var xmldata = []byte(`
<book>
	<author>William H. Gaddis</author>
	<title>The Recognitions</title>
	<review>One of the great seminal American novels of the 20th century.</review>
</book>
<book>
	<author>William H. Gaddis</author>
	<title>JR</title>
	<review>Won the National Book Award.</end_tag_error>
</book>
<book>
	<author>Austin Tappan Wright</author>
	<title>Islandia</title>
	<review>An example of earlier 20th century American utopian fiction.</review>
</book>
<book>
	<author>John Hawkes</author>
	<title>The Beetle Leg</title>
	<review>A lyrical novel about the construction of Ft. Peck Dam in Montana.</review>
</book>
<book>
	<author>
		<first_name>T.E.</first_name>
		<last_name>Porter</last_name>
	</author>
	<title>King's Day</title>
	<review>A magical novella.</review>
</book>`)

var jsondata = []byte(`
 {"book":{"author":"William H. Gaddis","review":"One of the great seminal American novels of the 20th century.","title":"The Recognitions"}}
{"book":{"author":"Austin Tappan Wright","review":"An example of earlier 20th century American utopian fiction.","title":"Islandia"}}
{"book":{"author":"John Hawkes","review":"A lyrical novel about the construction of Ft. Peck Dam in Montana.","title":"The Beetle Leg"}}
{"book":{"author":{"first_name":"T.E.","last_name":"Porter"},"review":"A magical novella.","title":"King's Day"}}
{ "here":"we", "put":"in", "an":error }`)

var jsonWriter = new(bytes.Buffer)
var xmlWriter = new(bytes.Buffer)

var jsonErrLog = new(bytes.Buffer)
var xmlErrLog = new(bytes.Buffer)

func TestXmlReader(t *testing.T) {
	// create Reader for xmldata
	xmlReader := bytes.NewReader(xmldata)

	// read XML from Readerand pass Map value with the raw XML to handler
	err := HandleXmlReader(xmlReader, bxmaphandler, bxerrhandler)
	if err != nil {
		t.Fatal("err:", err.Error())
	}

	// get the JSON
	j := make([]byte, jsonWriter.Len())
	_, _ = jsonWriter.Read(j)

	// get the errors
	e := make([]byte, xmlErrLog.Len())
	_, _ = xmlErrLog.Read(e)

	// print the input
	fmt.Println("XmlReader, xmldata:\n", string(xmldata))
	// print the result
	fmt.Println("XmlReader, result :\n", string(j))
	// print the errors
	fmt.Println("XmlReader, errors :\n", string(e))
}

func bxmaphandler(m Map) bool {
	j, err := m.JsonIndent("", "  ", true)
	if err != nil {
		return false
	}

	_, _ = jsonWriter.Write(j)
	// put in a NL to pretty up printing the Writer
	_, _ = jsonWriter.Write([]byte("\n"))
	return true
}

func bxerrhandler(err error) bool {
	// write errors to file
	_, _ = xmlErrLog.Write([]byte(err.Error())) 
	// _, _ = xmlErrWriter.Write(*raw) // used for HandleXmlReader, no raw XML passed
	 _, _ = xmlErrLog.Write([]byte("\n")) // pretty up
	return true
}

func TestJsonReader(t *testing.T) {
	jsonReader := bytes.NewReader(jsondata)

	// read all the JSON
	err := HandleJsonReader(jsonReader, bjmaphandler, bjerrhandler)
	if err != nil {
		t.Fatal("err:", err.Error())
	}

	// get the XML
	x := make([]byte, xmlWriter.Len())
	_, _ = xmlWriter.Read(x)

	// get the errors
	e := make([]byte, jsonErrLog.Len())
	_, _ = jsonErrLog.Read(e)

	// print the input
	fmt.Println("JsonReader, jsondata:\n", string(jsondata))
	// print the result
	fmt.Println("JsonReader, result  :\n", string(x))
	// print the errors
	fmt.Println("JsonReader, errors :\n", string(e))
}

func bjmaphandler(m Map) bool {
	x, err := m.XmlIndent("  ", "  ")
	if err != nil {
		return false
	}
	_, _ = xmlWriter.Write(x)
	// put in a NL to pretty up printing the Writer
	_, _ = xmlWriter.Write([]byte("\n"))
	return true
}

func bjerrhandler(err error) bool {
	// write errors to file
	_, _ = jsonErrLog.Write([]byte(err.Error()))
	_, _ = jsonErrLog.Write([]byte("\n")) // pretty up
	return true
}
