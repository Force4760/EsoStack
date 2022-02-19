package initProject

import (
	"fmt"

	"github.com/force4760/esostack/io/colorize"
)

// Init a new Pipes Project
func Init(num int) {
	// The min number of files is 1
	// The max number of files is 26
	if num < 1 || num > 26 {
		fmt.Println(
			colorize.Error("The number of files must be between 1 and 26."),
		)
		return
	}

	// get information
	name, project, desc := get_input()

	// Try to create the project folder
	err := makeDir(project)
	if err != nil {
		// if there was a problem
		fmt.Println(
			colorize.Error(err.Error()),
		)
		return
	}

	// Try to create the src folder
	err = makeDir(project + "/src")
	if err != nil {
		// if there was a problem
		fmt.Println(
			colorize.Error(err.Error()),
		)
		return
	}

	srcFiles(project, num)

	mainFile(project)

	readmeFile(name, project, desc)

	chrsFile(project)

	fmt.Println(
		colorize.Colorize("Project Created", colorize.GREEN),
	)
}
