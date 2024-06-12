package cases

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleCase(s string) string {
	transformer := cases.Title(language.English)
	return transformer.String(s)
}
