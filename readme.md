<h2>mxj - to/from maps, XML and JSON</h2>
Marshal/Unmarshal XML to/from JSON and map[string]interface{} values, and extract/modify values from maps by key or key-path, including wildcards.  

mxj supplants the legacy x2j and j2x packages. If you want the old syntax, use mxj/x2j and mxj/j2x packages.

<h4>Notices</h4>
	2014-11-09: IncludeTagSeqNum() adds "_seq" key with XML doc positional information.
	            (NOTE: PreserveXmlList() is similar and will be here soon.)
	2014-09-18: inspired by NYTimes fork, added PrependAttrWithHyphen() to allow stripping hyphen from attribute tag.
	2014-08-02: AnyXml() and AnyXmlIndent() will try to marshal arbitrary values to XML.
	2014-04-28: ValuesForPath() and NewMap() now accept path with indexed array references.

<h4>Basic Unmarshal XML / JSON / struct</h4>
   - type Map map[string]interface{}

Create a Map value, 'm', from any map[string]interface{} value, 'v':
   - m := Map(v)

Unmarshal / marshal XML as a Map value, 'm':
   - m, err := NewMapXml(xmlValue) // unmarshal
   - xmlValue, err := m.Xml()      // marshal

Unmarshal XML from an io.Reader as a Map value, 'm':
   - m, err := NewMapReader(xmlReader)         // repeated calls, as with an os.File Reader, will process stream
   - m, raw, err := NewMapReaderRaw(xmlReader) // 'raw' is the raw XML that was decoded

Marshal Map value, 'm', to an XML Writer (io.Writer):
   - err := m.XmlWriter(xmlWriter)<br>
   - raw, err := m.XmlWriterRaw(xmlWriter) // 'raw' is the raw XML that was written on xmlWriter
	
Also, for prettified output:
   - xmlValue, err := m.XmlIndent(prefix, indent, ...)<br>
   - err := m.XmlIndentWriter(xmlWriter, prefix, indent, ...)<br>
   - raw, err := m.XmlIndentWriterRaw(xmlWriter, prefix, indent, ...)

Bulk process XML with error handling (note: handlers must return a boolean value):
   - err := HandleXmlReader(xmlReader, mapHandler(Map), errHandler(error))<br>
   - err := HandleXmlReaderRaw(xmlReader, mapHandler(Map, []byte), errHandler(error, []byte))

Converting XML to JSON: see Examples for NewMapXml and HandleXmlReader.

There are comparable functions and methods for JSON processing.

Arbitrary structure values can be decoded to / encoded from Map values:
   - m, err := NewMapStruct(structVal)<br>
   - err := m.Struct(structPointer)

<h4>Extract / modify Map values</h4>
To work with XML tag values, JSON or Map key values or structure field values, decode the XML, JSON
or structure to a Map value, 'm', or cast a map[string]interface{} value to a Map value, 'm', then:
   - paths := m.PathsForKey(key)<br>
   - path := m.PathForKeyShortest(key)<br>
   - values, err := m.ValuesForKey(key, subkeys)<br>
   - values, err := m.ValuesForPath(path, subkeys)<br>
   - count, err := m.UpdateValuesForPath(newVal, path, subkeys)

Get everything at once, irrespective of path depth:
   - leafnodes := m.LeafNodes()<br>
   - leafvalues := m.LeafValues()

A new Map with whatever keys are desired can be created from the current Map and then encoded in XML
or JSON. (Note: keys can use dot-notation.)
   - newMap := m.NewMap("oldKey_1:newKey_1", "oldKey_2:newKey_2", ..., "oldKey_N:newKey_N")<br>
   - newXml := newMap.Xml()   // for example<br>
   - newJson := newMap.Json() // ditto<br>

<h4>Usage</h4>

The package is fairly well self-documented with examples. (http://godoc.org/github.com/clbanning/mxj)

Also, the subdirectory "examples" contains a wide range of examples, several taken from golang-nuts discussions.

<h4>XML parsing conventions</h4>

   - Attributes are parsed to map[string]interface{} values by prefixing a hyphen, '-',
     to the attribute label. (Unless overridden by PrependAttrWithHyphen(false).)
   - If the element is a simple element and has attributes, the element value
     is given the key '#text' for its map[string]interface{} representation.  (See
     the 'atomFeedString.xml' test data, below.)

<h4>XML encoding conventions</h4>

   - 'nil' Map values, which may represent 'null' JSON values, are encoded as '\<tag/\>'.
      NOTE: the operation is not symmetric as '\<tag/\>' elements are decoded as 'tag:""' Map values,
            which, then, encode in JSON as '"tag":""' values..

<h4>Running "go test"</h4>

Because there are no guarantees on the sequence map elements are retrieved, the tests have been 
written for visual verification in most cases.  One advantage is that you can easily use the 
output from running "go test" as examples of calling the various functions and methods.

<h4>Motivation</h4>

I make extensive use of JSON for messaging and typically unmarshal the messages into
map[string]interface{} variables.  This is easily done using json.Unmarshal from the
standard Go libraries.  Unfortunately, many legacy solutions use structured
XML messages; in those environments the applications would have to be refitted to
interoperate with my components.

The better solution is to just provide an alternative HTTP handler that receives
XML messages and parses it into a map[string]interface{} variable and then reuse
all the JSON-based code.  The Go xml.Unmarshal() function does not provide the same
option of unmarshaling XML messages into map[string]interface{} variables. So I wrote
a couple of small functions to fill this gap and released them as the x2j package.

Over the next year and a half additional features were added, and the companion j2x
package was released to address XML encoding of arbitrary JSON and map[string]interface{}
values.  As part of a refactoring of our production system and looking at how we had been
using the x2j and j2x packages we found that we rarely performed direct XML-to-JSON or
JSON-to_XML conversion and that working with the XML or JSON as map[string]interface{}
values was the primary value.  Thus, everything was refactored into the mxj package.

