package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

type Line struct {
	text        string
	indentation int
	lineNumber  int
	filename    string
}

var BlankLine *regexp.Regexp
var NormalLine *regexp.Regexp

func init() {
	BlankLine = regexp.MustCompile("^[\\s\\t]*$")
	NormalLine = regexp.MustCompile("^(\\t*)(.+)$")
}

func LinesFromText(text string, filename string) []Line {
	rawLines := strings.Split(text, "\n")
	lines := make([]Line, 0, len(rawLines))

	for i, rawLine := range rawLines {
		// Ignore blank lines
		if BlankLine.MatchString(rawLine) {
			continue
		}

		// Count indentation
		match := NormalLine.FindStringSubmatch(rawLine)
		if len(match) != 3 {
			panic("Invalid line")
		}

		lines = append(lines, Line{
			text:        match[2],
			indentation: len(match[1]),
			lineNumber:  i + 1,
			filename:    filename,
		})

	}

	return lines
}

func (l Line) String() string {
	return fmt.Sprintf("line %d of %s", l.lineNumber, l.filename)
}

type LineError struct {
	line     int
	message  string
	filename string
}

func (l Line) Error(message string, fmtParams ...interface{}) LineError {
	return LineError{
		line:     l.lineNumber,
		filename: l.filename,
		message:  fmt.Sprintf(message, fmtParams...),
	}
}

func (e LineError) Error() string {
	return fmt.Sprintf("Error in `%s` on line %d: %s", e.filename, e.line, e.message)
}
