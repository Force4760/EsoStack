package ast

// Datastructure conteining the operations of an individual file
type FunctionFile []Expression

// Datastructure that holds all different files
type Program struct {
	Files map[string]Expression
}
