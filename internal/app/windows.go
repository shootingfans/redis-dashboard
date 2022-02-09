package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/shootingfans/redis-dashboard/internal/locales"
)

func makeMainWindows() fyne.Window {
	app := fyne.CurrentApp()
	w := app.NewWindow("redis-dashboard")
	w.SetContent(
		container.NewVBox(
			makeToolBar(w),
			widget.NewSeparator(),
		),
	)
	w.Resize(fyne.NewSize(
		float32(app.Preferences().FloatWithFallback(preferenceKeyOfMainAppWindowWidth, defaultMainAppWindowsWidth)),
		float32(app.Preferences().FloatWithFallback(preferenceKeyOfMainAppWindowHeight, defaultMainAppWindowsHeight)),
	))
	w.SetOnClosed(func() {
		app.Preferences().SetFloat(preferenceKeyOfMainAppWindowWidth, float64(w.Canvas().Size().Width))
		app.Preferences().SetFloat(preferenceKeyOfMainAppWindowHeight, float64(w.Canvas().Size().Height))
	})
	return w
}

func makeToolBar(parent fyne.Window) *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			makeSettingDialog(parent).Show()
		}),
	)
}

func makeSettingDialog(parent fyne.Window) dialog.Dialog {
	var items []*widget.FormItem
	items = append(items, makeLanguageSelector(), makeLightOrDarkRadioGroup())
	submit := func(confirm bool) {
		if !confirm {
			return
		}
		var windowsReboot bool
		windowsReboot = editTheme(items[1].Widget.(*widget.RadioGroup).Selected) || windowsReboot
		// language changed
		// note: language will change other options string ,so this must handle at last
		windowsReboot = editLanguage(items[0].Widget.(*widget.Select).Selected) || windowsReboot
		if windowsReboot {
			currentApp.EventManager().Trigger(eventNameOfRebootWindows, nil)
		}
	}
	dia := dialog.NewForm(
		locales.Get(locales.TITLE_SETTING_WINDOWS),
		locales.Get(locales.BUTTON_SAVE),
		locales.Get(locales.BUTTON_CANCEL),
		items, submit, parent,
	)
	dia.Resize(fyne.NewSize(defaultSettingDialogWidth, defaultSettingDialogHeight))
	return dia
}

func makeLanguageSelector() *widget.FormItem {
	s := widget.NewSelect(locales.GetLanguageNameList(), nil)
	s.SetSelected(locales.CurrentLanguageName())
	return widget.NewFormItem(locales.Get(locales.LABEL_SELECT_LANGUAGE), s)
}

func makeLightOrDarkRadioGroup() *widget.FormItem {
	s := widget.NewRadioGroup([]string{locales.Get(locales.LABEL_THEME_LIGHT), locales.Get(locales.LABEL_THEME_DARK)}, nil)
	s.Required = true
	s.Horizontal = true
	s.SetSelected(locales.Get(currentTheme().Tag()))
	return widget.NewFormItem(locales.Get(locales.LABEL_SELECT_THEME), s)
}
