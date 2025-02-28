package jsn_test

import (
	"encoding/json"
	"fmt"

	"github.com/cristalhq/jsn"
)

func ExampleO() {
	j := jsn.O{
		"hello": "world",
		"do": jsn.A{
			"go", "mod", "tidy",
		},
		"x": jsn.N("123456.00000000000000000000000000000000000000001"),
	}

	printJSON(j)

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

	printJSON(j)

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

func ExampleN() {
	j := jsn.O{
		"x": 123456.00000000000000000000000000000000000000001,
		"y": jsn.N("123456.00000000000000000000000000000000000000001"),
		"z": "123456.00000000000000000000000000000000000000001",
	}

	printJSON(j)

	// Output:
	//{
	//   "x": 123456,
	//   "y": 123456.00000000000000000000000000000000000000001,
	//   "z": "123456.00000000000000000000000000000000000000001"
	//}
}

func ExampleF() {
	nums := []jsn.F{
		jsn.F(0),
		jsn.F(1),
		jsn.F(10),
		jsn.F(10.12),
		jsn.F(-1),
		jsn.F(-1.23),
		jsn.F(-0.1),
	}

	for _, f := range nums {
		printJSON(f)
	}

	// Output:
	// 0.0
	// 1.0
	// 10.0
	// 10.12
	// -1.0
	// -1.23
	// -0.1
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

func printJSON(v any) {
	raw, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(raw))
}
