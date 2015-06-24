package lexer

import (
	"fmt"
	"regexp"
)

type Command struct {
	Type    CommandType
	Value   string
	Comment string
}

type CommandType int

const (
	CmdXExpression CommandType = iota + 1
	CmdAliasDefinition
	CmdAliasCall
	CmdLiteral
	CmdDescription
	CmdExample
	CmdExampleMatch
	CmdExampleNonMatch
	CmdToken
	CmdGroup
	CmdSelect
	CmdCase
)

var CommandRegexp *regexp.Regexp
var LiteralRegexp *regexp.Regexp

func init() {
	CommandRegexp = regexp.MustCompile("^([A-z0-9\\s]+)(?:\\:(?:\\s(.+))?)?$")
	LiteralRegexp = regexp.MustCompile("^\\'(.+)\\'(?:\\:\\s(.+))?$")
}

// CommandFromText converts a text string to a Command
// Expects a single command only, with no surrounding whitespace
func CommandFromText(text string) (command Command, err error) {
	// First test if it's a regexp literal
	match := LiteralRegexp.FindStringSubmatch(text)
	if len(match) == 3 {
		command.Type = CmdLiteral
		command.Value = match[1]
		command.Comment = match[2]
		return command, nil
	}

	match = CommandRegexp.FindStringSubmatch(text)

	if len(match) != 3 {
		return Command{}, fmt.Errorf("Invalid command syntax: %s", text)
	}

	command.Comment = match[2]

	switch match[1] {
	case "XExpression":
		command.Type = CmdXExpression
	case "Alias":
		command.Type = CmdAliasDefinition
	case "Description":
		command.Type = CmdDescription
	case "Example":
		command.Type = CmdExample
	case "Match":
		command.Type = CmdExampleMatch
	case "Non Match":
		command.Type = CmdExampleNonMatch
	case "Select":
		command.Type = CmdSelect
	case "Case":
		command.Type = CmdCase
	default:
		command.Type = CmdAliasCall
		command.Value = match[1]
	}

	return command, nil
}
