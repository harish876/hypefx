package utils

import "testing"

func TestReplaceFileContent(t *testing.T) {
	filePath := "example.go"
	ReplaceFileContent(filePath, "foobar", "baz")
}
