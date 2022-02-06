package main

import (
	"flag"
	"github.com/shootingfans/redis-dashboard/internal/app"
	"github.com/shootingfans/redis-dashboard/internal/locales"
	"github.com/shootingfans/redis-dashboard/internal/logger"
	"github.com/shootingfans/redis-dashboard/internal/plugins"
	"os"
)

var (
	pluginFolder = flag.String("plugin-folder", "./plugins", locales.Get(locales.FLAG_PLUGIN_FOLDER_DESCRIPTION))
)

func main() {
	flag.Parse()
	if err := plugins.Initialize(*pluginFolder); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	gui := app.NewGUI()
	if err := gui.Start(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
