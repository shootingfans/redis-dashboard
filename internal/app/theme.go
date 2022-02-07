package app

import (
	"fyne.io/fyne/v2"
	"image/color"
)

type customTheme struct {
	parent fyne.Theme
}

func (c *customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return c.parent.Color(name, variant)
}

func (c *customTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return resourceFontsBoldTtf
	}
	return resourceFontsMediumTtf
}

func (c *customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return c.parent.Icon(name)
}

func (c *customTheme) Size(name fyne.ThemeSizeName) float32 {
	return c.parent.Size(name)
}

func newTheme(parent fyne.Theme) fyne.Theme {
	return &customTheme{parent: parent}
}
