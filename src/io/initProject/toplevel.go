package initProject

import (
	"fmt"
	"io/ioutil"
)

// Create the project ReadMe
func readmeFile(name, project, desc string) {
	readmeText := []byte(
		fmt.Sprintf(
			"# %s\n\n## Project Information\n* Author: %s\n* Description: %s\n* Language: EsoStack",
			project,
			name,
			desc,
		),
	)

	ioutil.WriteFile(project+"/README.md", readmeText, 0644)
}

// Create the Characters file
func chrsFile(folder string) {
	docsText := []byte(`# Operations

Arithmetic    ->  + - * / % ^

Comparisson   ->  < <= > >= == !=

Logic         ->  and  or  xor  not
              ->  &&   ||  XX   !!

Math          ->  sqrt  fact  floor  ceil  abs  max  min  sim  exp  log  sin  cos  tan
              ->  √x    x!    ⌊x⌋     ⌈x⌉   |x|   ↑    ↓    ±  

Stack         -> swap  drop  dup  dup2  rot  over size

Control Flow  -> if{}{}  for{}`)

	ioutil.WriteFile(folder+"/ops.md", docsText, 0644)
}

// Create the function log file
func funcLogFile(folder string) {
	funcText := []byte("# Functions\n")

	ioutil.WriteFile(folder+"/funcLog.md", funcText, 0644)
}
