package interpret

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/force4760/esostack/src/interpreter/stack"
)

//////////////////////////////////////////////////////////////
// FILE INTERPRET COMMAND                                   //
//////////////////////////////////////////////////////////////

func RunMain(path string) {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert []byte to string and print to screen
	funcs, err := getFuncs(content)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Create the stack
	stk := stack.NewStack([]float64{})

	for _, i := range funcs {
		// Ignore blank lines
		if i == "" {
			continue
		}

		// Interpret the file contents
		err = RunFile(formatPath(i), stk)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(stk)
}

func getFuncs(content []byte) ([]string, error) {
	line := string(content)
	line = strings.ReplaceAll(line, " ", "")

	funcs := strings.Split(line, "\n")

	for _, i := range funcs {
		if i == "" {
			continue
		}

		if !exists(i) {
			return []string{}, NotExistsError(i)
		}
	}

	return funcs, nil
}

func exists(path string) bool {
	_, err := os.Stat(
		formatPath(path),
	)
	return !errors.Is(err, os.ErrNotExist)
}

func formatPath(str string) string {
	return fmt.Sprintf("./funcs/%s.stk", str)
}
