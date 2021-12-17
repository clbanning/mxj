
// ======================== newMapToXmlIndent

func (mv Map) MarshalXml(rootTag ...string) ([]byte, error) {
	m := map[string]interface{}(mv)
	var err error
	// s := new(string)
	// b := new(strings.Builder)
	b := new(bytes.Buffer)
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
						err = marshalMapToXmlIndent(false, b, DefaultRootTag, m, p)
						goto done
					}
				}
			}
			err = marshalMapToXmlIndent(false, b, key, value, p)
		}
	} else if len(rootTag) == 1 {
		err = marshalMapToXmlIndent(false, b, rootTag[0], m, p)
	} else {
		err = marshalMapToXmlIndent(false, b, DefaultRootTag, m, p)
	}
done:
	return b.Bytes(), err
}
