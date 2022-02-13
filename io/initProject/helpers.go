package initProject

import (
	"os"
	"strings"

	"github.com/liamg/flinch/widgets"
)

// Create a directory at the specified path
func makeDir(path string) error {
	return os.Mkdir(path, 0755)
}

// Get name, projec and description from the user
func get_input() (string, string, string) {
	name, _ := widgets.Input("Full Name: ")
	project, _ := widgets.Input("Project Name: ")
	description, _ := widgets.Input("Short Description: ")

	// replace spaces
	project = strings.ReplaceAll(project, " ", "_")

	return name, project, description
}
