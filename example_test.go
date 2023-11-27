package jsn_test

import (
	"encoding/json"
	"fmt"

	"github.com/cristalhq/jsn"
)

func Example() {
	j := jsn.O{
		"hello": "world",
		"do": jsn.A{
			"go", "mod", "tidy",
		},
		"x": jsn.N("123456.00000000000000000000000000000000000000001"),
	}

	raw, _ := json.MarshalIndent(j, "", "  ")
	fmt.Printf("%s\n", raw)

	// Output:
	//{
	//   "do": [
	//     "go",
	//     "mod",
	//     "tidy"
	//   ],
	//   "hello": "world",
	//   "x": 123456.00000000000000000000000000000000000000001
	//}
}

func ExampleAO() {
	j := jsn.AO{
		{"a": 1},
		{"b": 2, "c": 3},
		{"d": jsn.O{"1": "2"}},
	}

	raw, _ := json.MarshalIndent(j, "", "  ")
	fmt.Printf("%s\n", raw)

	// Output:
	// [
	//   {
	//     "a": 1
	//   },
	//   {
	//     "b": 2,
	//     "c": 3
	//   },
	//   {
	//     "d": {
	//       "1": "2"
	//     }
	//   }
	// ]
}

func ExampleNumber() {
	j := jsn.O{
		"x": 123456.00000000000000000000000000000000000000001,
		"y": jsn.N("123456.00000000000000000000000000000000000000001"),
		"z": "123456.00000000000000000000000000000000000000001",
	}

	raw, _ := json.MarshalIndent(j, "", "  ")
	fmt.Printf("%s\n", raw)

	// Output:
	//{
	//   "x": 123456,
	//   "y": 123456.00000000000000000000000000000000000000001,
	//   "z": "123456.00000000000000000000000000000000000000001"
	//}
}

func ExampleIndent() {
	j := jsn.O{
		"x": 123456.00000000000000000000000000000000000000001,
		"y": map[string]any{
			"a": []int{1, 2, 3},
			"b": "ccc",
		},
	}

	fmt.Printf("%s\n", jsn.Indent(j))

	// Output:
	// {
	//   "x": 123456,
	//   "y": {
	//     "a": [
	//       1,
	//       2,
	//       3
	//     ],
	//     "b": "ccc"
	//   }
	// }
}
