package lexer

import "testing"

func TestLexerParse(t *testing.T) {
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

	ast, err := Parse(textSample, "test")
	if err != nil {
		t.Errorf("Error parsing:\n\t%s\n", err)
	}
	ast.DebugPrint("")
}
