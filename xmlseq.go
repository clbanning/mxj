// Copyright 2012-2015 Charles Banning. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

// xmlseq.go - version of xml.go with sequence # injection on Decoding and sorting on Encoding.
// Also, handles comments, directives and process instructions.

package mxj

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// ------------------- NewMapXmlSeq & NewMapXmlSeqReader ... -------------------------

// THIS IS EXPERIMENTAL!
//
// It is only useful if you want to re-encode the Map as XML using mv.XmlSeq(), etc., to preserve the original structure.
//
// NewMapXmlSeq - convert a XML doc into a Map with elements id'd with decoding sequence int - #seq.
// If the optional argument 'cast' is 'true', then values will be converted to boolean or float64 if possible.
// NOTE: "#seq" key/value pairs are removed on encoding with mv.XmlSeq() / mv.XmlSeqIndent().
//	• attributes are an array - map["#attr"][]map["attr_key"]interface{}
//	• all simple elements are decoded as map["#text"]interface{} with a "#seq" k:v pair, as well.
//	• lists always decode as map["list_tag"][]map[string]interface{} where the array elements include
//	  a "#seq" k:v pair based on sequence they are decoded.  Thus, XML like:
//	      <ltag>value 1</ltag>
//	      <newtag>value 2</newtag>
//	      <ltag>value 3</ltag>
//	  will encode in proper sequence even though the Map representation merges all "ltag" elements in an array.
//	• comments - "<!--comment-->" -  are decoded as map["#comment"]map["#text"]"cmnt_text" with a "#seq" k:v pair.
//	• directives - "<!text>" - are decoded as map["#directive"]map[#text"]"directive_text" with a "#seq" k:v pair.
//	• process instructions  - "<?instr?>" - are decoded as map["#procinst"]interface{} where the #procinst value
//	  is of map[string]interface{} type with the following keys: #target, #inst, and #seq.
//	• note: "<![CDATA[" syntax is lost in xml.Decode parser - and is not handled here, either.
//	   and: "\r\n" is converted to "\n"
func NewMapXmlSeq(xmlVal []byte, cast ...bool) (Map, error) {
	var r bool
	if len(cast) == 1 {
		r = cast[0]
	}
	return xmlSeqToMap(xmlVal, r)
}

// THIS IS EXPERIMENTAL!
//
// It is only useful if you want to re-encode the Map as XML using mv.XmlSeq(), etc., to preserve the original structure.
//
// Get next XML doc from an io.Reader as a Map value.  Returns Map value.
// See NewMapXmlSeq for "#seq" key insertion.
func NewMapXmlSeqReader(xmlReader io.Reader, cast ...bool) (Map, error) {
	var r bool
	if len(cast) == 1 {
		r = cast[0]
	}

	// build the node tree
	return xmlSeqReaderToMap(xmlReader, r)
}

// THIS IS EXPERIMENTAL!
//
// It is only useful if you want to re-encode the Map as XML using mv.XmlSeq(), etc., to preserve the original structure.
//
// Get next XML doc from an io.Reader as a Map value.  Returns Map value and slice with the raw XML.
//	NOTES: 1. Due to the implementation of xml.Decoder, the raw XML off the reader is buffered to []byte
//	          using a ByteReader. If the io.Reader is an os.File, there may be significant performance impact.
//	          See the examples - getmetrics1.go through getmetrics4.go - for comparative use cases on a large
//	          data set. If the io.Reader is wrapping a []byte value in-memory, however, such as http.Request.Body
//	          you CAN use it to efficiently unmarshal a XML doc and retrieve the raw XML in a single call.
//	       2. The 'raw' return value may be larger than the XML text value.
func NewMapXmlSeqReaderRaw(xmlReader io.Reader, cast ...bool) (Map, []byte, error) {
	var r bool
	if len(cast) == 1 {
		r = cast[0]
	}
	// create TeeReader so we can retrieve raw XML
	buf := make([]byte, XmlWriterBufSize)
	wb := bytes.NewBuffer(buf)
	trdr := myTeeReader(xmlReader, wb)

	// build the node tree
	m, err := xmlSeqReaderToMap(trdr, r)

	// retrieve the raw XML that was decoded
	b := make([]byte, wb.Len())
	_, _ = wb.Read(b)

	if err != nil {
		return nil, b, err
	}

	return m, b, nil
}

// xmlReaderToTree() - parse a XML io.Reader to a map[string]interface{} value
func xmlSeqReaderToMap(rdr io.Reader, r bool) (map[string]interface{}, error) {
	// parse the Reader
	p := xml.NewDecoder(rdr)
	p.CharsetReader = XmlCharsetReader
	return xmlSeqToMapParser("", nil, p, r)
}

