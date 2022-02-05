package main

import (
	"flag"
	"github.com/shootingfans/redis-dashboard/internal/locales"
	"github.com/shootingfans/redis-dashboard/internal/plugins"
	"log"
)

var (
	pluginFolder = flag.String("plugin-folder", "./plugins", locales.Get(locales.FLAG_PLUGIN_FOLDER_DESCRIPTION))
)

func main() {
	flag.Parse()
	if err := plugins.Initialize(*pluginFolder); err != nil {
		log.Fatal(err)
	}
}
