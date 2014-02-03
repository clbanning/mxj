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
*/