// xmlSeqToMap - convert a XML doc into map[string]interface{} value
func xmlSeqToMap(doc []byte, r bool) (map[string]interface{}, error) {
	b := bytes.NewReader(doc)
	p := xml.NewDecoder(b)
	p.CharsetReader = XmlCharsetReader
	return xmlSeqToMapParser("", nil, p, r)
}

// ===================================== where the work happens =============================

// xmlSeqToMapParser - load a 'clean' XML doc into a map[string]interface{} directly.
// Add #seq tag value for each element decoded - to be used for Encoding later.
func xmlSeqToMapParser(skey string, a []xml.Attr, p *xml.Decoder, r bool) (map[string]interface{}, error) {
	// NOTE: all attributes and sub-elements parsed into 'na', 'na' is returned as value for 'skey'
	// Unless 'skey' is a simple element w/o attributes, in which case the xml.CharData value is the value.
	var n, na map[string]interface{}
	var seq int // for including seq num when decoding

	// Allocate maps and load attributes, if any.
	if skey != "" {
		// 'n' only needs one slot - save call to runtime•hashGrow()
		// 'na' we don't know
		n = make(map[string]interface{}, 1)
		na = make(map[string]interface{})
		if len(a) > 0 {
			// xml.Attr is decoded into: map["#attr"][]map[string]interface{}
			aa := make([]interface{}, len(a))
			for n, v := range a {
				am := make(map[string]interface{}, 1)
				am[v.Name.Local] = cast(v.Value, r)
				aa[n] = interface{}(am)
			}
			na["#attr"] = aa // note: a will encode in sequence of decoding
		}
	}
	for {
		t, err := p.Token()
		if err != nil {
			if err != io.EOF {
				return nil, errors.New("xml.Decoder.Token() - " + err.Error())
			}
			return nil, err
		}
		switch t.(type) {
		case xml.StartElement:
			tt := t.(xml.StartElement)

			// First call to xmlSeqToMapParser() doesn't pass xml.StartElement - the map key.
			// So when the loop is first entered, the first token is the root tag along
			// with any attributes, which we process here.
			//
			// Subsequent calls to xmlSeqToMapParser() will pass in tag+attributes for
			// processing before getting the next token which is the element value,
			// which is done above.
			if skey == "" {
				return xmlSeqToMapParser(tt.Name.Local, tt.Attr, p, r)
			}

			// If not initializing the map, parse the element.
			// len(nn) == 1, necessarily - it is just an 'n'.
			nn, err := xmlSeqToMapParser(tt.Name.Local, tt.Attr, p, r)
			if err != nil {
				return nil, err
			}

			// The nn map[string]interface{} value is a na[nn_key] value.
			// We need to see if nn_key already exists - means we're parsing a list.
			// This may require converting na[nn_key] value into []interface{} type.
			// First, extract the key:val for the map - it's a singleton.
			var key string
			var val interface{}
			for key, val = range nn {
				break
			}

			// add "#seq" k:v pair -
			// Sequence number included even in list elements - this should allow us
			// to properly resequence even something goofy like:
			//     <list>item 1</list>
			//     <subelement>item 2</subelement>
			//     <list>item 3</list>
			// where all the "list" subelements are decoded into an array.
			switch val.(type) {
			case map[string]interface{}:
				val.(map[string]interface{})["#seq"] = seq
				seq++
			case interface{}: // a non-nil simple element: string, float64, bool
				v := map[string]interface{}{"#text": val}
				v["#seq"] = seq
				seq++
				val = v
			}

			// 'na' holding sub-elements of n.
			// See if 'key' already exists.
			// If 'key' exists, then this is a list, if not just add key:val to na.
			if v, ok := na[key]; ok {
				var a []interface{}
				switch v.(type) {
				case []interface{}:
					a = v.([]interface{})
				default: // anything else - note: v.(type) != nil
					a = []interface{}{v}
				}
				a = append(a, val)
				na[key] = a
			} else {
				na[key] = val // save it as a singleton
			}
		case xml.EndElement:
			// len(n) > 0 if this is a simple element w/o xml.Attrs.
			if len(n) == 0 {
				// If len(na)==0 we have an empty element == "";
				// it has no xml.Attr nor xml.CharData.
				// Empty element content will be  map["etag"]map["#text"]""
				// after #seq injection - map["etag"]map["#seq"]seq - after return.
				if len(na) > 0 {
					n[skey] = na
				} else {
					n[skey] = "" // empty element
				}
			}
			return n, nil
		case xml.CharData:
			tt := string(t.(xml.CharData))
			// clean up possible noise
			tt = strings.Trim(tt, "\t\r\b\n ")
			if len(tt) > 0 {
				// every simple element is a #text and has #seq associated with it
				na["#text"] = cast(tt, r)
				na["#seq"] = seq
				seq++
			}
		case xml.Comment:
			cm := make(map[string]interface{}, 2)
			cm["#text"] = string(t.(xml.Comment))
			cm["#seq"] = seq
			seq++
			na["#comment"] = cm
		case xml.Directive:
			dm := make(map[string]interface{}, 2)
			dm["#text"] = string(t.(xml.Directive))
			dm["#seq"] = seq
			seq++
			na["#directive"] = dm
		case xml.ProcInst:
			pm := make(map[string]interface{}, 3)
			pm["#target"] = t.(xml.ProcInst).Target
			pm["#inst"] = string(t.(xml.ProcInst).Inst)
			pm["#seq"] = seq
			seq++
			na["#procinst"] = pm
		default:
			// noop - shouldn't ever get here, now, since we handle all token types
		}
	}
}

