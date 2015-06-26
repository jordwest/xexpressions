package writer

import (
	"os"
	"testing"

	"github.com/jordwest/xexpressions/compiler"
	"github.com/jordwest/xexpressions/lexer"
)

func TestWriteJavascript(t *testing.T) {
	ast, err := lexer.ParseFile("../compiler/builtin.xexpr")
	if err != nil {
		t.Errorf("Error parsing:\n\t%s\n", err)
	}

	demos, err := lexer.ParseFile("../demos.xexp")
	if err != nil {
		t.Errorf("Error parsing:\n\t%s\n", err)
	}

	ast.Include(demos)

	output, _, err := compiler.CompileRoot(*ast)
	if err != nil {
		t.Errorf("Error parsing:\n\t%s\n", err)
	}

	err = WriteRegexps(output, "../templates/javascript.js", os.Stdout)

}
