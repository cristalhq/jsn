# jsn

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Go package to easily construct JSON without defined types.

## Features

* Very simple API.
* Dependency-free.
* Clean and tested code.

## Install

Go version 1.18+

```
go get github.com/cristalhq/jsn
```

## Example

```go
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
```

See examples: [example_test.go](example_test.go).

## Documentation

See [these docs][pkg-url] for more details.

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/jsn/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/jsn/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/jsn
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/jsn
[reportcard-img]: https://goreportcard.com/badge/cristalhq/jsn
[reportcard-url]: https://goreportcard.com/report/cristalhq/jsn
[coverage-img]: https://codecov.io/gh/cristalhq/jsn/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/jsn
[version-img]: https://img.shields.io/github/v/release/cristalhq/jsn
[version-url]: https://github.com/cristalhq/jsn/releases