// ------------------ END: NewMapXml & NewMapXmlReader -------------------------

// ------------------ mv.Xml & mv.XmlWriter - from j2x ------------------------

// THIS IS EXPERIMENTAL!
//
// It should ONLY be used on Map values that were decoded using NewMapXmlSeq() & co.
//
// Encode a Map as XML with elements sorted on #seq.  The companion of NewMapXmlSeq().
// The following rules apply.
//    - The key label "#text" is treated as the value for a simple element with attributes.
//    - The "#seq" key is used to seqence the subelements but is ignored for writing.
//    - The "#attr" map key identifies the array of attribute map[string]interface{} values.
//    - The "#comment" map key identifies a comment in the value "#text" map entry - <!--comment-->.
//    - The "#directive" map key identifies a directive in the value "#text" map entry - <!directive>.
//    - The "#procinst" map key identifies a process instruction in the value "#target" and "#inst"
//      map entries - <?target inst?>.
//    - Value type encoding:
//          > string, bool, float64, int, int32, int64, float32: per "%v" formating
//          > []bool, []uint8: by casting to string
//          > structures, etc.: handed to xml.Marshal() - if there is an error, the element
//            value is "UNKNOWN"
//    - Elements with only attribute values or are null are terminated using "/>" unless XmlGoEmptyElemSystax() called.
//    - If len(mv) == 1 and no rootTag is provided, then the map key is used as the root tag, possible.
//      Thus, `{ "key":"value" }` encodes as "<key>value</key>".
func (mv Map) XmlSeq(rootTag ...string) ([]byte, error) {
	m := map[string]interface{}(mv)
	var err error
	s := new(string)
	p := new(pretty) // just a stub

	if len(m) == 1 && len(rootTag) == 0 {
		for key, value := range m {
			// if it an array, see if all values are map[string]interface{}
			// we force a new root tag if we'll end up with no key:value in the list
			// so: key:[string_val, bool:true] --> <doc><key>string_val</key><bool>true</bool></doc>
			switch value.(type) {
			case []interface{}:
				for _, v := range value.([]interface{}) {
					switch v.(type) {
					case map[string]interface{}: // noop
					default: // anything else
						err = mapToXmlSeqIndent(false, s, DefaultRootTag, m, p)
						goto done
					}
				}
			}
			err = mapToXmlSeqIndent(false, s, key, value, p)
		}
	} else if len(rootTag) == 1 {
		err = mapToXmlSeqIndent(false, s, rootTag[0], m, p)
	} else {
		err = mapToXmlSeqIndent(false, s, DefaultRootTag, m, p)
	}
done:
	return []byte(*s), err
}

// The following implementation is provided only for symmetry with NewMapXmlReader[Raw]
// The names will also provide a key for the number of return arguments.

// THIS IS EXPERIMENTAL!
//
// It should ONLY be used on Map values that were decoded using NewMapXmlSeq() & co.
//
// Writes the Map as  XML on the Writer.
// See Xml() for encoding rules.
func (mv Map) XmlSeqWriter(xmlWriter io.Writer, rootTag ...string) error {
	x, err := mv.XmlSeq(rootTag...)
	if err != nil {
		return err
	}

	_, err = xmlWriter.Write(x)
	return err
}

