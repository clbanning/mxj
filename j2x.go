// Copyright 2012-2014 Charles Banning. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

// j2x.go - wrappers for end-to-end transformatioin of JSON to XML
// For (mostly) backwards compatibility with j2x package.

package mxj

// Wrappers for end-to-end JSON to XML transformation.

/*
import "io"

// FromJson() --> ToXml().
func JsonToXml(jsonVal []byte) ([]byte, error) {
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return nil, merr
	}
	return m.Xml()
}

// FromJson() --> ToXmlWriter().
func JsonToXmlWriter(jsonVal []byte, xmlWriter io.Writer) (*[]byte, error) {
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return nil, merr
	}
	return m.XmlWriter(xmlWriter)
}

// FromJsonReader() --> ToXml().
func JsonReaderToXml(jsonReader io.Reader) ([]byte, *[]byte, error) {
	m, raw, merr := NewMapJsonReader(jsonReader)
	if merr != nil {
		return nil, raw, merr
	}
	x, xerr := m.Xml()
	return x, raw, xerr
}

// FromJsonReader() --> ToXmlWriter().  Handy for transforming bulk message sets.
func JsonReaderToXmlWriter(jsonReader io.Reader, xmlWriter io.Writer) (*[]byte, *[]byte, error) {
	m, jraw, merr := NewMapJsonReader(jsonReader)
	if merr != nil {
		return jraw, nil, merr
	}
	xraw, xerr := m.XmlWriter(xmlWriter)
	return jraw, xraw, xerr
}
*/

// JSON wrappers for Map methods implementing key path and value functions.

/*
// Wrap PathsForKey for JSON.
func JsonPathsForKey(jsonVal []byte, key string) ([]string, error) {
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return nil, merr
	}
	paths := m.PathsForKey(key)
	return paths, nil
}

// Wrap PathForKeyShortest for JSON.
func JsonPathForKeyShortest(jsonVal []byte, key string) (string, error) {
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return "", merr
	}
	path := m.PathForKeyShortest(key)
	return path, nil
}

// Wrap ValuesForKey for JSON.
func JsonValuesForKey(jsonVal []byte, key string, subkeys ...string) ([]interface{}, error) {
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return nil, merr
	}
	return m.ValuesForKey(key, subkeys...)
}

// Wrap ValuesForKeyPath for JSON.
func JsonValuesForKeyPath(jsonVal []byte, path string, subkeys ...string) ([]interface{}, error) {
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return nil, merr
	}
	return m.ValuesForPath(path, subkeys...)
}

// Wrap UpdateValuesForPath for JSON
//	'jsonVal' is XML value
//	'newKeyValue' is the value to replace an existing value at the end of 'path'
//	'path' is the dot-notation path with the key whose value is to be replaced at the end
//	       (can include wildcard character, '*')
//	'subkeys' are key:value pairs of key:values that must match for the key
func JsonUpdateValsForPath(jsonVal []byte, newKeyValue interface{}, path string, subkeys ...string) ([]byte, error) {
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return 0, err
	}
	n, err := m.UpdateValuesForPath(newJsonValue, path, subkeys...)
	if err != nil {
		return nil, err
	}
	return m.Json() 
}

// Wrap NewMap for JSON and return as JSON
// 'jsonVal' is an JSON value
// 'keypairs' are "oldKey:newKey" values that conform to 'keypairs' in (Map)NewMap.
func JsonNewJson(jsonVal []byte, keypairs ...string) ([]byte, error) (
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return nil, merr
	}
	n, nerr := m.NewMap(keypairs...)
	if nerr != nil {
		return nil, nerr
	}
	return n.Json()
}

// Wrap NewMap for JSON and return as XML
// 'jsonVal' is an JSON value
// 'keypairs' are "oldKey:newKey" values that conform to 'keypairs' in (Map)NewMap.
func JsonNewXml(jsonVal []byte, keypairs ...string) ([]byte, error) (
	m, merr := NewMapJson(jsonVal)
	if merr != nil {
		return nil, merr
	}
	n, nerr := m.NewMap(tagpairs...)
	if nerr != nil {
		return nil, nerr
	}
	return n.Xml()
}

*/

