package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/shootingfans/redis-dashboard/internal/locales"
	"github.com/shootingfans/redis-dashboard/internal/logger"
)

const appUniqueId = "com.shootingfans.redis-dashboard"

const (
	preferenceKeyOfSettingLanguage     = "Setting.Language"
	preferenceKeyOfMainAppWindowWidth  = "Size.main.width"
	preferenceKeyOfMainAppWindowHeight = "Size.main.height"
	preferenceKeyOfTheme               = "Setting.Theme"
)

const (
	defaultMainAppWindowsWidth  = 800
	defaultMainAppWindowsHeight = 600
	defaultSettingDialogWidth   = 300
	defaultSettingDialogHeight  = 100
)

// App define application interface
type App interface {
	// Start is start the application
	Start() error

	// Stop is stop the application
	Stop()
}

type guiApp struct {
	startDisable int32
}

func (g *guiApp) initialize() error {
	p := app.NewWithID(appUniqueId)
	locales.SetLanguage(currentLanguage())
	p.Lifecycle().SetOnStarted(func() {
		logger.Info(locales.Get(locales.LOG_INFO_APPLICATION_STARTED))
	})
	p.Lifecycle().SetOnStopped(func() {
		logger.Info(locales.Get(locales.LOG_INFO_APPLICATION_STOPED))
	})
	setAppTheme()
	g.renderMain()
	return nil
}

func (g *guiApp) renderMain() {
	main := makeMainWindows()
	//main.SetMaster()
	main.Show()
}

func (g *guiApp) Start() error {
	//g.startDisable = 1
	//for atomic.CompareAndSwapInt32(&g.startDisable, 1, 0) {
	//	fmt.Println("start...", g.startDisable)
	if err := g.initialize(); err != nil {
		return err
	}
	fyne.CurrentApp().Run()
	//}
	return nil
}

func (g *guiApp) Stop() {
	fyne.CurrentApp().Quit()
}

// NewGUI return new gui application
func NewGUI() App {
	return new(guiApp)
}

func rebootMainWindows() {
	main := makeMainWindows()
	main.CenterOnScreen()
	main.Show()
	logger.Info(locales.Get(locales.LOG_INFO_RENEW_RENDER_WINDOWS))
	fyne.CurrentApp().Driver().AllWindows()[0].Close()
}

func currentLanguage() string {
	return fyne.CurrentApp().Preferences().StringWithFallback(preferenceKeyOfSettingLanguage, locales.CurrentLanguage().String())
}
