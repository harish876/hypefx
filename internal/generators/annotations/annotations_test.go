package annotations

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnnotationParserWithNoArgs(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelInfo)
	comment := "// @get"
	result, _ := ParseAnnotations(comment)
	assert.Equal(t, result.Directive, "get")
	assert.Equal(t, len(result.Params), 0)
}
func TestAnnotationParserWithNoArgsWithBrackets(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelInfo)
	comment := "//@static()"
	result, _ := ParseAnnotations(comment)
	assert.Equal(t, result.Directive, "static")
	assert.Equal(t, FromDirective(result.Directive), STATIC)
	assert.Equal(t, len(result.Params), 0)
}

func TestAnnotationParserWithArgs(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelInfo)
	comment := "// @get(key1 = val1 ,key2 = val2,key3=val3)"
	result, err := ParseAnnotations(comment)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, result.Directive, "get")
	assert.Equal(t, len(result.Params), 3)
	assert.Equal(t, result.Params["key1"], "val1")
	assert.Equal(t, result.Params["key2"], "val2")
	assert.Equal(t, result.Params["key3"], "val3")
}
