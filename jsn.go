package jsn

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// A represents JSON array.
type A = []any

// O represents JSON object.
type O = map[string]any

// AO represents JSON array of JSON objects.
type AO = []O

// N represents JSON number.
type N string

func (n N) MarshalJSON() ([]byte, error) {
	return []byte(n), nil
}

// Unmarshal only 1 JSON entity from the input.
// Disallows unknown fields if the argument is a struct.
func Unmarshal(b []byte, v any) error {
	return UnmarshalFrom(bytes.NewReader(b), v)
}

// UnmarshalFrom only 1 JSON entity from the input.
// Disallows unknown fields if the argument is a struct.
func UnmarshalFrom(r io.Reader, v any) error {
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()

	if err := d.Decode(v); err != nil {
		return err
	}
	if d.More() {
		return errMore
	}
	return nil
}

var errMore = errors.New("body must contain only one JSON entity")
