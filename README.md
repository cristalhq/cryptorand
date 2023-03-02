# cryptorand

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Go package provides `math/rand.Source64` implementation based on `crypto/rand`.

## Features

* Simple API.
* Dependency-free.
* Clean and tested code.

## Install

Go version 1.19+

```
go get github.com/cristalhq/cryptorand
```

## Example

```go
r := rand.New(cryptorand.Source)

m := map[string]struct{}{}
for i := 0; i < 100; i++ {
	s := fmt.Sprint(r.Float64())
	m[s] = struct{}{}
}

fmt.Printf("Have %d unique floats", len(m))

// Output:
// Have 100 unique floats
```

See examples: [example_test.go](example_test.go).

## Documentation

See [these docs][pkg-url] for more details.

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/cryptorand/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/cryptorand/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/cryptorand
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/cryptorand
[reportcard-img]: https://goreportcard.com/badge/cristalhq/cryptorand
[reportcard-url]: https://goreportcard.com/report/cristalhq/cryptorand
[coverage-img]: https://codecov.io/gh/cristalhq/cryptorand/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/cryptorand
[version-img]: https://img.shields.io/github/v/release/cristalhq/cryptorand
[version-url]: https://github.com/cristalhq/cryptorand/releases
