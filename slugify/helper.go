package slugify

import (
	"errors"
	"reflect"
	"strconv"
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

func separatorChecker(separator string) error {
	if separator == "" {
		return errors.New("empty separator")
	}

	if len(separator) < 1 {
		return errors.New("empty separator")
	}

	maxLen := 5
	if len(separator) > maxLen {
		errorMsg := "separator too long, max: " + strconv.Itoa(maxLen) + ", got: " + separator
		return errors.New(errorMsg)
	}

	return nil
}

type config struct {
	Separator string
	LowerCase bool
}

type options func(*config)

func WithSeparator(sep string) options {
	return func(c *config) {
		c.Separator = sep
	}
}

func WithLowerCase(enable bool) options {
	return func(c *config) {
		c.LowerCase = enable
	}
}

func Slugify(s string, option ...options) (string, error) {
	if len(s) < 1 {
		return "", errors.New("slugify must contain at least 1 character")
	}

	if err := checkString(s[:1]); err != nil {
		return "", err
	}

	cfg := &config{
		Separator: "-",
		LowerCase: true,
	}

	for _, opt := range option {
		opt(cfg)
	}

	// remove end spaces
	trimmed := strings.TrimRight(s, " ")

	// changes space with -
	var betaSlug string
	if cfg.LowerCase {
		betaSlug = strings.ToLower(trimmed)
	} else {
		betaSlug = trimmed
	}
	// remove symbols
	cleanedText := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			return r // Keep letters, digits, and spaces
		}
		return -1 // Drop other characters (symbols)
	}, betaSlug)
	cleanedText = strings.ReplaceAll(cleanedText, "_", "")
	cleanedText = strings.ReplaceAll(cleanedText, "-", "")
	if cfg.Separator != "-" {
		errSeparator := separatorChecker(cfg.Separator)
		if errSeparator != nil {
			return "", errors.New("invalid separator: " + errSeparator.Error())
		}

		cleanedText = strings.ReplaceAll(cleanedText, "  ", cfg.Separator)
		cleanedText = strings.ReplaceAll(cleanedText, " ", cfg.Separator)
	} else {
		cleanedText = strings.ReplaceAll(cleanedText, "  ", "-")
		cleanedText = strings.ReplaceAll(cleanedText, " ", "-")
	}

	// if character > 40
	if len(cleanedText) > 40 {
		cleanedText = cleanedText[:40]
	}

	return cleanedText, nil
}
