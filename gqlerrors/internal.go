package gqlerrors

import (
	"github.com/graphql-go/graphql/language/ast"
)

type InternalError struct {
	Original error
	Nodes    []ast.Node
}
