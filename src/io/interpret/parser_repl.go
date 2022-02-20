package interpret

import (
	"bufio"
	"fmt"
	"io"
)

//////////////////////////////////////////////////////////////
// PARSER REPL COMMAND                                      //
//////////////////////////////////////////////////////////////

func Parser(in io.Reader) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)

		// Get the input
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		parsed, err := parserReturn(line)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(parsed.ToJson())
		}

	}
}