// THIS IS EXPERIMENTAL!
//
// It should ONLY be used on Map values that were decoded using NewMapXmlSeq() & co.
//
// Writes the Map as  XML on the Writer. []byte is the raw XML that was written.
// See Xml() for encoding rules.
func (mv Map) XmlSeqWriterRaw(xmlWriter io.Writer, rootTag ...string) ([]byte, error) {
	x, err := mv.XmlSeq(rootTag...)
	if err != nil {
		return x, err
	}

	_, err = xmlWriter.Write(x)
	return x, err
}

// THIS IS EXPERIMENTAL!
//
// It should ONLY be used on Map values that were decoded using NewMapXmlSeq() & co.
//
// Writes the Map as pretty XML on the Writer.
// See Xml() for encoding rules.
func (mv Map) XmlSeqIndentWriter(xmlWriter io.Writer, prefix, indent string, rootTag ...string) error {
	x, err := mv.XmlSeqIndent(prefix, indent, rootTag...)
	if err != nil {
		return err
	}

	_, err = xmlWriter.Write(x)
	return err
}

// THIS IS EXPERIMENTAL!
//
// It should ONLY be used on Map values that were decoded using NewMapXmlSeq() & co.
//
// Writes the Map as pretty XML on the Writer. []byte is the raw XML that was written.
// See Xml() for encoding rules.
func (mv Map) XmlSeqIndentWriterRaw(xmlWriter io.Writer, prefix, indent string, rootTag ...string) ([]byte, error) {
	x, err := mv.XmlSeqIndent(prefix, indent, rootTag...)
	if err != nil {
		return x, err
	}

	_, err = xmlWriter.Write(x)
	return x, err
}

// -------------------- END: mv.Xml & mv.XmlWriter -------------------------------

// ---------------------- XmlSeqIndent ----------------------------

// THIS IS EXPERIMENTAL!
//
// It should ONLY be used on Map values that were decoded using NewMapXmlSeq() & co.
//
// Encode a map[string]interface{} as a pretty XML string.
// See mv.XmlSeq() for encoding rules.
func (mv Map) XmlSeqIndent(prefix, indent string, rootTag ...string) ([]byte, error) {
	m := map[string]interface{}(mv)

	var err error
	s := new(string)
	p := new(pretty)
	p.indent = indent
	p.padding = prefix

	if len(m) == 1 && len(rootTag) == 0 {
		// this can extract the key for the single map element
		// use it if it isn't a key for a list
		for key, value := range m {
			if _, ok := value.([]interface{}); ok {
				err = mapToXmlSeqIndent(true, s, DefaultRootTag, m, p)
			} else {
				err = mapToXmlSeqIndent(true, s, key, value, p)
			}
		}
	} else if len(rootTag) == 1 {
		err = mapToXmlSeqIndent(true, s, rootTag[0], m, p)
	} else {
		err = mapToXmlSeqIndent(true, s, DefaultRootTag, m, p)
	}
	return []byte(*s), err
}

