package main

import (
	"fmt"

	"github.com/iniadil/go-slugify/slugify"
)

func main() {
	//slug, err := go_slugify.Slugify("Slugify - make your title better")
	//slug, _ := slugify.Slugify("Slugify - make !@#$%^&*()")
	//
	//fmt.Println(slug)
	//if err == nil {
	//}

	newSlug, _ := slugify.Slugify("Slugify - make !@#$%^&*()", slugify.WithLowerCase(false))
	fmt.Println(newSlug)
}
