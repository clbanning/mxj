package mxj

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"testing"
)

var jdata = []byte(`{ "key1":"string", "key2":34, "key3":true, "key4":"unsafe: <>&", "key5":null }`)
var jdata2 = []byte(`{ "key1":"string", "key2":34, "key3":true, "key4":"unsafe: <>&" },
	{ "key":"value in new JSON string" }`)

func TestJsonHeader(t *testing.T) {
	fmt.Println("\n----------------  json_test.go ...")
}

func TestNewMapJson(t *testing.T) {

	m, merr := NewMapJson(jdata)
	if merr != nil {
		t.Fatal("NewMapJson, merr:", merr.Error())
	}

	fmt.Println("NewMapJson, jdata:", string(jdata))
	fmt.Printf("NewMapJson, m    : %#v\n", m)
}

func TestNewMapJsonNumber(t *testing.T) {

	JsonUseNumber = true

	m, merr := NewMapJson(jdata)
	if merr != nil {
		t.Fatal("NewMapJson, merr:", merr.Error())
	}

	fmt.Println("NewMapJson, jdata:", string(jdata))
	fmt.Printf("NewMapJson, m    : %#v\n", m)

	JsonUseNumber = false
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

		fmt.Println("NewMapJsonReader, jb:", string(jb))
		fmt.Printf("NewMapJsonReader, m : %#v\n", m)
	}
}

func TestNewMapJsonReaderNumber(t *testing.T) {

	JsonUseNumber = true

	rdr := bytes.NewBuffer(jdata2)

	for {
		m, jb, merr := NewMapJsonReaderRaw(rdr)
		if merr != nil && merr != io.EOF {
			t.Fatal("NewMapJsonReader, merr:", merr.Error())
		}
		if merr == io.EOF {
			break
		}

		fmt.Println("NewMapJsonReader, jb:", string(jb))
		fmt.Printf("NewMapJsonReader, m : %#v\n", m)
	}

	JsonUseNumber = false
}

func TestJson(t *testing.T) {

	m, _ := NewMapJson(jdata)

	j, jerr := m.Json()
	if jerr != nil {
		t.Fatal("Json, jerr:", jerr.Error())
	}

	fmt.Println("Json, jdata:", string(jdata))
	fmt.Println("Json, j    :", string(j))

	j, _ = m.Json(true)
	fmt.Println("Json, j safe:", string(j))
}

func TestJsonWriter(t *testing.T) {
	mv := Map(map[string]interface{}{"this": "is a", "float": 3.14159, "and": "a", "bool": true})

	w := new(bytes.Buffer)
	raw, err := mv.JsonWriterRaw(w)
	if err != nil {
		t.Fatal("err:", err.Error())
	}

	b := make([]byte, w.Len())
	_, err = w.Read(b)
	if err != nil {
		t.Fatal("err:", err.Error())
	}

	fmt.Println("JsonWriter, raw:", string(raw))
	fmt.Println("JsonWriter, b  :", string(b))
}

