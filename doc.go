// mxj - A collection of map[string]interface{} and associated XML and JSON utilities.
// Copyright 2012-2014 Charles Banning. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

/*
Marshal/Unmarshal XML to/from JSON and map[string]interface{} values, and extract values from maps by key or key-path, including wildcards.

mxj supplants the legacy x2j and j2x packages. If you want the old syntax, just uncomment the content of x2j.go
or j2x.go, which are included as part of this package.

SUMMARY

	type Map map[string]interface{}

	Create a Map value, 'm', from any map[string]interface{} value, 'v':
		m := Map(v)

	Unmarshal / marshal XML as a Map value, 'm':
		m, err := NewMapXml(xmlValue) // unmarshal
		xmlValue, err := m.Xml()      // marshal

	Unmarshal XML from an io.Reader as a Map value, 'm':
		m, err := NewMapReader(xmlReader)         // repeated calls, as with an os.File Reader, will process stream
		m, raw, err := NewMapReaderRaw(xmlReader) // 'raw' points to the raw XML that was decoded

	Marshal Map value, 'm', to an XML Writer (io.Writer):
		err := m.XmlWriter(xmlWriter)         // or: m.XmlIndentWriter(xmlWriter, prefix, indent)
		err, raw := m.XmlWriterRaw(xmlWriter) // 'raw' points to the raw XML that was written on xmlWriter

	Bulk process XML with error handling:
		err := HandleXmlReader(xmlReader, mapHandler(Map), errHandler(error))
		err := HandleXmlReaderRaw(xmlReader, mapHandler(Map, *[]byte), errHandler(error, *[]byte))

	Converting XML to JSON: see Examples for NewMapXml and HandleXmlReader.

	There are comparable functions and methods for JSON processing.

	Arbitrary structure values can be decoded to / encoded from Map values:
		m, err := NewMapStruct(structVal)
		err := m.Struct(structPointer)

	To work with XML tag values, JSON or Map key values or structure field values, decode the XML, JSON
	or structure to a Map value, 'm', or cast a map[string]interface{} value to a Map value, 'm', then:
		paths := m.PathsForKey(key)
		path := m.PathForKeyShortest(key)
		values, err := m.ValuesForKey(key, subkeys)
		values, err := m.ValuesForPath(path, subkeys)
		count, err := m.UpdateValuesForPath(newVal, path, subkeys)

	A new Map with whatever keys are desired can be created from the current Map and then encoded in XML
	or JSON. (Note: keys can use dot-notation.)
		newMap := m.NewMap("oldKey_1:newKey_1", "oldKey_2:newKey_2", ..., "oldKey_N:newKey_N")
		newXml := newMap.Xml()     // for example
		newJson := newMap.Json()   // ditto 

XML PARSING CONVENTIONS

   - Attributes are parsed to map[string]interface{} values by prefixing a hyphen, '-',
     to the attribute label.
   - If the element is a simple element and has attributes, the element value
     is given the key '#text' for its map[string]interface{} representation. 

*/
package mxj
