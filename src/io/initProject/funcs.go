package initProject

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/force4760/esostack/src/io/colorize"
	"github.com/liamg/flinch/widgets"
)

// Create the Characters file
func FuncFile() {
	name, _ := widgets.Input("Function Name: ")
	if name == "" {
		fmt.Println(
			colorize.Error(
				"The provided Name is not valid",
			),
		)
		return
	}

	desc, _ := widgets.Input("Function Description: ")

	docsText := []byte(fmt.Sprintf(`( Description: %s 
	Args:
		(arg-1-name): <arg-1-desc>
	
	Vals:
		(val-1-name): <val-1-desc>
)`, desc))

	err := ioutil.WriteFile(
		fmt.Sprintf("./funcs/%s.stk", name),
		docsText,
		0644,
	)

	if err != nil {
		fmt.Println(
			colorize.Error("There is no function folder in this directory."),
		)
		return
	}

	append(name, desc)

}

func append(name, desc string) {
	file, err := os.OpenFile("./funcLog.md", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(
			colorize.Error("No Function Log file."),
		)
		return
	}

	defer file.Close()

	funcLog := fmt.Sprintf(
		"\n* %s: %s\n",
		name, desc,
	)

	_, err = file.WriteString(funcLog)
	if err != nil {
		fmt.Println(
			colorize.Error("Error witing to the Function Log file."),
		)
		return
	}

	fmt.Println(
		colorize.Colorize(
			"Function Created.",
			colorize.GREEN,
		),
	)
}
