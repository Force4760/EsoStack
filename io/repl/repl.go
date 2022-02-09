package repl

import (
	"fmt"
	"os"

	"github.com/Force4760/pipes/io/colorize"
)

var PROMPT = colorize.Colorize("-> ", colorize.GREEN)
var MSG = colorize.Colorize("Hello! This is the Pipes programming language!", colorize.BLUE)

func Repl() {
	fmt.Println(MSG)

	Lexer(os.Stdin)
}
