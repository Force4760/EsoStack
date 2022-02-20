package funcs

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/force4760/esostack/src/interpreter/stack"
	"github.com/force4760/esostack/src/io/colorize"
	"github.com/force4760/esostack/src/io/interpret"
)

// Process Play/Run functionality
func PlayFunc(in, out binding.String) func() {
	return func() {
		// Get the content of the input
		s, err := in.Get()
		if err != nil {
			s = ""
		}

		// Create an empty stack
		stk := stack.NewStack([]float64{})

		// Interpret the input
		err = interpret.Interpret(s, stk)

		if err != nil {
			// Show Error
			out.Set(colorize.UndoError(err))
		} else {
			// Show the Stack
			out.Set(stk.StringNoColor())
		}
	}
}
