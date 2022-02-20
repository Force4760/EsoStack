package funcs

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/force4760/esostack/io/colorize"
	"github.com/force4760/esostack/io/interpret"
	"github.com/force4760/esostack/src/stack"
)

func PlayFunc(in, out binding.String) func() {
	return func() {
		s, err := in.Get()
		if err != nil {
			s = ""
		}

		stk := stack.NewStack([]float64{})

		err = interpret.Interpret(s, stk)

		if err != nil {
			out.Set(colorize.UndoError(err))
		} else {
			out.Set(stk.StringNoColor())
		}
	}
}
