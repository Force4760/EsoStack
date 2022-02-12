package ast

import "encoding/json"

// Datastructure conteining the operations of an individual file
type FunctionFile []Expression

func (f *FunctionFile) String() string {
	return toStr(*f)
}

func (f *FunctionFile) ToJson() string {
	js, err := json.MarshalIndent(f, "", "  ")

	if err != nil {
		return ""
	}

	return string(js)
}

// Datastructure that holds all different files
type Program struct {
	Files map[string]Expression
}
