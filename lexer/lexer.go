package lexer

import "fmt"

func ParseFile(filename string) {
	panic("Not implemented")
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
			fmt.Printf("[%d] Creating child of root node\n", line.lineNumber)
			newChild := rootNode.CreateChild()
			newChild.command = newCommand
			newChild.line = line
			currentNode = &newChild
			currentIndentation = line.indentation
			continue
		}

		// Is this a child of the last element?
		if line.indentation-currentIndentation == 1 {
			fmt.Printf("[%d] Creating child of previous node\n", line.lineNumber)
			newChild := (*currentNode).CreateChild()
			newChild.command = newCommand
			newChild.line = line
			currentNode = &newChild
			currentIndentation = line.indentation
			continue
		}

		// Same level as previous node
		if currentIndentation == line.indentation {
			fmt.Printf("[%d] Creating child of previous node's parent\n", line.lineNumber)
			newChild = (*currentNode).parent.CreateChild()
			newChild.command = newCommand
			newChild.line = line
			currentNode = &newChild
			currentIndentation = line.indentation
			continue
		}

		// Outdent until we get to the right level
		fmt.Printf("[%d] Outdenting... ", line.lineNumber)
		parent := (*currentNode).parent
		for i := currentIndentation; i > line.indentation; i-- {
			fmt.Printf("%d... ", i)
			parent = parent.parent
		}
		fmt.Printf("Done\n")

		newChild = parent.CreateChild()
		newChild.command = newCommand
		newChild.line = line
		currentNode = &newChild
		currentIndentation = line.indentation
	}

	return rootNode, nil
}
