package annotations

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"

	sitter "github.com/harish876/go-tree-sitter"
	"github.com/harish876/go-tree-sitter/golang"
	"github.com/harish876/hypefx/internal/utils"
)

type HttpMethod int

const (
	GET HttpMethod = iota
	POST
	PUT
	DELETE
	INVALID_METHOD
)

type QueryExecutionParams struct {
	Cursor     *sitter.QueryCursor
	Query      *sitter.Query
	Node       *sitter.Node
	SourceCode []byte
}

func NewQueryExecutionParams(cursor *sitter.QueryCursor, query *sitter.Query, node *sitter.Node, sourceCode []byte) *QueryExecutionParams {
	return &QueryExecutionParams{
		Cursor:     cursor,
		Query:      query,
		Node:       node,
		SourceCode: sourceCode,
	}
}

func GetQueryCursor(lang *sitter.Language, sourceCode []byte, query []byte) (*QueryExecutionParams, error) {
	node, _ := sitter.ParseCtx(context.Background(), sourceCode, lang)

	sitterQuery, _ := sitter.NewQuery(query, lang)
	queryCursor := sitter.NewQueryCursor()

	return NewQueryExecutionParams(queryCursor, sitterQuery, node, sourceCode), nil
}

type AnnotatedRouteDetails struct {
	Method HttpMethod
}

func FromAnnotaion(route string) (*AnnotatedRouteDetails, error) {
	//do the string pre-processing here
	pattern := `^//\s*@(\w+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(route)
	if len(matches) < 2 {
		return nil, fmt.Errorf("invalid annotation format\n")
	}

	method := FromMethod(matches[1])
	if method == INVALID_METHOD {
		return nil, fmt.Errorf("invalid http method %s\n", matches[2])
	}
	return &AnnotatedRouteDetails{
		Method: method,
	}, nil
}

func FromMethod(method string) HttpMethod {
	formattedMethod := utils.UpperCase(method)
	switch formattedMethod {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	default:
		return INVALID_METHOD
	}
}

func FromEnum(methodEnum HttpMethod) string {
	switch methodEnum {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	default:
		return "INVALID_METHOD"
	}
}

type HandlerDetails struct {
	HandlerName string
	Method      HttpMethod
}

func FromHandlerDetails(handlerName string, method HttpMethod) HandlerDetails {
	return HandlerDetails{
		HandlerName: handlerName,
		Method:      method,
	}
}

func GetHandlerDetailsFromAnnotations(sourceCode []byte) ([]HandlerDetails, string, error) {
	lang := golang.GetLanguage()
	query := []byte(`
	(
		(package_clause
			(package_identifier) @packageName
		)
		(comment) @routeDetails
		.
		(function_declaration
			name: (identifier) @handlerName
		)
	  )
	`)
	q, err := GetQueryCursor(lang, sourceCode, query)
	if q.Node.HasError() {
		return nil, "", fmt.Errorf("Syntax Tree has errors")
	}

	if err != nil {
		return nil, "", err
	}
	q.Cursor.Exec(q.Query, q.Node)
	handlers := make([]HandlerDetails, 0)
	var packageName string
	for {
		m, ok := q.Cursor.NextMatch()
		if !ok {
			break
		}
		m = q.Cursor.FilterPredicates(m, q.SourceCode)
		var annotatedRoute string
		for _, c := range m.Captures {
			if c.Node.Type() == "package_identifier" && packageName == "" {
				packageName = c.Node.Content(sourceCode)
			} else if c.Node.Type() == "comment" {
				annotatedRoute = c.Node.Content(sourceCode)
			} else if c.Node.Type() == "identifier" {
				handlerName := c.Node.Content(sourceCode)
				if annotatedRoute == "" {
					slog.Debug("GetHandlerDetailsFromAnnotations", "route details not set for", handlerName)
					continue
				} else {
					routeDetails, err := FromAnnotaion(annotatedRoute)
					if err != nil {
						slog.Debug("GetHandlerDetailsFromAnnotations", "bad annotated route Route Details", routeDetails, "Comment", annotatedRoute)
						continue
					}
					slog.Debug("GetHandlerDetailsFromAnnotations", "Annotated route Handler", handlerName, "Comment", annotatedRoute)

					handlers = append(handlers, FromHandlerDetails(
						handlerName,
						routeDetails.Method,
					))
				}
				annotatedRoute = ""
			}
		}
	}
	return handlers, packageName, nil
}
