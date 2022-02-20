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

// Create the full Editor pgae
// Split layout with an Entry and a Label
func MakeEditorPage(in, out binding.String) *container.Split {
	input := editor(in)
	output := label(out)

	content := container.NewVSplit(
		input,
		output,
	)

	content.SetOffset(0.9)

	return content
}

//////////////////////////////////////////////////////////////
// EDITOR COMPONENTS                                        //
//////////////////////////////////////////////////////////////

// Create the Output Label
func label(data binding.String) *container.Scroll {
	return container.NewHScroll(
		widget.NewLabelWithData(
			data,
		),
	)
}

// Create the input Editor
func editor(data binding.String) *widget.Entry {
	input := widget.NewMultiLineEntry()

	input.Bind(data)

	input.SetPlaceHolder("Enter Code ...")

	input.Wrapping = fyne.TextWrapWord

	return input
}
