package compiler

import (
	"fmt"
	"testing"

	"github.com/jordwest/xexpressions/lexer"
)

func TestCompile(t *testing.T) {
	textSample := `
Alias: Date Separator
	Description: A date separator
	'[- /.]'

XExpression: Date
	Description: This matches a valid date between year 1900 and 2099
	Example:
		Match: 2015-02-21
		Match: 2015/02/21
		Match: 2015 02 21
		Non Match: 3056-02-16
		Non Match: 2015-2-8
		Non Match: 2015-05-34
		Non Match: 2015-13-05

	Group[Capture]: Year
		'(19|20)': Must be in range 1900-2099
		Digit
		Digit
	Date Separator
	Group[Capture]: Month
		Select:
			Case: Jan to Sept
				'0[1-9]'
			Case: Oct to Dec
				'1[012]'
	Date Separator
	Group[Capture]: Day
		Select:
			Case: 1st to 9th
				'0[1-9]'
			Case: 10th to 29th
				'[12][0-9]'
			Case: 30th to 31st
				'3[01]'
		`

	global := `
Alias: Digit
	'\d'
	`

	globalAst, err := lexer.Parse(global, "compiler_test.go > global")

	_, err = lexer.Parse(textSample, "compiler_test.go > textSample")
	ast, err := lexer.ParseFile("../demos.xexp")
	if err != nil {
		t.Errorf("Error parsing:\n\t%s\n", err)
	}

	globalAst.Append(ast)

	//output, scope, err := CompileExpression(*ast.Children()[1], Scope{})
	output, scope, err := CompileRoot(*globalAst)
	if err != nil {
		t.Errorf("Error parsing:\n\t%s\n", err)
	}
	fmt.Printf("Output:\n%s\n", output)
	scope.DebugPrint()
}
