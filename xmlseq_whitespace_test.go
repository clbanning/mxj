package mxj

import "testing"

var whiteSpaceDataSeqTest = []byte(`<books>
   <book seq="1" ser="5">
      <author>William T. Gaddis </author>
      <title> The Recognitions </title>
      <review> One of the great seminal American novels of the 20th century.</review>
   </book>
   <book seq="2">
      <author>Austin Tappan Wright</author>
      <title>Islandia</title>
      <review>An example of earlier 20th century American utopian fiction.</review>
   </book>
   <book seq="3" ser="6">
      <author> John Hawkes </author>
      <title> The Beetle Leg </title>
      <review> A lyrical novel about the construction of Ft. Peck Dam in Montana. </review>
   </book>
</books>`)

func TestNewMapXmlSeqWhiteSpace(t *testing.T) {
	t.Run("Testing NewMapFormattedXmlSeq with WhiteSpacing", func(t *testing.T) {
		DisableTrimWhiteSpace(true)

		m, err := NewMapFormattedXmlSeq(whiteSpaceDataSeqTest)
		if err != nil {
			t.Fatal(err)
		}

		m1 := MapSeq(m)
		x, err := m1.XmlIndent("", "   ")
		if err != nil {
			t.Fatal(err)
		}

		if string(x) != string(whiteSpaceDataSeqTest) {
			t.Fatalf("expected\n'%s' \ngot \n'%s'", whiteSpaceDataSeqTest, x)
		}
	})
	DisableTrimWhiteSpace(false)
}
