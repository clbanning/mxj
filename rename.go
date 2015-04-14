package mxj

import (
	"errors"
	"strings"
)

// RenameKey renames a key in a Map.
// It works only for nested maps. It doesn't work for cases when it buried in a list.
func (mv Map) RenameKey(path string, newName string) error {
	m := map[string]interface{}(mv)
	return renameKey_(m, path, newName)
}

func renameKey_(m interface{}, path string, newName string) error {
	keys := strings.Split(path, ".")

	switch mValue := m.(type) {
	case map[string]interface{}:
		for key, value := range mValue {
			if key == keys[0] {
				if len(keys) == 1 {
					// renaming
					mValue[newName] = value
					delete(mValue, key)
					return nil
				} else {
					// keep looking for the full path to the key
					return renameKey_(value, strings.Join(keys[1:], "."), newName)
				}
			}
		}
		// TODO(mrsln): look for the key buried in a list
	}
	return errors.New("RenameKey didn't find the path: " + path)
}
