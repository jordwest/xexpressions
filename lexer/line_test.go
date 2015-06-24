package lexer

import (
	"fmt"
	"testing"
)

func TestLineSplit(t *testing.T) {
	textSample := `
XExpression: Date
	Description: This matches a valid date between year
	Example:
		Match: 12345
		Non Match: 78
	'20[0-9][0-9]'
		`

	fmt.Printf("Lines:\n %+v\n", LinesFromText(textSample))
}
