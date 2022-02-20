package docs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

//////////////////////////////////////////////////////////////
// DOCS                                                     //
//////////////////////////////////////////////////////////////

// Create the documentation page
func DocsPage() *widget.List {
	docData := allOps
	return widget.NewList(
		func() int {
			return len(docData)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(docData[i].OneLineString())
		})
}
