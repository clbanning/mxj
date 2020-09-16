// misc_test.go

package mxj

import (
	"fmt"
	"testing"
)

var miscdata = []byte(`
<doc>
	<elem1 name="elem1" seq="1">
		<sub1 name="sub1" seq="1">sub_value_1</sub1>
		<sub2 name="sub2" seq="2">sub_value_2</sub2>
	</elem1>
	<elem2 name="elem2" seq="2">element_2</elem2>
</doc>
`)

func TestMisc(t *testing.T) {
	PrependAttrWithHyphen(true) // be safe
	fmt.Println("\n------------------ misc_test.go ...")
}

func TestRoot(t *testing.T) {
	m, err := NewMapXml(miscdata)
	if err != nil {
		t.Fatal(err)
	}
	r, err := m.Root()
	if err != nil {
		t.Fatal(err)
	}
	if r != "doc" {
		t.Fatal("Root not doc:", r)
	}
}

func TestElements(t *testing.T) {
	m, err := NewMapXml(miscdata)
	if err != nil {
		t.Fatal(err)
	}
	e, err := m.Elements("doc")
	if err != nil {
		t.Fatal(err)
	}
	elist := []string{"elem1", "elem2"}
	for i, ee := range e {
		if ee != elist[i] {
			t.Fatal("error in list, elem#:", i, "-", ee, ":", elist[i])
		}
	}

	e, err = m.Elements("doc.elem1")
	if err != nil {
		t.Fatal(err)
	}
	elist = []string{"sub1", "sub2"}
	for i, ee := range e {
		if ee != elist[i] {
			t.Fatal("error in list, elem#:", i, "-", ee, ":", elist[i])
		}
	}
}

func TestAttributes(t *testing.T) {
	m, err := NewMapXml(miscdata)
	if err != nil {
		t.Fatal(err)
	}
	a, err := m.Attributes("doc.elem2")
	if err != nil {
		t.Fatal(err)
	}
	alist := []string{"name", "seq"}
	for i, aa := range a {
		if aa != alist[i] {
			t.Fatal("error in list, elem#:", i, "-", aa, ":", alist[i])
		}
	}

	a, err = m.Attributes("doc.elem1.sub2")
	if err != nil {
		t.Fatal(err)
	}
	alist = []string{"name", "seq"}
	for i, aa := range a {
		if aa != alist[i] {
			t.Fatal("error in list, elem#:", i, "-", aa, ":", alist[i])
		}
	}
}

func TestElementsAttrPrefix(t *testing.T) {
	SetAttrPrefix("__")
	m, err := NewMapXml(miscdata)
	if err != nil {
		t.Fatal(err)
	}
	e, err := m.Elements("doc")
	if err != nil {
		t.Fatal(err)
	}
	elist := []string{"elem1", "elem2"}
	for i, ee := range e {
		if ee != elist[i] {
			t.Fatal("error in list, elem#:", i, "-", ee, ":", elist[i])
		}
	}

	e, err = m.Elements("doc.elem1")
	if err != nil {
		t.Fatal(err)
	}
	elist = []string{"sub1", "sub2"}
	for i, ee := range e {
		if ee != elist[i] {
			t.Fatal("error in list, elem#:", i, "-", ee, ":", elist[i])
		}
	}
}

func TestAttributesAttrPrefix(t *testing.T) {
	SetAttrPrefix("__")
	m, err := NewMapXml(miscdata)
	if err != nil {
		t.Fatal(err)
	}
	a, err := m.Attributes("doc.elem2")
	if err != nil {
		t.Fatal(err)
	}
	alist := []string{"name", "seq"}
	for i, aa := range a {
		if aa != alist[i] {
			t.Fatal("error in list, elem#:", i, "-", aa, ":", alist[i])
		}
	}

	a, err = m.Attributes("doc.elem1.sub2")
	if err != nil {
		t.Fatal(err)
	}
	alist = []string{"name", "seq"}
	for i, aa := range a {
		if aa != alist[i] {
			t.Fatal("error in list, elem#:", i, "-", aa, ":", alist[i])
		}
	}
}

func TestElementsNoAttrPrefix(t *testing.T) {
	PrependAttrWithHyphen(false)
	m, err := NewMapXml(miscdata)
	if err != nil {
		t.Fatal(err)
	}
	e, err := m.Elements("doc")
	if err != nil {
		t.Fatal(err)
	}
	if len(e) != 2 {
		t.Fatal("didn't get 2 elements:", e)
	}

	e, err = m.Elements("doc.elem1")
	if err != nil {
		t.Fatal(err)
	}
	if len(e) != 4 {
		t.Fatal("didn't get 4 elements:", e)
	}
	PrependAttrWithHyphen(true)
}

func TestAttributesNoAttrPrefix(t *testing.T) {
	PrependAttrWithHyphen(false)
	m, err := NewMapXml(miscdata)
	if err != nil {
		t.Fatal(err)
	}
	a, err := m.Attributes("doc.elem2")
	if err != nil {
		t.Fatal(err)
	}
	if len(a) > 0 {
		t.Fatal("found attributes where there are none:", a)
	}

	a, err = m.Attributes("doc.elem1.sub2")
	if err != nil {
		t.Fatal(err)
	}
	if len(a) > 0 {
		t.Fatal("found attributes where there are none:", a)
	}
	PrependAttrWithHyphen(true)
}

var whiteSpaceData = []byte("<doc><elem3> hello world </elem3></doc>")

func TestPreserveSpaceDisableByDefault(t *testing.T) {
	const (
		path       = "doc.elem3"
		input      = " hello world "
		trimmed    = "hello world"
		notTrimmed = input
	)
	tcs := []struct {
		name           string
		args           []bool
		expectedOutput string
	}{
		{
			name:           "Default Preserve Space disabled should trim values",
			args:           nil, // nil will result in the `DisableWhiteSpace` to be skipped
			expectedOutput: trimmed,
		},
		{
			name:           "Single true is passed should not trim values",
			args:           []bool{true},
			expectedOutput: notTrimmed,
		},
		{
			name:           "Single false is passed should trim values",
			args:           []bool{false},
			expectedOutput: trimmed,
		},
		{
			name:           "No args are passed should not trim values",
			args:           []bool{},
			expectedOutput: notTrimmed,
		},
		{
			name:           "Extra arguments should be ignored",
			args:           []bool{true, false},
			expectedOutput: notTrimmed,
		},
		{
			name:           "Extra arguments should be ignored with false",
			args:           []bool{false, true},
			expectedOutput: trimmed,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.args != nil {
				DisableTrimWhiteSpace(tc.args...)
			}
			m, err := NewMapXml(whiteSpaceData)

			if err != nil {
				t.Fatal(err)
			}

			s, err := m.ValueForPath(path)
			if err != nil {
				t.Fatal(err)
			}
			if s != tc.expectedOutput {
				t.Fatalf("expected '%s' got '%s'", tc.expectedOutput, s)
			}
		})
	}
	// Set it back to false after all tests are done
	DisableTrimWhiteSpace(false)
}
