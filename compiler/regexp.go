package compiler

import "fmt"

type CaptureGroup struct {
	Name         string
	VariableName string
	Description  string
	Index        int // The capture group's index when the regular expression is executed
}

type Regexp struct {
	TextName      string         // The textual name of the regular expression
	VariableName  string         // The name of the regular expression in variable naming format
	Description   string         // A description of the regular expression, if any
	RegexpText    string         // The compiled text for the regular expression
	CaptureGroups []CaptureGroup // A list of the named capture groups in the regular expression
	Examples      []Example      // A list of the examples for this regexp
	Source        string         // The location of the original definition

	groupIndex int
}

func NewRegexp() Regexp {
	return Regexp{
		CaptureGroups: make([]CaptureGroup, 0),
		Examples:      make([]Example, 0),
	}
}

func (re *Regexp) AddCaptureGroup(name string, description string) {
	re.groupIndex++
	re.CaptureGroups = append(re.CaptureGroups, CaptureGroup{
		Name:        name,
		Description: description,
		Index:       re.groupIndex,
	})
}

func (re *Regexp) IncrementCaptureGroup(increment int) {
	re.groupIndex += increment
}

func (re Regexp) DebugPrint() {
	fmt.Printf("--- %s ---\nDescription: %s\n", re.TextName, re.Description)

	fmt.Printf("Examples:\n")
	for _, example := range re.Examples {
		fmt.Printf("\t%t\t%s\n", example.ShouldMatch, example.Text)
	}

	fmt.Printf("Capture Groups:\n")
	for _, group := range re.CaptureGroups {
		fmt.Printf("\t%+v\n", group)
	}

	fmt.Printf("Compiled: %s\n", re.RegexpText)
}
