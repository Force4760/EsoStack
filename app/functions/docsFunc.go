package funcs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/force4760/esostack/app/appTheme"
	docs "github.com/force4760/esostack/app/docsPage"
)

// Show the Documentation Dialog
func DocsFunc(win fyne.Window) func() {
	return func() {
		d := dialog.NewCustom(
			"Docs",
			"Ok",
			docs.DocsPage(),
			win,
		)

		d.Resize(appTheme.FavSize)

		d.Show()
	}
}
