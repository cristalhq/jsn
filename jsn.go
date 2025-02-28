package jsn

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
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

// F represents JSON number with always presented decimal part.
type F float64

func (f F) MarshalJSON() ([]byte, error) {
	s := strconv.FormatFloat(float64(f), 'f', -1, 64)

	if float64(f) == float64(int64(f)) {
		s += ".0"
	}
	return []byte(s), nil
}

// Unmarshal only 1 JSON entity from the input.
// Disallows unknown fields if the argument is a struct.
func Unmarshal(r io.Reader, v any) error {
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

// MustMarshal an object to JSON, and panics if any error occurs.
func MustMarshal(obj any) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return b
}

// Indent is like [json.MarshalIndent] but with predefined params.
// Panics if any error occurs.
func Indent(obj any) string {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

var errMore = errors.New("body must contain only one JSON entity")
