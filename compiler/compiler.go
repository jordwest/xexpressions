package compiler

import (
	"fmt"
	"strings"

	"github.com/jordwest/xexpressions/lexer"
)

func CompileRoot(node lexer.ASTNode) (output string, err error) {
	if !node.IsCommandType(0) {
		panic("Expecting root node")
	}

	scope := NewScope()
	output, err = compileNodeChildren(node, scope)

	return output, err
}

func CompileExpression(node lexer.ASTNode, scope Scope) (output string, err error) {
	if !node.IsCommandType(lexer.CmdXExpression) {
		panic("Compile should be passed an XExpression node only")
	}

	return compileNodeChildren(node, scope)
}

func compileGroup(node lexer.ASTNode, scope Scope) (output string, err error) {
	if !node.IsCommandType(lexer.CmdGroup) {
		panic("Expected Group node")
	}

	childOutput, err := compileNodeChildren(node, scope)
	if err != nil {
		return "", err
	}

	preventCapture := "?:"
	if node.Command().Params == "Capture" {
		preventCapture = ""
	}

	output = fmt.Sprintf("(%s%s)", preventCapture, childOutput)

	return output, err
}

func compileNodeChildren(node lexer.ASTNode, scope Scope) (output string, err error) {
	for _, child := range node.Children() {
		childOutput := ""

		switch child.Command().Type {
		case lexer.CmdLiteral:
			childOutput = child.Command().Value
		case lexer.CmdAliasCall:
			childOutput = child.Command().Value
		case lexer.CmdGroup:
			if childOutput, err = compileGroup(*child, scope); err != nil {
				return "", err
			}
		case lexer.CmdSelect:
			if childOutput, err = compileSelect(*child, scope); err != nil {
				return "", err
			}
		default:
			if childOutput, err = compileNodeChildren(*child, scope); err != nil {
				return "", err
			}
		}

		output += childOutput
	}

	return output, nil
}

func compileSelect(node lexer.ASTNode, scope Scope) (output string, err error) {
	cases := make([]string, len(node.Children()))

	for i, child := range node.Children() {
		if !child.IsCommandType(lexer.CmdCase) {
			return "", child.Line().Error("Only 'Case' statements are valid inside a 'Select' statement")
		}

		if cases[i], err = compileNodeChildren(*child, scope); err != nil {
			return "", err
		}
	}

	// Join the cases with pipes and put it all into a non-capture group
	output = fmt.Sprintf("(?:%s)", strings.Join(cases, "|"))
	return output, nil
}
