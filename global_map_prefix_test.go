package mxj

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var whiteSpaceDataSeqTest2 = []byte(`<books>
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

func TestSetGlobalKeyMapPrefix(t *testing.T) {
	prefixList := []struct {
		name  string
		value string
	}{
		{
			name:  "Testing with % as Map Key Prefix",
			value: "%",
		},
		{
			name:  "Testing with _ as Map Key Prefix",
			value: "_",
		},
		{
			name:  "Testing with - as Map Key Prefix",
			value: "-",
		},
		{
			name:  "Testing with & as Map Key Prefix",
			value: "&",
		},
	}

	for _, prefix := range prefixList {
		t.Run(prefix.name, func(t *testing.T) {

			// Testing MapSeq(Ordering) with whitespace and byte equivalence
			DisableTrimWhiteSpace(true)
			SetGlobalKeyMapPrefix(prefix.value)

			m, err := NewMapFormattedXmlSeq(whiteSpaceDataSeqTest2)
			if err != nil {
				t.Fatal(err)
			}

			m1 := MapSeq(m)
			x, err := m1.XmlIndent("", "   ")
			if err != nil {
				t.Fatal(err)
			}

			if string(x) != string(whiteSpaceDataSeqTest2) {
				t.Fatalf("expected\n'%s' \ngot \n'%s'", whiteSpaceDataSeqTest2, x)
			}
			DisableTrimWhiteSpace(false)

			// Testing Map with whitespace and deep equivalence
			DisableTrimWhiteSpace(true)
			m3, err := NewMapXml(whiteSpaceDataSeqTest2)
			if err != nil {
				t.Fatal(err)
			}

			m4 := Map(m3)

			if !cmp.Equal(m3, m4) {
				t.Errorf("Maps unmatched using %s", prefix.value)
			}
			DisableTrimWhiteSpace(false)

			// Testing MapSeq(Ordering) without whitespace and byte equivalence
			m5, err := NewMapFormattedXmlSeq(whiteSpaceDataSeqTest2)
			if err != nil {
				t.Fatal(err)
			}

			m6 := MapSeq(m5)

			if !cmp.Equal(m5, m6) {
				t.Errorf("Maps unmatched using %s", prefix.value)
			}

			// Testing Map without whitespace and deep equivalence
			m7, err := NewMapXml(whiteSpaceDataSeqTest2)
			if err != nil {
				t.Fatal(err)
			}

			m8 := Map(m7)

			if !cmp.Equal(m7, m8) {
				t.Errorf("Maps unmatched using %s", prefix.value)
			}
		})

	}
	SetGlobalKeyMapPrefix("#")

}
