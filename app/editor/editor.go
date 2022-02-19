package editor

import (
	"fyne.io/fyne/v2/widget"
)

func MakeEditor() *widget.Entry {
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("( Enter Code ... )")

	return input
}
