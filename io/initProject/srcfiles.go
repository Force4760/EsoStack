package initProject

import (
	"fmt"
	"io/ioutil"
)

// Create the src files
func srcFiles(folder string, n int) {
	// The min number of files is 1
	// The max number of files is 26
	if n < 1 || n > 26 {
		fmt.Println("Error: The number of files must be between 1 and 26")
	}

	for i := 0; i < n; i++ {
		// name of the file
		// n=0 -> A    n=1 -> B    ...    n=26 -> Z
		name := string(rune(65 + i))

		makeFile(folder, name)
	}
}

// Create a single function file
func makeFile(folder, name string) error {
	// Path for the file
	path := folder + "/src/" + name + ".pipes"

	// Comment to be written on the text
	text := []byte("( " + path + " )\n( {File Description} )")

	// Write the file
	return ioutil.WriteFile(path, text, 0644)
}
