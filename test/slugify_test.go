package test

import (
	"testing"

	"github.com/iniadil/go-slugify/slugify"
)

func BenchmarkSlugify_Table(b *testing.B) {
	cases := []struct {
		name string
		in   string
		opts []slugify.Options // atau []func(*config) tergantung penamaanmu
	}{
		{"short_default", "Halo Dunia", nil},
		{"long_default", "…panjang sekali …", nil},
		{"custom_sep", "Halo___Dunia", []slugify.Options{slugify.WithSeparator("_")}},
		{"no_lower", "Halo Dunia", []slugify.Options{slugify.WithLowerCase(false)}},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = slugify.Slugify(tc.in, tc.opts...)
			}
		})
	}
}
