package interpret

import (
	"fmt"
	"io/ioutil"

	"github.com/force4760/esostack/src/interpreter/stack"
)

//////////////////////////////////////////////////////////////
// FILE INTERPRET COMMAND                                   //
//////////////////////////////////////////////////////////////

func File(path string) {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert []byte to string and print to screen
	line := string(content)

	// Create the stack
	stk := stack.NewStack([]float64{})

	// Interpret the file contents
	err = Interpret(line, stk)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stk)
}
