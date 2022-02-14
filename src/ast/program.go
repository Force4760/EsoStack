package ast

import (
	"encoding/json"

	"github.com/Force4760/pipes/src/stack"
)

//////////////////////////////////////////////////////////////
// PROGRAM                                                  //
//////////////////////////////////////////////////////////////

// Datastructure that holds all different files
type Program struct {
	Files map[string]FunctionFile
}

//////////////////////////////////////////////////////////////
// FunctionFile                                             //
//////////////////////////////////////////////////////////////

// Datastructure conteining the operations of an individual file
type FunctionFile []Expression

// Convert FunctionFile to string
func (f *FunctionFile) String() string {
	return toStr(*f)
}

// Convert FunctionFile to json
func (f *FunctionFile) ToJson() string {
	js, err := json.MarshalIndent(f, "", "  ")

	if err != nil {
		return ""
	}

	return string(js)
}

// Evaluate every Node in an FunctionFile
func (f FunctionFile) Eval(s *stack.Stack) error {
	for _, i := range f {
		err := i.Eval(s)

		if err != nil {
			return err
		}
	}

	return nil
}
