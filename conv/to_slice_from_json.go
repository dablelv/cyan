package conv

import "encoding/json"

// JSONToSliceE converts the JSON-encoded data to []string.
// E.g. a JSON value ["foo", "bar", "baz"]
func JSONToSliceE(data []byte) ([]string, error) {
	var strs []string
	err := json.Unmarshal(data, &strs)
	return strs, err
}
