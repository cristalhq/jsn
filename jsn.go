package jsn

// A represents JSON array.
type A = []any

// O represents JSON object.
type O = map[string]any

// N represents JSON number.
type N string

func (n N) MarshalJSON() ([]byte, error) {
	return []byte(n), nil
}
