package mxj

import (
	"strings"
	"testing"
)

// Map keys became XML tag/attribute names with no validation or escaping.
// JSON->XML conversion (mxj's main use case, via AnyXml) lets an attacker
// who controls a JSON object's keys inject raw XML structure.

func TestTagNameInjectionRejected(t *testing.T) {
	m := Map{
		`foo><injected attr="pwned`: "bar",
	}
	if _, err := m.Xml(); err == nil {
		t.Fatal("expected an error for an XML tag name containing structural characters, got nil")
	}
}

func TestAttrNameInjectionRejected(t *testing.T) {
	m := Map{
		"doc": map[string]interface{}{
			`-attr" x="evil`: "v",
			"#text":          "val",
		},
	}
	if _, err := m.Xml(); err == nil {
		t.Fatal("expected an error for an XML attribute name containing structural characters, got nil")
	}
}

func TestTagNameInjectionRejectedSeq(t *testing.T) {
	// same PoC through the #seq-aware encoding path (mapToXmlSeqIndent)
	mv, err := NewMapXmlSeq([]byte(`<doc><child>ok</child></doc>`))
	if err != nil {
		t.Fatalf("setup: %v", err)
	}
	doc := mv["doc"].(map[string]interface{})
	delete(doc, "child")
	doc[`evil"><injected`] = map[string]interface{}{"#seq": 0, "#text": "x"}
	if _, err := mv.Xml(); err == nil {
		t.Fatal("expected an error for an XML tag name containing structural characters (seq path), got nil")
	}
}

func TestValidTagAndAttrNamesStillWork(t *testing.T) {
	m := Map{
		"doc": map[string]interface{}{
			"-id":   "1",
			"#text": "val",
		},
	}
	x, err := m.Xml()
	if err != nil {
		t.Fatalf("valid map should encode without error: %v", err)
	}
	if !strings.Contains(string(x), `<doc id="1">val</doc>`) {
		t.Fatalf("unexpected output: %s", x)
	}
}
