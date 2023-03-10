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
