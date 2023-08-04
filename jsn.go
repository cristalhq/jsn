package jsn

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
