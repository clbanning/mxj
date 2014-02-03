package mxj

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

// ------------------------------ write JSON -----------------------

// Just a wrapper on json.Marshal.
// If option unsafeEncoding is'true' then safe encoding of '<' and '>'
// is rolled back (see encoding/json#Marshal).
func (mv Map) Json(unsafeEncoding ...bool) ([]byte, error) {
	b, err := json.Marshal(mv)
	if len(unsafeEncoding) == 1 && unsafeEncoding[0] {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	}
	return b, err
}

func (mv Map) JsonErr(callErr ...error) ([]byte, error) {
	if  len(callErr) == 1 {
		return nil, callErr[0]
	}
	b, err := json.Marshal(mv)
	return b, err
}

// Just a wrapper on json.MarshalIndent.
// If option unsafeEncoding is'true' then safe encoding of '<' and '>'
// is rolled back (see encoding/json#Marshal).
func (mv Map) JsonIndent(prefix, indent string, unsafeEncoding ...bool) ([]byte, error) {
	b, err := json.MarshalIndent(mv, prefix, indent)
	if len(unsafeEncoding) == 1 && unsafeEncoding[0] {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	}
	return b, err
}

// The following implementation is provided for symmetry with NewMapJsonReader[Raw]
// The names will also provide a key for the number of return arguments.

// Writes the Map as JSON on the Writer. 
// Rolls back "safe" encoding of '<' and '>' - use with caution.
func (mv Map) JsonWriter(jsonWriter io.Writer) error {
	b, err := mv.Json()
	if err != nil {
		return err
	}

	_, err = jsonWriter.Write(b)
	return err
}

// Writes the Map as JSON on the Writer. *[]byte is the raw JSON that was written.
// Rolls back "safe" encoding of '<' and '>' - use with caution.
func (mv Map) JsonWriterRaw(jsonWriter io.Writer) (*[]byte, error) {
	b, err := mv.Json()
	if err != nil {
		return &b, err
	}

	_, err = jsonWriter.Write(b)
	return &b, err
}

// --------------------------- read JSON -----------------------------

// Just a wrapper on json.Unmarshal
//	Converting JSON to XML is a simple as:
//		...
//		mapVal, merr := mxj.NewMapJson(jsonVal)
//		if merr != nil {
//			// handle error
//		}
//		xmlVal, xerr := mapVal.Xml()
//		if xerr != nil {
//			// handle error
//		}
func NewMapJson(jsonVal []byte) (Map, error) {
	m := make(Map)
	err := json.Unmarshal(jsonVal, &m)
	return m, err
}

// Retrieve a Map value from an io.Reader.
func NewMapJsonReader(jsonReader io.Reader) (Map, error) {
	jb, err := getJson(jsonReader)
	if err != nil || len(*jb) == 0 {
		return nil, err
	}

	// Unmarshal the 'presumed' JSON string
	return NewMapJson(*jb)
}

// Retrieve a Map value and raw JSON - *[]byte - from an io.Reader.
func NewMapJsonReaderRaw(jsonReader io.Reader) (Map, *[]byte, error) {
	jb, err := getJson(jsonReader)
	if err != nil || len(*jb) == 0 {
		return nil, jb, err
	}

	// Unmarshal the 'presumed' JSON string
	m, merr := NewMapJson(*jb)
	return m, jb, merr
}

// Pull the next JSON string off the stream: just read from first '{' to its closing '}'.
func getJson(rdr io.Reader) (*[]byte, error) {
	bval := make([]byte, 1)
	jb := make([]byte, 0)
	var inQuote, inJson bool
	var parenCnt int

	// scan the input for a matched set of {...}
	// json.Unmarshal will handle syntax checking.
	for {
		_, err := rdr.Read(bval)
		if err != nil {
			if err == io.EOF && inJson && parenCnt > 0 {
				return nil, errors.New("no closing } for JSON string: " + string(jb))
			}
			return nil, err
		}
		switch bval[0] {
		case '{':
			if !inQuote {
				parenCnt++
				inJson = true
			}
		case '}':
			if !inQuote {
				parenCnt--
			}
			if parenCnt < 0 {
				return nil, errors.New("closing } without opening {: " + string(jb))
			}
		case '"':
			if inQuote {
				inQuote = false
			} else {
				inQuote = true
			}
		case '\n', '\r', '\t', ' ':
			if !inQuote {
				continue
			}
		}
		if inJson {
			jb = append(jb, bval...)
			if parenCnt == 0 {
				break
			}
		}
	}

	return &jb, nil
}

// ------------------------------- JSON Reader handler via Map values  -----------------------

// Default poll delay to keep Handler from spinning on an open stream
// like sitting on os.Stdin waiting for imput.
var jhandlerPollInterval = time.Duration(1e6)

// While unnecessary, we make HandleJsonReader() have the same signature as HandleXmlReader().
// This avoids treating one or other as a special case and discussing the underlying stdlib logic.

// Bulk process JSON using handlers that process a Map value.
//	'rdr' is an io.Reader for the JSON (stream).
//	'mapHandler' is the Map processing handler. Return of 'false' stops further processing.
//	'errHandler' is the error. Return of 'false' stop processing and returns error.
//	Note: mapHandler() and errHandler() calls are blocking, so reading and processing of messages is serialized.
//	      This means that you can stop reading the file on error or after processing a particular message.
//	      To have reading and handling run concurrently, pass argument(s) to a go routine in handler and return true.
func HandleJsonReader(jsonReader io.Reader, mapHandler func(Map) bool, errHandler func(error) bool) error {
	var n int
	for {
		m, merr := NewMapJsonReader(jsonReader)
		n++

		// handle error condition with errhandler
		if merr != nil && merr != io.EOF {
			merr = errors.New(fmt.Sprintf("[jsonReader: %d] %s", n, merr.Error()))
			if ok := errHandler(merr); !ok {
				// caused reader termination
				return merr
			}
			continue
		}

		// pass to maphandler
		if len(m) != 0 {
			if ok := mapHandler(m); !ok {
				break
			}
		} else if merr != io.EOF {
			<-time.After(jhandlerPollInterval)
		}

		if merr == io.EOF {
			break
		}
	}
	return nil
}

// Bulk process JSON using handlers that process a Map value and the raw JSON.
//	'rdr' is an io.Reader for the JSON (stream).
//	'mapHandler' is the Map and raw JSON - *[]byte - processing handler. Return of 'false' stops further processing.
//	'errHandler' is the error and raw JSON processing handler. Return of err != nil stop processing and returns error.
//	Note: mapHandler() and errHandler() calls are blocking, so reading and processing of messages is serialized.
//	      This means that you can stop reading the file on error or after processing a particular message.
//	      To have reading and handling run concurrently, pass argument(s) to a go routine in handler and return true.
func HandleJsonReaderRaw(jsonReader io.Reader, mapHandler func(Map, *[]byte) bool, errHandler func(error, *[]byte) bool) error {
	var n int
	for {
		m, raw, merr := NewMapJsonReaderRaw(jsonReader)
		n++

		// handle error condition with errhandler
		if merr != nil && merr != io.EOF {
			merr = errors.New(fmt.Sprintf("[jsonReader: %d] %s", n, merr.Error()))
			if ok := errHandler(merr, raw); !ok {
				// caused reader termination
				return merr
			}
			continue
		}

		// pass to maphandler
		if len(m) != 0 {
			if ok := mapHandler(m, raw); !ok {
				break
			}
		} else if merr != io.EOF {
			<-time.After(jhandlerPollInterval)
		}

		if merr == io.EOF {
			break
		}
	}
	return nil
}
