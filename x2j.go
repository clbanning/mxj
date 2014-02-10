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
*/

