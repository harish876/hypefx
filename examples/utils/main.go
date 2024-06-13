package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	startCaseTransformer = cases.Title(language.English)
	upperCaseTransformer = cases.Upper(language.English)
)

func StartCase(s string) string {
	return startCaseTransformer.String(s)
}

func UpperCase(s string) string {
	return upperCaseTransformer.String(s)
}
