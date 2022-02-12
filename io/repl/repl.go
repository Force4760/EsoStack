package repl

import (
	"github.com/Force4760/pipes/io/colorize"
)

var PROMPT = colorize.Colorize("-> ", colorize.GREEN)
var MSG = colorize.Colorize("Hello! This is the Pipes programming language!", colorize.BLUE)

func Repl() {
}
