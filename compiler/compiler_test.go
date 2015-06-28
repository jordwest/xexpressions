package compiler

import (
	"fmt"
	"testing"

	"github.com/jordwest/xexpressions/lexer"
)

func TestCompile(t *testing.T) {
	ast, err := lexer.ParseFile("../example/common.xexpr")
	demos, err := lexer.ParseFile("../example/demos.xexpr")
	if err != nil {
		t.Errorf("Error parsing:\n\t%s\n", err)
	}

	ast.Include(demos)

	//output, scope, err := CompileExpression(*ast.Children()[1], Scope{})
	output, _, err := CompileRoot(*ast)
	if err != nil {
		t.Errorf("Error parsing:\n\t%s\n", err)
	}
	fmt.Printf("Output:\n%s\n", output)
}
