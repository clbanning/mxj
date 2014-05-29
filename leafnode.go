package mxj

// leafnode.go - return leaf nodes with paths and values for the Map
// inspired by: https://groups.google.com/forum/#!topic/golang-nuts/3JhuVKRuBbw

import (
	"strconv"
)

// LeafNode - a terminal path value in a Map.
type LeafNode struct {
	Path  string      // a dot-notation representation of the path
	Value interface{} // the leaf value
}

// LeafNodes - returns an array of all LeafNode values for the Map
func (mv Map)LeafNodes() []LeafNode {
	l := make([]LeafNode, 0)
	getLeafNotes("", "", map[string]interface{}(mv), &l)
	return l
}

func getLeafNotes(path, node string, mv interface{}, l *[]LeafNode) {
	if path != "" {
		path += "."
	}
	path += node
	switch mv.(type) {
	case map[string]interface{}:
		for k, v := range mv.(map[string]interface{}) {
			getLeafNotes(path, k, v, l)
		}
	case []interface{}:
		for i, v := range mv.([]interface{}) {
			getLeafNotes(path, strconv.Itoa(i), v, l)
		}
	default:
		// can't walk any further, so create leaf
		n := LeafNode{path, mv}
		*l = append(*l, n)
	}
}

// LeafPaths - all paths that terminate in LeafNode values.
func (mv Map) LeafPaths() []string {
	ln := mv.LeafNodes()
	ss := make([]string,len(ln))
	for i := 0 ; i < len(ln); i++ {
		ss[i] = ln[i].Path
	}
	return ss
}

// LeafValues - all terminal values in the Map.
func (mv Map) LeafValues() []interface{} {
	ln := mv.LeafNodes()
	vv := make([]interface{},len(ln))
	for i := 0 ; i < len(ln); i++ {
		vv[i] = ln[i].Value
	}
	return vv
}
