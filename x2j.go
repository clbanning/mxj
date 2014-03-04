// Copyright 2012-2014 Charles Banning. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

// x2j - for (mostly) backwards compatibility with x2j package

package mxj

// Wrappers for end-to-end XML to JSON transformation.

/*
import "io"

// FromXml() --> ToJson().
func XmlToJson(xmlVal []byte, safeEncoding ...bool) ([]byte, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	return m.Json(safeEncoding...)
}

// FromXml() --> ToJsonWriter().
func XmlToJsonWriter(xmlVal []byte, jsonWriter io.Writer, safeEncoding ...bool) (*[]byte, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	return m.JsonWriter(jsonWriter, safeEncoding...)
}

// FromXmlReader() --> ToJson().
func XmlReaderToJson(xmlReader io.Reader, safeEncoding ...bool) ([]byte, *[]byte, error) {
	m, raw, merr := NewMapXmlReader(xmlReader)
	if merr != nil {
		return nil, raw, merr
	}
	j, jerr := m.Json(safeEncoding...)
	return j, raw, jerr
}

// FromXmlReader() --> ToJsonWriter().  Handy for bulk transformation of documents.
func XmlReaderToJsonWriter(xmlReader io.Reader, jsonWriter io.Writer, safeEncoding ...bool) (*[]byte, *[]byte, error) {
	m, xraw, merr := NewMapXmlReader(xmlReader)
	if merr != nil {
		return xraw, nil, merr
	}
	jraw, jerr := m.JsonWriter(jsonWriter, safeEncoding...)
	return xraw, jraw, jerr
}
*/

// XML wrappers for Map methods implementing tag path and value functions.

/*
// Wrap PathsForKey for XML.
func XmlPathsForTag(xmlVal []byte, tag string) ([]string, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	paths := m.PathsForKey(tag)
	return paths, nil
}

// Wrap PathForKeyShortest for XML.
func XmlPathForTagShortest(xmlVal []byte, tag string) (string, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return "", merr
	}
	path := m.PathForKeyShortest(tag)
	return path, nil
}

// Wrap ValuesForKey for XML.
// 'attrs' are key:value pairs for attributes, where key is attribute label prepended with a hypen, '-'.
func XmlValuesForTag(xmlVal []byte, tag string, attrs ...string) ([]interface{}, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	return m.ValuesForKey(tag, attrs...)
}

// Wrap ValuesForPath for XML.
// 'attrs' are key:value pairs for attributes, where key is attribute label prepended with a hypen, '-'.
func XmlValuesForTagPath(xmlVal []byte, path string, attrs ...string) ([]interface{}, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	return m.ValuesForPath(path, attrs...)
}

// Wrap UpdateValuesForPath for XML
//	'xmlVal' is XML value
//	'newTagValue' is the value to replace an existing value at the end of 'path'
//	'path' is the dot-notation path with the tag whose value is to be replaced at the end
//	       (can include wildcard character, '*')
//	'subkeys' are key:value pairs of tag:values that must match for the tag
func XmlUpdateValsForPath(xmlVal []byte, newTagValue interface{}, path string, subkeys ...string) ([]byte, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return 0, err
	}
	n, err := m.UpdateValuesForPath(newTagValue, path, subkeys...)
	if err != nil {
		return nil, err
	}
	return m.Xml() 
}

// Wrap NewMap for XML and return as XML
// 'xmlVal' is an XML value
// 'tagpairs' are "oldTag:newTag" values that conform to 'keypairs' in (Map)NewMap.
func XmlNewXml(xmlVal []byte, tagpairs ...string) ([]byte, error) (
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	n, nerr := m.NewMap(tagpairs...)
	if nerr != nil {
		return nil, nerr
	}
	return n.Xml()
}

// Wrap NewMap for XML and return as JSON
// 'xmlVal' is an XML value
// 'tagpairs' are "oldTag:newTag" values that conform to 'keypairs' in (Map)NewMap.
func XmlNewJson(xmlVal []byte, tagpairs ...string) ([]byte, error) (
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	n, nerr := m.NewMap(tagpairs...)
	if nerr != nil {
		return nil, nerr
	}
	return n.Json()
}

*/

