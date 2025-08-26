package helper

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
)

func checkString(s string) error {
	// check string
	if len(s) == 0 {
		return errors.New("empty string")
	}

	if reflect.TypeOf(s).Kind() != reflect.String {
		return errors.New("not a string")
	}
	return nil
}

func Slugify(s string) (string, error) {
	if len(s) < 1 {
		return "", errors.New("slugify must contain at least 1 character")
	}

	if err := checkString(s[:1]); err != nil {
		return "", err
	}

	// remove end spaces
	trimmed := strings.TrimRight(s, " ")

	// changes space with -
	betaSlug := strings.ToLower(trimmed)
	// remove symbols
	cleanedText := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			return r // Keep letters, digits, and spaces
		}
		return -1 // Drop other characters (symbols)
	}, betaSlug)
	cleanedText = strings.ReplaceAll(cleanedText, "_", "")
	cleanedText = strings.ReplaceAll(cleanedText, "-", "")
	cleanedText = strings.ReplaceAll(cleanedText, "  ", "-")
	cleanedText = strings.ReplaceAll(cleanedText, " ", "-")

	// if character > 40
	if len(cleanedText) > 40 {
		cleanedText = cleanedText[:40]
	}

	return cleanedText, nil
}
