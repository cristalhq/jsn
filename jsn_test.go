package jsn

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	testCases := []struct {
		input  string
		want   any
		errStr string
	}{
		{
			input: `"abc"`,
			want:  "abc",
		},
		{
			input: `123`,
			want:  float64(123),
		},
		{
			input: `{"abc": 123}`,
			want:  map[string]any{"abc": float64(123)},
		},
		{
			input: `{"abc": 123} `,
			want:  map[string]any{"abc": float64(123)},
		},
		{
			input: `        ["abc", 123] `,
			want:  []any{"abc", float64(123)},
		},
		{
			input:  `{"abc": 123}a`,
			errStr: "body must contain only one JSON object",
		},
		{
			input:  `{"abc`,
			errStr: "unexpected EOF",
		},
	}

	for _, tc := range testCases {
		var val any
		err := Unmarshal([]byte(tc.input), &val)
		if err != nil {
			if tc.errStr == "" {
				t.Fatal(err)
			}
			if have := err.Error(); have != tc.errStr {
				t.Fatalf("\nhave: %+v\nwant: %+v\n", have, tc.errStr)
			}
			continue
		}

		if !reflect.DeepEqual(val, tc.want) {
			t.Fatalf("\nhave: %+v\nwant: %+v\n", val, tc.want)
		}
	}
}
