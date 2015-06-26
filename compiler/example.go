package compiler

import (
	"regexp"

	"github.com/jordwest/xexpressions/lexer"
)

type Example struct {
	ShouldMatch bool
	Text        string
	Line        lexer.Line
}

func NewExample(match bool, text string, line lexer.Line) Example {
	return Example{
		ShouldMatch: match,
		Text:        text,
		Line:        line,
	}
}

func (e Example) Run(re string) (pass bool) {
	regexp := regexp.MustCompile(re)
	matched := regexp.MatchString(e.Text)
	if e.ShouldMatch {
		return matched
	} else {
		return !matched
	}
}
