package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/force4760/esostack/app/appTheme"
	"github.com/force4760/esostack/app/components"
)

// Create the Input and Output data bindings
func createBindings() (binding.String, binding.String) {
	// Input binding
	inputStr := binding.NewString()
	inputStr.Set("")

	// Output Binding
	outputStr := binding.NewString()
	outputStr.Set("")

	return inputStr, outputStr
}

// Create the content of the app
func createContent(win fyne.Window) *fyne.Container {
	in, out := createBindings()

	// Input Editor
	input := components.MakeEditorPage(in, out)

	// App Toolbar
	toolbar := components.MakeToolBar(in, out, win)

	content := container.NewBorder(
		toolbar,
		nil, nil, nil,
		input,
	)

	return content
}

// Setup the App and the Window
func setupWin() fyne.Window {
	// Create the app and window
	myApp := app.New()
	myWindow := myApp.NewWindow("Stack")

	// Setup the app and window
	myApp.Settings().SetTheme(
		&appTheme.Theme{},
	)
	myWindow.Resize(
		appTheme.FavSize,
	)

	return myWindow
}
