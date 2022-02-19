package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/force4760/esostack/app/appTheme"
	docs "github.com/force4760/esostack/app/docsPage"
	"github.com/force4760/esostack/app/editor"
	runpage "github.com/force4760/esostack/app/runPage"
)

func RunApp() {
	myApp := app.New()
	myApp.Settings().SetTheme(&appTheme.Theme{})
	myWindow := myApp.NewWindow("StackIT")

	content := container.NewAppTabs(
		container.NewTabItem(
			"Edit",
			editor.MakeEditor(),
		),
		container.NewTabItem(
			"Run",
			runpage.MakeRun(),
		),
		container.NewTabItem(
			"Docs",
			docs.DocsPage(),
		),
	)

	myWindow.Resize(
		fyne.NewSize(360, 440),
	)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