func TestJsonLiteralEscapes(t *testing.T) {
	maps := []Map{
		{"value": `\u003c`},
		{"value": `\u003e`},
		{"value": `\u0026`},
		{`\u003c\u003e\u0026`: "escaped key"},
		{"nested": Map{"values": []interface{}{`\\u003c`, `before \u003e after`, `\u0026 <>&`}}},
		{"value": `<>&`},
		{"value": `\<\>\&`},
		{"value": "quotes: \" and newline: \n and separators: \u2028\u2029"},
		{},
		nil,
	}
	options := []struct {
		name string
		args []bool
	}{
		{"default", nil},
		{"false", []bool{false}},
		{"true", []bool{true}},
		{"multiple", []bool{true, true}},
	}
	encoders := []struct {
		name string
		call func(Map, []bool) ([]byte, error)
	}{
		{"Json", func(m Map, args []bool) ([]byte, error) { return m.Json(args...) }},
		{"JsonIndent", func(m Map, args []bool) ([]byte, error) { return m.JsonIndent("", "  ", args...) }},
		{"JsonWriter", func(m Map, args []bool) ([]byte, error) {
			var buf bytes.Buffer
			err := m.JsonWriter(&buf, args...)
			return buf.Bytes(), err
		}},
		{"JsonWriterRaw", func(m Map, args []bool) ([]byte, error) {
			var buf bytes.Buffer
			b, err := m.JsonWriterRaw(&buf, args...)
			if !bytes.Equal(b, buf.Bytes()) {
				t.Errorf("JsonWriterRaw returned %q but wrote %q", b, buf.Bytes())
			}
			return b, err
		}},
		{"JsonIndentWriter", func(m Map, args []bool) ([]byte, error) {
			var buf bytes.Buffer
			err := m.JsonIndentWriter(&buf, "", "  ", args...)
			return buf.Bytes(), err
		}},
		{"JsonIndentWriterRaw", func(m Map, args []bool) ([]byte, error) {
			var buf bytes.Buffer
			b, err := m.JsonIndentWriterRaw(&buf, "", "  ", args...)
			if !bytes.Equal(b, buf.Bytes()) {
				t.Errorf("JsonIndentWriterRaw returned %q but wrote %q", b, buf.Bytes())
			}
			return b, err
		}},
	}
	for _, encoder := range encoders {
		for _, option := range options {
			t.Run(encoder.name+"/"+option.name, func(t *testing.T) {
				for _, m := range maps {
					b, err := encoder.call(m, option.args)
					if err != nil {
						t.Fatal(err)
					}
					var decoded Map
					if err := json.Unmarshal(b, &decoded); err != nil {
						t.Errorf("invalid JSON for %#v: %q: %v", m, b, err)
						continue
					}
					want, err := json.Marshal(m)
					if err != nil {
						t.Fatal(err)
					}
					var expected Map
					if err := json.Unmarshal(want, &expected); err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(decoded, expected) {
						t.Errorf("decoded %#v, want %#v", decoded, expected)
					}
					if bytes.HasSuffix(b, []byte("\n")) {
						t.Errorf("unexpected trailing newline: %q", b)
					}
				}
			})
		}
	}
}

func TestJsonEncodingCompatibility(t *testing.T) {
	m := Map{"html": `<>&`, "literal": `\u003c\u003e\u0026`, "nested": Map{"value": "line\ntext"}}
	safe, err := m.Json(true)
	if err != nil {
		t.Fatal(err)
	}
	wantSafe, err := json.Marshal(m)
	if err != nil || !bytes.Equal(safe, wantSafe) {
		t.Fatalf("safe JSON = %q, want %q, error = %v", safe, wantSafe, err)
	}
	unsafe, err := m.Json()
	if err != nil || !bytes.Contains(unsafe, []byte(`"html":"<>&"`)) {
		t.Fatalf("default JSON = %q, error = %v", unsafe, err)
	}
	for _, format := range []struct{ prefix, indent string }{
		{"", ""}, {"", "  "}, {"", "\t"}, {"> ", "  "}, {`\u003c`, `\u0026`},
	} {
		for _, safeEncoding := range []bool{false, true} {
			compact, err := m.Json(safeEncoding)
			if err != nil {
				t.Fatal(err)
			}
			var expected bytes.Buffer
			if err := json.Indent(&expected, compact, format.prefix, format.indent); err != nil {
				t.Fatal(err)
			}
			b, err := m.JsonIndent(format.prefix, format.indent, safeEncoding)
			if err != nil || !bytes.Equal(b, expected.Bytes()) {
				t.Errorf("prefix %q, indent %q, safe %v: got %q, want %q, error %v",
					format.prefix, format.indent, safeEncoding, b, expected.Bytes(), err)
			}
			if safeEncoding {
				want, err := json.MarshalIndent(m, format.prefix, format.indent)
				if err != nil || !bytes.Equal(b, want) {
					t.Errorf("safe indent = %q, want %q, error = %v", b, want, err)
				}
			}
		}
	}
	for _, args := range [][]bool{nil, {false}, {true}} {
		invalid := Map{"value": make(chan int)}
		b, err := invalid.Json(args...)
		if err == nil || b != nil {
			t.Errorf("unsupported value: got %q, error %v", b, err)
		}
		b, err = invalid.JsonIndent("", "  ", args...)
		if err == nil || b != nil {
			t.Errorf("unsupported indented value: got %q, error %v", b, err)
		}
	}
}
