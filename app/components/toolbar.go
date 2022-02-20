package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	funcs "github.com/force4760/esostack/app/functions"
)

//////////////////////////////////////////////////////////////
// TOOLBAR                                                  //
//////////////////////////////////////////////////////////////

// Create the ToolBar for the App
// Play and Copy Btns
func MakeToolBar(in, out binding.String, win fyne.Window) *widget.Toolbar {
	return widget.NewToolbar(
		playBtn(in, out),
		widget.NewToolbarSpacer(),
		docsBtn(out, win),
	)
}

//////////////////////////////////////////////////////////////
// TOOLBAR ITEMS                                            //
//////////////////////////////////////////////////////////////

// Play/Run program toolbar button
func playBtn(in, out binding.String) widget.ToolbarItem {
	return widget.NewToolbarAction(
		theme.MediaPlayIcon(),
		funcs.PlayFunc(in, out),
	)
}

// Show documentation toolbar button
func docsBtn(out binding.String, win fyne.Window) widget.ToolbarItem {
	return widget.NewToolbarAction(
		theme.DocumentIcon(),
		funcs.DocsFunc(win),
	)
}
