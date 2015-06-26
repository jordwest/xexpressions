package compiler

import "github.com/jordwest/xexpressions/lexer"

type Scope struct {
	Aliases map[string]lexer.ASTNode
}

func NewScope() (newScope Scope) {
	newScope.Aliases = make(map[string]lexer.ASTNode)
	return newScope
}
