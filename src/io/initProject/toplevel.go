package initProject

import (
	"fmt"
	"io/ioutil"
)

// Create the main file
func mainFile(folder string) {
	mainText := []byte(`

}─┐   ┌──┐       ┌─────A───{
  │ ┌─┤  ├─A──┐ ┌┴┐
  └─┘ └──┘    │ └┬┘
              └──┘`)

	ioutil.WriteFile(folder+"/main.pipes", mainText, 0644)
}

// Create the project ReadMe
func readmeFile(name, project, desc string) {
	readmeText := []byte(
		fmt.Sprintf(
			"# %s\n\n## Project Information\n* Author: %s\n* Description: %s\n* Language: Pipes",
			project,
			name,
			desc,
		),
	)

	ioutil.WriteFile(project+"/README.md", readmeText, 0644)
}

// Create the Characters file
func chrsFile(folder string) {
	docsText := []byte(`# Chars
	
Pipes         ->  ─ │ ┌ ┐ └ ┘ ┤ ├ ┴ ┬ ┼

Arithmetic    ->  + - * / % ^

Comparisson   ->  < <= > >= == !=

Logic         ->  and  or  xor  not
              ->  &&   ||  XX   !!

Math          ->  sqrt  fact  floor  ceil  abs  max  min  sim  exp  log  sin  cos  tan
              ->  √x    x!    ⌊x⌋     ⌈x⌉   |x|   ↑    ↓    ±  

Stack         -> swap  drop  dup  dup2  rot  over

Control Flow  -> if{}{}  for{}`)

	ioutil.WriteFile(folder+"/chars.md", docsText, 0644)
}
