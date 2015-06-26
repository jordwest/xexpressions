package compiler

import (
	"fmt"

	"github.com/jordwest/xexpressions/lexer"
)

type Scope struct {
	Aliases       map[string]lexer.ASTNode
	CurrentRegexp Regexp
}

func NewScope() (newScope Scope) {
	newScope.Aliases = make(map[string]lexer.ASTNode)
	newScope.CurrentRegexp = NewRegexp()
	return newScope
}

func (s Scope) DebugPrint() {
	fmt.Printf("Scope has %d aliases:\n\t%+v", len(s.Aliases), s.Aliases)
	s.CurrentRegexp.DebugPrint()
}
