package initProject

import (
	"fmt"

	"github.com/Force4760/pipes/io/colorize"
)

// Init a new Pipes Project
func Init(num int) {
	// get information
	name, project, desc := get_input()

	// Try to create the project folder
	err := makeDir(project)
	if err != nil {
		// if there was a problem
		fmt.Println(
			colorize.Colorize(err.Error(), colorize.RED),
		)
		return
	}

	// Try to create the src folder
	err = makeDir(project + "/src")
	if err != nil {
		// if there was a problem
		fmt.Println(
			colorize.Colorize(err.Error(), colorize.RED),
		)
		return
	}

	srcFiles(project, num)

	mainFile(project)

	readmeFile(name, project, desc)

	chrsFile(project)

	fmt.Println(colorize.Colorize("Project Created", colorize.GREEN))
}
