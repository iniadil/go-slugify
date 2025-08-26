# go-slugify

A simple Go package to convert strings into URL-friendly slugs.

## Install

```bash
go get github.com/iniadil/go-slugify@latest
```

## Usage

Here is a basic usage example:

```go
import "github.com/iniadil/go-slugify/slugify"

slug, err := slugify.Slugify("Slugify - make !@#$%^&*()")
if err == nil {
	fmt.Println(slug) // Output: slugify-make
}
```

## With Separator
```go
// Using a custom separator
newSlug, _ := slugify.Slugify("a title with separator", slugify.WithSeparator("/"))
fmt.Println(newSlug) // Output: slugify/make
```

## API

- `Slugify(text string, opts ...Option) (string, error)`
    - Converts the input string to a slug.
    - Options:
        - `WithSeparator(sep string)` : set a custom separator.
        - `WithLowerCase(bool)` : set lowercase _(default true)_

## Benchmark

`go-slugify` has been benchmarked to demonstrate its performance.  
On a macOS machine with an Intel i7 CPU, the results show:

- Short strings: ~300 ns/op
- Long strings: ~580 ns/op
- Custom separator: ~320 ns/op
- Without lowercase conversion: ~228 ns/op

These results indicate that:

- **Fast**: All operations complete in sub-microsecond time.
- **Lightweight**: Minimal overhead, suitable for high-performance applications.
- **Consistent**: Options like custom separators or disabling lowercase do not introduce significant cost.

This makes `go-slugify` a reliable choice for projects that require efficient slug generation, such as REST APIs, microservices, or CLI tools.



## License

See [LICENSE](LICENSE) for details.

---
Maintained by [iniadil](https://github.com/iniadil)