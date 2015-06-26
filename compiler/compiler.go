package compiler

import (
	"fmt"
	"strings"

	"github.com/jordwest/xexpressions/lexer"
)

func CompileRoot(node lexer.ASTNode) (regexps []Regexp, scope Scope, err error) {
	if !node.IsCommandType(0) {
		panic("Expecting root node")
	}

	regexps = make([]Regexp, 0)

	scope = NewScope()
	for _, child := range node.Children() {

		command := child.Command()
		switch command.Type {
		case lexer.CmdAliasDefinition:
			if scope, err = DefineAlias(*child, scope); err != nil {
				return regexps, scope, err
			}
		case lexer.CmdXExpression:
			regexp, scope, err := CompileExpression(*child, scope)
			if err != nil {
				return regexps, scope, err
			}
			regexps = append(regexps, regexp)
		}
	}

	return regexps, scope, err
}

func DefineAlias(node lexer.ASTNode, scope Scope) (parentScope Scope, err error) {
	if !node.IsCommandType(lexer.CmdAliasDefinition) {
		panic("Expecting alias definition command")
	}

	scope.Aliases[node.Command().Comment] = node

	return scope, nil
}

func CompileExpression(node lexer.ASTNode, scope Scope) (regexp Regexp, parentScope Scope, err error) {
	output := ""

	if !node.IsCommandType(lexer.CmdXExpression) {
		panic("Compile should be passed an XExpression node only")
	}

	// List of examples should be isolated in this expression's scope
	scope.CurrentRegexp = NewRegexp()
	output, scope, err = compileNodeChildren(node, scope)

	// Run tests on the examples
	for _, example := range scope.CurrentRegexp.Examples {
		pass := example.Run(output)
		if pass {
			fmt.Printf(" ✔ Match test passed on %s\n", example.Line.String())
		} else {
			fmt.Printf(" ✕ Match test failed for\n\t%s\non %s\n", example.Text, example.Line.String())
		}
	}

	scope.CurrentRegexp.RegexpText = output

	scope.CurrentRegexp.DebugPrint()

	return scope.CurrentRegexp, scope, err
}

func compileGroup(node lexer.ASTNode, scope Scope) (output string, parentScope Scope, err error) {
	if !node.IsCommandType(lexer.CmdGroup) {
		panic("Expected Group node")
	}

	childOutput, scope, err := compileNodeChildren(node, scope)
	if err != nil {
		return "", scope, err
	}

	preventCapture := "?:"
	if node.Command().Params == "Capture" {
		preventCapture = ""
	}

	output = fmt.Sprintf("(%s%s)", preventCapture, childOutput)

	return output, scope, err
}

func compileNodeChildren(node lexer.ASTNode, scope Scope) (output string, parentScope Scope, err error) {
	for _, child := range node.Children() {
		childOutput := ""

		switch child.Command().Type {
		case lexer.CmdLiteral:
			childOutput = child.Command().Value
		case lexer.CmdAliasCall:
			if childOutput, err = callAlias((*child).Command().Value, scope); err != nil {
				return "", scope, child.Line().Error(err.Error())
			}
		case lexer.CmdGroup:
			if childOutput, scope, err = compileGroup(*child, scope); err != nil {
				return "", scope, err
			}
		case lexer.CmdSelect:
			if childOutput, scope, err = compileSelect(*child, scope); err != nil {
				return "", scope, err
			}
		case lexer.CmdExample:
			if childOutput, scope, err = compileExamples(*child, scope); err != nil {
				return "", scope, err
			}
		default:
			if childOutput, scope, err = compileNodeChildren(*child, scope); err != nil {
				return "", scope, err
			}
		}

		output += childOutput
	}

	return output, scope, nil
}

func compileSelect(node lexer.ASTNode, scope Scope) (output string, parentScope Scope, err error) {
	cases := make([]string, len(node.Children()))

	for i, child := range node.Children() {
		if !child.IsCommandType(lexer.CmdCase) {
			return "", scope, child.Line().Error("Only 'Case' statements are valid inside a 'Select' statement")
		}

		if cases[i], scope, err = compileNodeChildren(*child, scope); err != nil {
			return "", scope, err
		}
	}

	// Join the cases with pipes and put it all into a non-capture group
	output = fmt.Sprintf("(?:%s)", strings.Join(cases, "|"))
	return output, scope, nil
}

func compileExamples(node lexer.ASTNode, scope Scope) (output string, parentScope Scope, err error) {
	if !node.IsCommandType(lexer.CmdExample) {
		panic("Expecting example command")
	}

	for _, child := range node.Children() {
		command := child.Command()
		switch command.Type {
		case lexer.CmdExampleMatch:
			scope.CurrentRegexp.Examples = append(scope.CurrentRegexp.Examples, NewExample(true, command.Comment, child.Line()))
		case lexer.CmdExampleNonMatch:
			scope.CurrentRegexp.Examples = append(scope.CurrentRegexp.Examples, NewExample(false, command.Comment, child.Line()))
		}
	}

	return "", scope, nil
}

func callAlias(aliasName string, scope Scope) (output string, err error) {
	node, ok := scope.Aliases[aliasName]
	if !ok {
		return "", fmt.Errorf("Cannot find alias with name %s", aliasName)
	}

	// The scope should be private to the alias call - nothing inside the alias
	// can affect the parent scope
	output, _, err = compileNodeChildren(node, scope)
	return output, err
}
