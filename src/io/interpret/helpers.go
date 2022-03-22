package interpret

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/force4760/esostack/src/interpreter/stack"
	"github.com/force4760/esostack/src/io/colorize"
)

//////////////////////////////////////////////////////////////
// REPL ADITIONAL COMMANDS                                  //
//////////////////////////////////////////////////////////////

// Prompt for the repl
var PROMPT = colorize.Colorize("-> ", colorize.GREEN)

// Intro message for the repl
var MSG = colorize.Colorize("Hello! This is the Pipes programming language!", colorize.BLUE)

// Error to show when no file path was given
var ErrNoPath = colorize.Error("No path was provided!")

// Repl changing functions
// clear -> clear the repl
// empty -> empty the stack
// stack -> print the stack
// load
func Commands(line string, stk *stack.Stack) bool {
	cmds := strings.Split(line, " ")

	switch cmds[0] {
	case "clear", "cls", ":c":
		// Clear the repl
		ClearScreen()
		return true

	case "empty", "ept", ":e":
		// Empty the stack
		stk.Empty()
		return true

	case "stack", "stk", ":s":
		// Print the stack
		fmt.Println(stk)
		return true

	case "load", "ld", ":l":
		// Load and evaluate a File
		loadFile(cmds, stk)
		return true

	case "help", ":h":
		// Print the available commands
		fmt.Println(`
  stack (stk) (:s) -> printe the current state of the stack
  empty (ept) (:e) -> empty the stack
  clear (cls) (:c) -> clear the terminal screen
  load  (ld)  (:l) -> load and evaluate a file in the current directory
  help        (:h) -> show this message`)
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

func loadFile(cmds []string, stk *stack.Stack) {
	if len(cmds) != 2 {
		fmt.Println(ErrNoPath)
		return
	}

	path := fmt.Sprintf("./funcs/%s.es", cmds[1])

	err := RunFile(path, stk)
	if err != nil {
		fmt.Println(
			colorize.Error(err.Error()),
		)
	}
}
