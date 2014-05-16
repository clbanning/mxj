package mxj

import (
	"fmt"
	"io"
	"os"
)

// ReadMapsFromXmlFile - creates an array from a file of JSON values.
func ReadMapsFromJsonFile(name string) ([]Map, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("file %s is not a regular file", name)
	}

	fh, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	am := make([]Map,0)
	for {
		m, raw, err := NewMapJsonReaderRaw(fh)
		if err != nil && err != io.EOF {
			return am, fmt.Errorf("error: %s - reading: %s", err.Error(), string(raw))
		}
		if len(m) > 0 {
			am = append(am, m)
		}
		if err == io.EOF {
			break
		}
	}
	return am, nil
}

// ReadMapsFromXmlFile - creates an array from a file of XML values.
func ReadMapsFromXmlFile(name string) ([]Map, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("file %s is not a regular file", name)
	}

	fh, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	am := make([]Map,0)
	for {
		m, raw, err := NewMapXmlReaderRaw(fh)
		if err != nil && err != io.EOF {
			return am, fmt.Errorf("error: %s - reading: %s", err.Error(), string(raw))
		}
		if len(m) > 0 {
			am = append(am, m)
		}
		if err == io.EOF {
			break
		}
	}
	return am, nil
}
