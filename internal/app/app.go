package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/shootingfans/redis-dashboard/internal/locales"
	"github.com/shootingfans/redis-dashboard/internal/logger"
)

type App interface {
	Start() error
}

type guiApp struct {
}

func (g *guiApp) initialize() error {
	p := app.NewWithID("com.shootingfans.redis-dashboard")
	p.Lifecycle().SetOnStarted(func() {
		logger.Info(locales.Get(locales.LOG_INFO_APPLICATION_STARTED))
	})
	p.Lifecycle().SetOnStopped(func() {
		logger.Info(locales.Get(locales.LOG_INFO_APPLICATION_STOPED))
	})
	return nil
}

func (g *guiApp) Start() error {
	if err := g.initialize(); err != nil {
		return err
	}
	fyne.CurrentApp().Run()
	return nil
}

func (g *guiApp) Stop() {
	fyne.CurrentApp().Quit()
}

func NewGUI() App {
	return new(guiApp)
}
