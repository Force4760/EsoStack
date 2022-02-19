package interpret

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/force4760/esostack/io/colorize"
	"github.com/force4760/esostack/src/stack"
)

//////////////////////////////////////////////////////////////
// REPL ADITIONAL COMMANDS                                  //
//////////////////////////////////////////////////////////////

// Prompt for the repl
var PROMPT = colorize.Colorize("-> ", colorize.GREEN)

// Intro message for the repl
var MSG = colorize.Colorize("Hello! This is the Pipes programming language!", colorize.BLUE)

// Repl changing functions
// clear -> clear the repl
// empty -> empty the stack
// stack -> print the stack
func Commands(line string, stk *stack.Stack) bool {
	switch line {
	case "clear", "cls":
		// Clear the repl
		ClearScreen()
		return true

	case "empty", "ept":
		// Empty the stack
		stk.Empty()
		return true

	case "stack", "stk":
		// Print the stack
		fmt.Println(stk)
		return true

	case "help":
		// Print the available commands
		fmt.Println(`
	stack (stk) -> printe the current state of the stack
	empty (ept) -> empty the stack
	clear (cls) -> clear the terminal screen
	help        -> show this message`)
		return true

	}

	return false
}

//////////////////////////////////////////////////////////////
// CLEAR SCREEN COMMAND                                     //
//////////////////////////////////////////////////////////////

// Clear Command Map
// OS - clear cmd
var clearScreen = map[string]func(){
	"darwin": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},

	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},

	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

// Clears the terminal window
func ClearScreen() {
	fun, valid := clearScreen[runtime.GOOS]
	if valid {
		fun()
	}
}
