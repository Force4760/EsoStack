package components

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	funcs "github.com/force4760/esostack/app/functions"
)

//////////////////////////////////////////////////////////////
// TOOLBAR                                                  //
//////////////////////////////////////////////////////////////

func MakeToolBar(in, out binding.String) *widget.Toolbar {
	return widget.NewToolbar(
		playBtn(in, out),
		widget.NewToolbarSeparator(),
		copyBtn(in, out),
		pasteBtn(in, out),
	)
}

//////////////////////////////////////////////////////////////
// TOOLBAR ITEMS                                            //
//////////////////////////////////////////////////////////////

func playBtn(in, out binding.String) widget.ToolbarItem {
	return widget.NewToolbarAction(
		theme.MediaPlayIcon(),
		funcs.PlayFunc(in, out),
	)
}

func copyBtn(in, out binding.String) widget.ToolbarItem {
	return widget.NewToolbarAction(
		theme.ContentCopyIcon(),
		func() {
			fmt.Println("copy")
		},
	)
}

func pasteBtn(in, out binding.String) widget.ToolbarItem {
	return widget.NewToolbarAction(
		theme.ContentPasteIcon(),
		func() {
			fmt.Println("paste")
		},
	)
}
