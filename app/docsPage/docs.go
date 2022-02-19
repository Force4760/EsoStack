package docs

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Create the documentation page
func DocsPage() *container.Scroll {
	return container.NewVScroll(
		widget.NewAccordion(
			newItem("Math", mathOpDocs),
			newItem("Comparisson", compOpDocs),
			newItem("Logic", logicOpDocs),
			newItem("Stack", stackOpDocs),
			newItem("Functions", funcOpDocs),
		),
	)
}

// Create a single item for the acordion
func newItem(title, content string) *widget.AccordionItem {
	return widget.NewAccordionItem(
		title,
		widget.NewRichTextFromMarkdown(
			content,
		),
	)
}
