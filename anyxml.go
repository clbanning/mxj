package mxj

// Encode arbitrary value as XML.  Note: there are no guarantees.
func AnyXml(v interface{}, rootTag ...string) ([]byte, error) {
	var err error
	s := new(string)
	p := new(pretty)

	var rt string
	if len(rootTag) == 1 {
		rt = rootTag[0]
	} else {
		rt = DefaultRootTag
	}

	var ss string
	var b []byte
	switch v.(type) {
	case []interface{}:
		ss = "<" + rt + ">"
		for _, vv := range v.([]interface{}) {
			switch vv.(type) {
			case map[string]interface{}:
				m := vv.(map[string]interface{})
				if len(m) == 1 {
					for tag, val := range m {
						err = mapToXmlIndent(false, s, tag, val, p)
					}
				} else {
					err = mapToXmlIndent(false, s, rt, vv, p)
				}
			default:
				err = mapToXmlIndent(false, s, rt, vv, p)
			}
			if err != nil {
				break
			}
		}
		ss += *s + "</" + rt + ">\n"
		b = []byte(ss)
	case map[string]interface{}:
		m := Map(v.(map[string]interface{}))
		b, err = m.Xml(rootTag...)
	default:
		err = mapToXmlIndent(false, s, rt, v, p)
		b = []byte(*s)
	}

	return b, err
}


// Encode an arbitrary value as a pretty XML string. Note: there are no guarantees.
func AnyXmlIndent(v interface{}, prefix, indent string, rootTag ...string) ([]byte, error) {
	var err error
	s := new(string)
	p := new(pretty)
	p.indent = indent
	p.padding = prefix

	var rt string
	if len(rootTag) == 1 {
		rt = rootTag[0]
	} else {
		rt = DefaultRootTag
	}

	var ss string
	var b []byte
	switch v.(type) {
	case []interface{}:
		ss = "<" + rt + ">\n"
		p.Indent()
		for _, vv := range v.([]interface{}) {
			switch vv.(type) {
			case map[string]interface{}:
				m := vv.(map[string]interface{})
				if len(m) == 1 {
					for tag, val := range m {
						err = mapToXmlIndent(true, s, tag, val, p)
					}
				} else {
					err = mapToXmlIndent(true, s, rt, vv, p)
				}
			default:
				err = mapToXmlIndent(true, s, rt, vv, p)
			}
			if err != nil {
				break
			}
		}
		ss += *s + "</" + rt + ">\n"
		b = []byte(ss)
	case map[string]interface{}:
		m := Map(v.(map[string]interface{}))
		b, err = m.XmlIndent(prefix, indent, rootTag...)
	default:
		err = mapToXmlIndent(true, s, rt, v, p)
		b = []byte(*s)
	}

	return b, err
}
