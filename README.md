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

## License

See [LICENSE](LICENSE) for details.

---
Maintained by [iniadil](https://github.com/iniadil)