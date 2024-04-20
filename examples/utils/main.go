package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func StartCase(s string) string {
	transformer := cases.Title(language.English)
	return transformer.String(s)
}
