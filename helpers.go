package api

import (
	"bytes"
	"encoding/json"
)

// ExactMatch ..
func ExactMatch(x, y []int64) bool {
	return len(NotExistsIn(x, y)) > 0 && len(NotExistsIn(y, x)) > 0
}

// NotExistsIn ..
func NotExistsIn(x, y []int64) []int64 {
	var diff []int64

	m := make(map[int64]bool)

	for _, i := range x {
		m[i] = true
	}

	for _, i := range y {
		if _, ok := m[i]; !ok {
			diff = append(diff, i)
		}
	}

	return diff
}

// UnescapedMarshal ..
func UnescapedMarshal(data interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}

	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(data)

	return buf.Bytes(), err
}
