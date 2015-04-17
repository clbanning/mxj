package mxj

import "strings"

// Removes the path.
func (mv Map) Remove(path string) error {
	m := map[string]interface{}(mv)
	return remove(m, path)
}

func remove(m interface{}, path string) error {
	val, err := prevValueByPath(m, path)
	if err != nil {
		return err
	}

	keys := strings.Split(path, ".")
	key := keys[len(keys)-1]

	delete(val, key)
	return nil
}
