package ast

type Node interface {
	Node() string
}

type Expression interface {
	Node
	expressionNode()
}
