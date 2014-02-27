// keyvalues_test.go - test keyvalues.go methods

package mxj

import (
	// "bytes"
	"fmt"
	// "io"
	"testing"
)

func TestKVHeader(t *testing.T) {
	fmt.Println("\n----------------  keyvalues_test.go ...\n")
}

var doc1 = []byte(`
<doc> 
   <books>
      <book seq="1">
         <author>William T. Gaddis</author>
         <title>The Recognitions</title>
         <review>One of the great seminal American novels of the 20th century.</review>
      </book>
      <book seq="2">
         <author>Austin Tappan Wright</author>
         <title>Islandia</title>
         <review>An example of earlier 20th century American utopian fiction.</review>
      </book>
      <book seq="3">
         <author>John Hawkes</author>
         <title>The Beetle Leg</title>
         <review>A lyrical novel about the construction of Ft. Peck Dam in Montana.</review>
      </book>
      <book seq="4"> 
         <author>
            <first_name>T.E.</first_name>
            <last_name>Porter</last_name>
         </author>
         <title>King's Day</title>
         <review>A magical novella.</review>
      </book>
   </books>
</doc>
`)

var doc2 = []byte(`
<doc>
   <books>
      <book seq="1">
         <author>William T. Gaddis</author>
         <title>The Recognitions</title>
         <review>One of the great seminal American novels of the 20th century.</review>
      </book>
   </books>
	<book>Something else.</book>
</doc>
`)

// the basic demo/test case - a small bibliography with mixed element types
func TestPathsForKey(t *testing.T) {
	fmt.Println("PathsForKey, doc1 ...")
	m, merr := NewMapXml(doc1)
	if merr != nil {
		t.Fatal("merr:", merr.Error())
	}
	fmt.Println("PathsForKey, doc1#author")
	ss := m.PathsForKey("author")
	fmt.Println("... ss:", ss)

	fmt.Println("PathsForKey, doc1#books")
	ss = m.PathsForKey("books")
	fmt.Println("... ss:", ss)

	fmt.Println("PathsForKey, doc2 ...")
	m, merr = NewMapXml(doc2)
	if merr != nil {
		t.Fatal("merr:", merr.Error())
	}
	fmt.Println("PathForKey, doc2#book")
	ss = m.PathsForKey("book")
	fmt.Println("... ss:", ss)

	fmt.Println("PathForKeyShortest, doc2#book")
	s := m.PathForKeyShortest("book")
	fmt.Println("... s :", s)
}

func TestValuesForKey(t *testing.T) {
	fmt.Println("ValuesForKey ...")
	m, merr := NewMapXml(doc1)
	if merr != nil {
		t.Fatal("merr:", merr.Error())
	}
	fmt.Println("ValuesForKey, doc1#author")
	ss, sserr := m.ValuesForKey("author")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForKey, doc1#book")
	ss, sserr = m.ValuesForKey("book")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForKey, doc1#book,-seq:3")
	ss, sserr = m.ValuesForKey("book", "-seq:3")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForKey, doc1#book, author:William T. Gaddis")
	ss, sserr = m.ValuesForKey("book", "author:William T. Gaddis")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForKey, doc1#author, -seq:1")
	ss, sserr = m.ValuesForKey("author", "-seq:1")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {	// should be len(ss) == 0
		fmt.Println("... ss.v:", v)
	}
}

func TestValuesForPath(t *testing.T) {
	fmt.Println("ValuesForPath ...")
	m, merr := NewMapXml(doc1)
	if merr != nil {
		t.Fatal("merr:", merr.Error())
	}
	fmt.Println("ValuesForPath, doc.books.book.author")
	ss, sserr := m.ValuesForPath("doc.books.book.author")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForPath, doc.books.book")
	ss, sserr = m.ValuesForPath("doc.books.book")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForPath, doc.books.book -seq=3")
	ss, sserr = m.ValuesForPath("doc.books.book", "-seq:3")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForPath, doc.books.* -seq=3")
	ss, sserr = m.ValuesForPath("doc.books.*", "-seq:3")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForPath, doc.*.* -seq=3")
	ss, sserr = m.ValuesForPath("doc.*.*", "-seq:3")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}
}

func TestValuesForNotKey( t *testing.T) {
	fmt.Println("ValuesForNotKey ...")
	m, merr := NewMapXml(doc1)
	if merr != nil {
		t.Fatal("merr:", merr.Error())
	}
	fmt.Println("ValuesForPath, doc.books.book !author:William T. Gaddis")
	ss, sserr := m.ValuesForPath("doc.books.book", "!author:William T. Gaddis")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForPath, doc.books.book !author:*")
	ss, sserr = m.ValuesForPath("doc.books.book", "!author:*")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {	// expect len(ss) == 0
		fmt.Println("... ss.v:", v)
	}

	fmt.Println("ValuesForPath, doc.books.book !unknown:*")
	ss, sserr = m.ValuesForPath("doc.books.book", "!unknown:*")
	if sserr != nil {
		t.Fatal("sserr:", sserr.Error())
	}
	for _, v := range ss {
		fmt.Println("... ss.v:", v)
	}
}
