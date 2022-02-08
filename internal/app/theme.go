package app

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"

	"github.com/shootingfans/redis-dashboard/internal/locales"
)

type Theme string

const (
	lightTheme Theme = "light"
	darkTheme  Theme = "dark"
)

func (t Theme) Tag() locales.Tag {
	if tg, ok := themeTagMapper[t]; ok {
		return tg
	}
	return locales.LABEL_THEME_DARK
}

var themeTagMapper = map[Theme]locales.Tag{
	lightTheme: locales.LABEL_THEME_LIGHT,
	darkTheme:  locales.LABEL_THEME_DARK,
}

var themes = map[Theme]func() fyne.Theme{
	lightTheme: func() fyne.Theme {
		return theme.LightTheme()
	},
	darkTheme: func() fyne.Theme {
		return theme.DarkTheme()
	},
}

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

func newThemeByTheme(t Theme) fyne.Theme {
	if fn, ok := themes[t]; ok {
		return newTheme(fn())
	}
	return newTheme(fyne.CurrentApp().Settings().Theme())
}

func setAppTheme() {
	fyne.CurrentApp().Settings().SetTheme(newThemeByTheme(currentTheme()))
}

func currentTheme() Theme {
	return Theme(fyne.CurrentApp().Preferences().StringWithFallback(preferenceKeyOfTheme, string(darkTheme)))
}

func getThemeByTranslate(translate string) Theme {
	for tm, tg := range themeTagMapper {
		if locales.Get(tg) == translate {
			return tm
		}
	}
	return darkTheme
}
