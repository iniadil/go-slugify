package test

import (
	"strconv"
	"testing"

	"github.com/iniadil/go-slugify/slugify"
)

func BenchmarkSlugify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		title := "test" + strconv.Itoa(i)
		slug, _ := slugify.Slugify(title)
	}
}
