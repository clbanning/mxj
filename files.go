package mxj

import (
	"fmt"
	"io"
	"os"
)

type MapRaw struct {
	M Map
	R []byte
}

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

// ReadMapsFromJsonFileRaw - creates an array of MapRaw from a file of JSON values.
func ReadMapsFromJsonFileRaw(name string) ([]MapRaw, error) {
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

	am := make([]MapRaw,0)
	for {
		mr := new(MapRaw)
		mr.M, mr.R, err = NewMapJsonReaderRaw(fh)
		if err != nil && err != io.EOF {
			return am, fmt.Errorf("error: %s - reading: %s", err.Error(), string(mr.R))
		}
		if len(mr.M) > 0 {
			am = append(am, *mr)
		}
		if err == io.EOF {
			break
		}
	}
	return am, nil
}

// ReadMapsFromXmlFile - creates an array from a file of XML values.
func ReadMapsFromXmlFile(name string) ([]Map, error) {
	x := XmlWriterBufSize
	XmlWriterBufSize = 0
	defer func() {
		XmlWriterBufSize = x
	}()

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

// ReadMapsFromXmlFileRaw - creates an array of MapRaw from a file of XML values.
// NOTE: the slice with the raw XML is clean with no extra capacity - unlike NewMapXmlReaderRaw().
// It is slow at parsing a file from disk and is intended for relatively small utility files.
func ReadMapsFromXmlFileRaw(name string) ([]MapRaw, error) {
	x := XmlWriterBufSize
	XmlWriterBufSize = 0
	defer func() {
		XmlWriterBufSize = x
	}()


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

	am := make([]MapRaw,0)
	for {
		mr := new(MapRaw)
		mr.M, mr.R, err = NewMapXmlReaderRaw(fh)
		if err != nil && err != io.EOF {
			return am, fmt.Errorf("error: %s - reading: %s", err.Error(), string(mr.R))
		}
		if len(mr.M) > 0 {
			am = append(am, *mr)
		}
		if err == io.EOF {
			break
		}
	}
	return am, nil
}
