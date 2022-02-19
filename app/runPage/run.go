package runpage

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MakeRun() *container.Scroll {
	return container.NewVScroll(
		widget.NewTextGridFromString(
			fmt.Sprintf(
				"kh jhjj jh\n kh jhkh\n hasjhf %d",
				1,
			),
		),
	)
}
