package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

//////////////////////////////////////////////////////////////
// EDITOR PAGE                                              //
//////////////////////////////////////////////////////////////

func editor(data binding.String) *widget.Entry {
	input := widget.NewMultiLineEntry()

	input.Bind(data)

	input.SetPlaceHolder("Enter Code ...")

	input.Wrapping = fyne.TextWrapWord

	return input
}

func MakeEditorPage(in, out binding.String) *container.Split {
	input := editor(in)

	content := container.NewVSplit(
		input,
		container.NewHScroll(
			widget.NewLabelWithData(
				out,
			),
		),
	)

	content.SetOffset(0.9)

	return content
}
