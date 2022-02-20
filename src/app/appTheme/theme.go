package appTheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:generate fyne bundle -o appTheme/bundled.go appTheme/Font.ttf

//////////////////////////////////////////////////////////////
// COLOR THEME                                              //
//////////////////////////////////////////////////////////////

type Theme struct{}

func (m *Theme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return &color.NRGBA{R: 0x2E, G: 0x34, B: 0x40, A: 0xFF}

	case theme.ColorNameForeground:
		return &color.NRGBA{R: 0xD8, G: 0xDE, B: 0xE9, A: 0xFF}

	case theme.ColorNamePrimary:
		return &color.NRGBA{R: 0x5E, G: 0x81, B: 0xAC, A: 0xAA}

	case theme.ColorNameFocus:
		return &color.NRGBA{R: 0x81, G: 0xA1, B: 0xC1, A: 0x66}

	case theme.ColorNameInputBackground:
		return &color.NRGBA{R: 0x4C, G: 0x56, B: 0x6A, A: 0x66}

	}

	return theme.DefaultTheme().Color(n, v)
}

func (m *Theme) Font(s fyne.TextStyle) fyne.Resource {
	return resourceFontTtf
}

func (m *Theme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *Theme) Size(n fyne.ThemeSizeName) float32 {
	if n == theme.SizeNameText {
		return theme.DefaultTheme().Size(n) + 4
	}

	return theme.DefaultTheme().Size(n)
}

// Window size
var FavSize = fyne.NewSize(360, 440)
