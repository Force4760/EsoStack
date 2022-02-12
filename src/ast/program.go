package ast

// Datastructure conteining the operations of an individual file
type FunctionFile []Expression

func (f *FunctionFile) String() string {
	return toStr(*f)
}

// Datastructure that holds all different files
type Program struct {
	Files map[string]Expression
}
