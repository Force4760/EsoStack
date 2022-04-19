package initProject

import (
	"fmt"

	"github.com/force4760/esostack/src/io/colorize"
)

// Init a new EsoStack Project
func Init() {
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
	err = makeDir(project + "/funcs")
	if err != nil {
		// if there was a problem
		fmt.Println(
			colorize.Error(err.Error()),
		)
		return
	}

	readmeFile(name, project, desc)

	chrsFile(project)

	funcLogFile(project)

	mainFile(project)

	fmt.Println(
		colorize.Colorize("Project Created", colorize.GREEN),
	)
}
