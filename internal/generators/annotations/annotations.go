package annotations

import (
	"fmt"
	"strings"

	"log/slog"
)

/*
	Format could be like this
	@get,@put
	@static(key1=val1,key2=val2)
	@ put()
*/

type AnnotationParser struct {
	Comment string
}

type Annotation struct {
	Directive string            // this would be the @[command]
	Params    map[string]string //this would be the structured key value pairs in parantheses
}

func NewAnnotation() Annotation {
	return Annotation{
		Params: make(map[string]string),
	}
}

func (r *AnnotationParser) Parse(annotation *Annotation) (int, error) {
	currIdx, err := r.ParseDirective(annotation)
	if err != nil {
		return currIdx, err
	}
	return r.ParseArgs(currIdx, annotation)
}

func (r *AnnotationParser) ParseDirective(annotation *Annotation) (int, error) {
	idx := 0
	token := r.Comment

	if token[:1] != "@" {
		fmt.Println("error here")
		return 1, fmt.Errorf("invalid directive name. should always start with @a: %s", token[:1])
	}

	idx += 1 //advance to next byte
	for {
		if token[idx] == ' ' {
			idx++
		} else {
			break
		}
	}

	start := idx
	for {
		if idx == len(token) || token[idx] == '(' {
			break
		}
		idx++
	}
	directive := token[start:idx]
	slog.Debug("Parse Direcive", "Directive", directive)
	slog.Debug("Parse Directive", "Args", token[idx:])
	annotation.Directive = directive
	return idx, nil
}

func (r *AnnotationParser) ParseArgs(idx int, annotation *Annotation) (int, error) {
	token := r.Comment

	if idx >= len(token) {
		return idx, nil
	}

	if token[idx] != '(' {
		return idx + 2, fmt.Errorf("invalid option name: %b", token[idx])
	}
	idx++ //move to the next byte

	for {
		if idx >= len(token) {
			break
		}

		for {
			if token[idx] != ' ' {
				break
			}
			idx++
		}

		var key string
		for {
			if token[idx] == '=' {
				break
			} else if token[idx] == ')' {
				return idx, nil
			} else if token[idx] != ' ' {
				key += string(token[idx])
			}
			idx++
		}

		slog.Debug("ParseArgs", "Key", key)
		idx++ //move ahead of the =

		for {
			if token[idx] != ' ' {
				break
			}
			idx++
		}

		var value string
		for {
			if idx >= len(token) || token[idx] == ')' || token[idx] == ',' {
				break
			} else if token[idx] != ' ' {
				value += string(token[idx])
			}
			idx++
		}
		slog.Debug("ParseArgs", "Value", value)
		annotation.Params[key] = value
		idx++ //move ahead of the ) or the ,

	}
	return len(token), nil
}

func ParseAnnotations(comment string) (Annotation, error) {
	var token string
	if strings.HasPrefix(comment, "//") {
		token = strings.TrimPrefix(comment, "//")
		idx := 0
		for {
			if token[idx] != ' ' {
				break
			}
			idx++
		}
		token = token[idx:]

	} else {
		return Annotation{}, fmt.Errorf("invalid comment syntax. should start with '//' or '// ', got %v", token)
	}
	p := AnnotationParser{Comment: token}
	annotation := NewAnnotation()
	if _, err := p.Parse(&annotation); err != nil {
		fmt.Println(err)
	}
	return annotation, nil
}
