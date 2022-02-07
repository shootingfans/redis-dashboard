package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/shootingfans/redis-dashboard/internal/locales"
	"golang.org/x/text/language"
)

func makeMainWindows() fyne.Window {
	app := fyne.CurrentApp()
	w := app.NewWindow("redis-dashboard")
	w.SetContent(
		container.NewVBox(
			widget.NewButton(locales.Get(locales.BUTTON_EDIT_SETTING), func() {
				makeSettingDialog(w).Show()
			}),
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

func makeToolBar() {
	widget.NewToolbar()
}

func makeSettingDialog(parent fyne.Window) dialog.Dialog {
	var items []*widget.FormItem
	items = append(items, makeLanguageSelector())
	submit := func(confirm bool) {
		if !confirm {
			return
		}
		// language changed
		lan := items[0].Widget.(*widget.Select).Selected
		for _, lang := range locales.AllLanguages() {
			if lang[0] == lan {
				if locales.CurrentLanguage().String() != lang[1] {
					locales.SetLanguage(lang[1])
					fyne.CurrentApp().Preferences().SetString(preferenceKeyOfSettingLanguage, lang[1])
					rebootMainWindows()
				}
				break
			}
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
	languages := locales.AllLanguages()
	var options []string
	var selected string
	nowLanguage := fyne.CurrentApp().Preferences().StringWithFallback(preferenceKeyOfSettingLanguage, language.English.String())
	for _, lang := range languages {
		options = append(options, lang[0])
		if lang[1] == nowLanguage {
			selected = lang[0]
		}
	}
	s := widget.NewSelect(options, nil)
	s.SetSelected(selected)
	return widget.NewFormItem(locales.Get(locales.LABEL_SELECT_LANGUAGE), s)
}
