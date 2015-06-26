package lexer

import "io/ioutil"

// ParseFile converts a .xexpr file into an AST
func ParseFile(filename string) (*ASTNode, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Parse(string(data), filename)
}

func Parse(text string, filename string) (*ASTNode, error) {
	rootNode := NewASTNode(nil, Command{})

	lines := LinesFromText(text, filename)
	currentIndentation := 0
	currentNode := &rootNode

	for _, line := range lines {
		var newChild *ASTNode

		newCommand, err := CommandFromText(line.text)

		if err != nil {
			return rootNode, line.Error(err.Error())
		}

		if line.indentation < 0 {
			panic("Negative indents!?")
		}

		if line.indentation-currentIndentation > 1 {
			// Cannot jump by more than one indentation
			return rootNode, line.Error("Too many indents! Check your tabs. Went from %d indents to %d indents", currentIndentation, line.indentation)
		}

		// Is this first level node?
		if line.indentation == 0 {
			newChild := rootNode.CreateChild()
			newChild.command = newCommand
			newChild.line = line
			currentNode = &newChild
			currentIndentation = line.indentation
			continue
		}

		// Is this a child of the last element?
		if line.indentation-currentIndentation == 1 {
			newChild := (*currentNode).CreateChild()
			newChild.command = newCommand
			newChild.line = line
			currentNode = &newChild
			currentIndentation = line.indentation
			continue
		}

		// Same level as previous node
		if currentIndentation == line.indentation {
			newChild = (*currentNode).parent.CreateChild()
			newChild.command = newCommand
			newChild.line = line
			currentNode = &newChild
			currentIndentation = line.indentation
			continue
		}

		// Outdent until we get to the right level
		parent := (*currentNode).parent
		for i := currentIndentation; i > line.indentation; i-- {
			parent = parent.parent
		}

		newChild = parent.CreateChild()
		newChild.command = newCommand
		newChild.line = line
		currentNode = &newChild
		currentIndentation = line.indentation
	}

	return rootNode, nil
}
