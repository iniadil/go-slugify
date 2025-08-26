package main

import (
	"fmt"

	go_slugify "github.com/iniadil/go-slugify/helper"
)

func main() {
	//slug, err := go_slugify.Slugify("Slugify - make your title better")
	slug, err := go_slugify.Slugify("Slugify - make !@#$%^&*()")

	if err == nil {
		fmt.Println(slug)
	}
}
