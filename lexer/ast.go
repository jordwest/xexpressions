package lexer

import "fmt"

type ASTNode struct {
	parent   *ASTNode
	children []*ASTNode
	command  Command
	line     Line // Keep track of the original line for later error reporting
	order    int
}

var num = 1

func NewASTNode(parent *ASTNode, command Command) *ASTNode {
	newNode := &ASTNode{}
	newNode.parent = parent
	newNode.children = make([]*ASTNode, 0)
	newNode.command = command
	newNode.order = num
	newNode.line = Line{}
	num++
	return newNode
}

func (n *ASTNode) CreateChild() *ASTNode {
	newChild := NewASTNode(n, Command{})
	n.children = append(n.children, newChild)
	return newChild
}

func (n *ASTNode) Append(node *ASTNode) {
	n.children = append(n.children, node.children...)
}

func (n *ASTNode) Command() Command     { return n.command }
func (n *ASTNode) Line() Line           { return n.line }
func (n *ASTNode) Children() []*ASTNode { return n.children }
func (n *ASTNode) Parent() *ASTNode     { return n.parent }

func (n *ASTNode) DebugPrint(prependLine string) {
	if n.parent == nil {
		fmt.Printf("%sROOT NODE -- %d children\n", prependLine, len(n.children))
	} else {
		fmt.Printf("%sCMD %d -- Order: %d, Value: %s, Params: %s, Comment: %s, %d children\n", prependLine, n.command.Type, n.order, n.command.Value, n.command.Params, n.command.Comment, len(n.children))
	}

	for _, child := range n.children {
		child.DebugPrint(prependLine + "\t")
	}
}

// IsCommandType confirms whether this node is a command of a certain type
func (n *ASTNode) IsCommandType(commandType CommandType) bool {
	return (n.command.Type == commandType)
}
