package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/force4760/esostack/app/appTheme"
	"github.com/force4760/esostack/app/components"
)

func createBindings() (binding.String, binding.String) {
	inputStr := binding.NewString()
	inputStr.Set("")

	outputStr := binding.NewString()
	outputStr.Set("")

	return inputStr, outputStr
}

func setupWin() fyne.Window {
	// Create the app and window
	myApp := app.New()
	myWindow := myApp.NewWindow("Stack")

	// Setup the app and window
	myApp.Settings().SetTheme(
		&appTheme.Theme{},
	)
	myWindow.Resize(
		fyne.NewSize(360, 440),
	)

	return myWindow
}

func createContent() *fyne.Container {
	in, out := createBindings()

	input := components.MakeEditorPage(in, out)

	toolbar := components.MakeToolBar(in, out)

	content := container.NewBorder(
		toolbar,
		nil, nil, nil,
		input,
	)

	return content
}