// where the work actually happens
// returns an error if an attribute is not atomic
func mapToXmlSeqIndent(doIndent bool, s *string, key string, value interface{}, pp *pretty) error {
	var endTag bool
	var isSimple bool
	var elen int
	p := &pretty{pp.indent, pp.cnt, pp.padding, pp.mapDepth, pp.start}

	switch value.(type) {
	case map[string]interface{}, []byte, string, float64, bool, int, int32, int64, float32:
		if doIndent {
			*s += p.padding
		}
		if key != "#comment" && key != "#directive" && key != "#procinst" {
			*s += `<` + key
		}
	}
	switch value.(type) {
	case map[string]interface{}:
		val := value.(map[string]interface{})

		if key == "#comment" {
			*s += `<!--` + val["#text"].(string) + `-->`
			break
		}

		if key == "#directive" {
			*s += `<!` + val["#text"].(string) + `>`
			break
		}

		if key == "#procinst" {
			*s += `<?` + val["#target"].(string) + ` ` + val["#inst"].(string) + `?>`
			break
		}

		haveAttrs := false
		// process attributes first 
		// They are in sequence in the array.
		if v, ok := val["#attr"].([]interface{}); ok {
			for _, vv := range v {
				for k, av := range vv.(map[string]interface{}) {
					switch av.(type) {
					case string, float64, bool, int, int32, int64, float32:
						*s += ` ` + k + `="` + fmt.Sprintf("%v", av) + `"`
					case []byte:
						*s += ` ` + k + `="` + fmt.Sprintf("%v", string(av.([]byte))) + `"`
					default:
						return fmt.Errorf("invalid attribute value for: %s", k)
					}
				}
			}
			haveAttrs = true
		}

		// only attributes?
		if len(val) == 0 {
			break
		}

		// simple element? Note: '#text" is an invalid XML tag.
		// a simple elment
		if v, ok := val["#text"]; ok && ((len(val) == 3 && haveAttrs) || (len(val) == 2 && !haveAttrs)) {
			*s += ">" + fmt.Sprintf("%v", v)
			endTag = true
			elen = 1
			isSimple = true
			break
		}

		// we now need to sequence everything except attributes
		// 'kv' will hold everything that needs to be written
		kv := make([]keyval, 0)
		var kvv keyval
		for k, v := range val {
			if k == "#attr" { // already processed
				continue
			}
			if k == "#seq" { // ignore - just for sorting
				continue
			}
			switch v.(type) {
			case []interface{}:
				// unwind the array as separate entries
				for _, vv := range v.([]interface{}) {
					kvv = keyval{k, vv}
					kv = append(kv, kvv)
				}
			default:
				kvv = keyval{k, v}
				kv = append(kv, kvv)
			}
		}

		// close tag with possible attributes
		*s += ">"
		if doIndent {
			*s += "\n"
		}
		// something more complex
		p.mapDepth++
		// PrintElemListSeq(elemListSeq(kv))
		sort.Sort(elemListSeq(kv))
		// PrintElemListSeq(elemListSeq(kv))
		i := 0
		for _, v := range kv {
			switch v.v.(type) {
			case []interface{}:
			default:
				if i == 0 && doIndent {
					p.Indent()
				}
			}
			i++
			mapToXmlSeqIndent(doIndent, s, v.k, v.v, p)
			switch v.v.(type) {
			case []interface{}: // handled in []interface{} case
			default:
				if doIndent {
					p.Outdent()
				}
			}
			i--
		}
		p.mapDepth--
		endTag = true
		elen = 1 // we do have some content ...
	case []interface{}:
		for _, v := range value.([]interface{}) {
			if doIndent {
				p.Indent()
			}
			mapToXmlSeqIndent(doIndent, s, key, v, p)
			if doIndent {
				p.Outdent()
			}
		}
		return nil
	case nil:
		// terminate the tag
		*s += "<" + key
		break
	default: // handle anything - even goofy stuff
		elen = 0
		switch value.(type) {
		case string, float64, bool, int, int32, int64, float32:
			v := fmt.Sprintf("%v", value)
			elen = len(v)
			if elen > 0 {
				*s += ">" + v
			}
		case []byte: // NOTE: byte is just an alias for uint8
			// similar to how xml.Marshal handles []byte structure members
			v := string(value.([]byte))
			elen = len(v)
			if elen > 0 {
				*s += ">" + v
			}
		default:
			var v []byte
			var err error
			if doIndent {
				v, err = xml.MarshalIndent(value, p.padding, p.indent)
			} else {
				v, err = xml.Marshal(value)
			}
			if err != nil {
				*s += ">UNKNOWN"
			} else {
				elen = len(v)
				if elen > 0 {
					*s += string(v)
				}
			}
		}
		isSimple = true
		endTag = true
	}
	if endTag {
		if doIndent {
			if !isSimple {
				*s += p.padding
			}
		}
		switch value.(type) {
		case map[string]interface{}, []byte, string, float64, bool, int, int32, int64, float32:
			if elen > 0 || useGoXmlEmptyElemSyntax {
				if elen == 0 {
					*s += ">"
				}
				*s += `</` + key + ">"
			} else {
				*s += `/>`
			}
		}
	} else if key != "#comment"  && key != "#directive" && key != "#procinst" {
		if useGoXmlEmptyElemSyntax {
			*s += "></" + key + ">"
		} else {
			*s += "/>"
		}
	}
	if doIndent {
		if p.cnt > p.start {
			*s += "\n"
		}
		p.Outdent()
	}

	return nil
}

// the element sort implementation

type keyval struct {
	k string
	v interface{}
}
type elemListSeq []keyval

func (e elemListSeq) Len() int {
	return len(e)
}

func (e elemListSeq) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e elemListSeq) Less(i, j int) bool {
	var iseq, jseq int
	var ok bool
	if iseq, ok = e[i].v.(map[string]interface{})["#seq"].(int); !ok {
		iseq = 9999999
	}

	if jseq, ok = e[j].v.(map[string]interface{})["#seq"].(int); !ok {
		jseq = 9999999
	}

	if iseq > jseq {
		return false
	}
	return true
}

func PrintElemListSeq(e elemListSeq) {
	for n, v := range e {
		fmt.Printf("%d: %v\n", n, v)
	}
}
