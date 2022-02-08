package app

import (
	"fyne.io/fyne/v2"

	"github.com/shootingfans/redis-dashboard/internal/locales"
	"github.com/shootingfans/redis-dashboard/internal/logger"
)

func editLanguage(selected string) bool {
	if s := locales.GetLanguageByName(selected); s != locales.CurrentLanguage().String() {
		logger.Info(locales.Get(locales.LOG_INFO_LANGUAGE_CHANGED, locales.CurrentLanguageName(), selected))
		locales.SetLanguage(s)
		fyne.CurrentApp().Preferences().SetString(preferenceKeyOfSettingLanguage, s)
		return true
	}
	return false
}

func editTheme(selected string) bool {
	if s := getThemeByTranslate(selected); s != currentTheme() {
		logger.Info(locales.Get(locales.LOG_INFO_THEME_CHANGED, locales.Get(currentTheme().Tag()), selected))
		fyne.CurrentApp().Preferences().SetString(preferenceKeyOfTheme, string(s))
		setAppTheme()
		return true
	}
	return false
}
