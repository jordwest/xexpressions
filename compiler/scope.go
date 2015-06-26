package compiler

import (
	"fmt"

	"github.com/jordwest/xexpressions/lexer"
)

type Scope struct {
	Aliases  map[string]lexer.ASTNode
	Examples []Example
}

func NewScope() (newScope Scope) {
	newScope.Aliases = make(map[string]lexer.ASTNode)
	newScope.Examples = make([]Example, 0)
	return newScope
}

func (s Scope) DebugPrint() {
	fmt.Printf("Scope has %d aliases:\n\t%+v", len(s.Aliases), s.Aliases)
	fmt.Printf("Scope has %d examples:\n\t%+v", len(s.Examples), s.Examples)
}
