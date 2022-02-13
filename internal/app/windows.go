package app

import (
	"errors"
	"regexp"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
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
			makeWorkspace(),
			layout.NewSpacer(),
			widget.NewSeparator(),
			makeButtonToolbar(),
		),
	)
	w.SetOnClosed(func() {
		currentApp.EventManager().Trigger(eventNameOfCloseWindow, w.Canvas().Size())()
	})
	w.CenterOnScreen()
	return w
}

func makeToolBar(parent fyne.Window) fyne.CanvasObject {
	rightToolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			makeSettingDialog(parent).Show()
		}),
	)
	leftToolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MenuIcon(), func() {
			currentApp.EventManager().Trigger(eventNameOfToggleLeftMenu, nil)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			makeCreateRedisDialog(parent).Show()
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			// todo delete current control redis
		}),
	)
	return container.NewHBox(leftToolbar, layout.NewSpacer(), rightToolbar)
}

func makeSettingDialog(parent fyne.Window) dialog.Dialog {
	var items []*widget.FormItem
	items = append(items, makeLanguageSelector(), makeLightOrDarkRadioGroup(), makeRecordWindowSize())
	submit := func(confirm bool) {
		if !confirm {
			return
		}
		var windowsReboot bool
		editRecordWindowSize(items[2].Widget.(*widget.Check).Checked)
		windowsReboot = editTheme(items[1].Widget.(*widget.RadioGroup).Selected) || windowsReboot
		// language changed
		// note: language will change other options string ,so this must handle at last
		windowsReboot = editLanguage(items[0].Widget.(*widget.Select).Selected) || windowsReboot
		if windowsReboot {
			currentApp.EventManager().Trigger(eventNameOfRebootWindow, nil)
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

func makeRecordWindowSize() *widget.FormItem {
	s := widget.NewCheck("", nil)
	s.SetChecked(currentRecordWindowSize())
	return widget.NewFormItem(locales.Get(locales.LABEL_RECORD_SIZE), s)
}

func makeWorkspace() fyne.CanvasObject {
	updateButtonLabel := func(selected string) {
		// todo find redis connect states
	}
	hostSelector := widget.NewSelect(nil, updateButtonLabel)
	currentApp.EventManager().FocusOn(eventNameOfRedisHostsChanged, func(_ string, _ interface{}) {
		var hosts []string
		list, _ := redisStore.List()
		for _, item := range list {
			hosts = append(hosts, item.Name)
		}
		hostSelector.Options = hosts
		hostSelector.Refresh()
	})
	controlRedis := func() {}
	controlButton := widget.NewButton(locales.Get(locales.BUTTON_CONNECT), controlRedis)
	leftMenu := container.NewVBox(
		container.NewHBox(hostSelector, controlButton),
	)
	currentApp.EventManager().FocusOn(eventNameOfToggleLeftMenu, func(_ string, _ interface{}) {
		if leftMenu.Visible() {
			leftMenu.Hide()
			return
		}
		leftMenu.Show()
	})
	currentApp.EventManager().Trigger(eventNameOfRedisHostsChanged, nil)
	workspace := container.NewVBox(
		widget.NewLabel("workspace"),
	)
	return container.NewGridWithColumns(2, leftMenu, workspace)
}

func makeButtonToolbar() fyne.CanvasObject {
	return container.NewHBox(
		layout.NewSpacer(),
		widget.NewLabel(appVersion),
	)
}

func makeCreateRedisDialog(parent fyne.Window) dialog.Dialog {
	var items []*widget.FormItem
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder(locales.Get(locales.LABEL_NEW_REDIS_NAME_PLACEHOLDER))
	nameEntry.Validator = func(name string) error {
		if redisStore.Exists(name) {
			return errors.New(locales.Get(locales.ERROR_ALREADY_EXISTS))
		}
		return nil
	}
	items = append(items, widget.NewFormItem(locales.Get(locales.LABEL_NEW_REDIS_NAME), nameEntry))
	endpointEntry := widget.NewEntry()
	endpointEntry.SetPlaceHolder(locales.Get(locales.LABEL_NEW_REDIS_PLACEHOLDER))
	endpointEntry.Validator = func(endpoints string) error {
		return checkEndpoints(strings.Split(endpoints, ",")...)
	}
	items = append(items, widget.NewFormItem(locales.Get(locales.LABEL_NEW_REDIS_ENDPOINT), endpointEntry))
	passwordEntry := widget.NewPasswordEntry()
	items = append(items, widget.NewFormItem(locales.Get(locales.LABEL_NEW_REDIS_PASSWORD), passwordEntry))
	submit := func(confirm bool) {
		if !confirm {
			return
		}
		_, err := redisStore.Add(redisEntry{
			Name:     nameEntry.Text,
			Hosts:    strings.Split(endpointEntry.Text, ","),
			Password: passwordEntry.Text,
		})
		if err != nil {
			dialog.ShowError(err, parent)
			return
		}
		currentApp.EventManager().Trigger(eventNameOfRedisHostsChanged, nil)
	}
	dg := dialog.NewForm(
		locales.Get(locales.TITLE_REDIS_CREATE_WINDOWS),
		locales.Get(locales.BUTTON_CONFIRM),
		locales.Get(locales.BUTTON_CANCEL),
		items, submit, parent)
	dg.Resize(fyne.NewSize(defaultNewRedisDialogWidth, defaultNewRedisDialogHeight))
	return dg
}

func checkEndpoints(endpoints ...string) error {
	if len(endpoints) == 0 {
		return errors.New(locales.Get(locales.ERROR_ENDPOINTS_REQUIRED))
	}
	r := regexp.MustCompile(`^([0-9a-zA-Z-_.]+)(:[0-9]+)?$`)
	for _, endpoint := range endpoints {
		if !r.MatchString(endpoint) {
			return errors.New(locales.Get(locales.ERROR_INCORRECT_ENDPOINT, endpoint))
		}
	}
	return nil
}
