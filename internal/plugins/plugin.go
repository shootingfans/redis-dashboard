package plugins

import (
	"fmt"
	"github.com/shootingfans/redis-dashboard/internal/locales"
	"github.com/shootingfans/redis-dashboard/internal/logger"
	"path/filepath"
	"plugin"
	"strings"
)

func Initialize(folder string) error {
	list, err := filepath.Glob(strings.TrimSuffix(folder, "/") + "/*.so")
	if err != nil {
		return fmt.Errorf(locales.Get(locales.ERROR_INCORRECT_PLUGIN_FOLDER), err)
	}
	var localesPlugin []locales.Plugin
	for _, item := range list {
		plug, err := plugin.Open(item)
		if err != nil {
			return fmt.Errorf(locales.Get(locales.ERROR_OPEN_PLUGIN_FAILED), err)
		}
		ins, err := buildPluginInstance(plug)
		if err != nil {
			return err
		}
		logger.Info(locales.Get(locales.LOG_INFO_PLUGIN_DISCOVERY, getPluginName(plug, item)))
		switch tp := ins.(type) {
		case locales.Plugin:
			localesPlugin = append(localesPlugin, tp)
		}
	}
	if err := locales.InjectionPlugin(localesPlugin...); err != nil {
		return fmt.Errorf(locales.Get(locales.ERROR_INITIALIZE_PLUGIN_FAILED), err)
	}
	return nil
}

func getPluginName(p *plugin.Plugin, filename string) string {
	name, _ := p.Lookup("Name")
	if name, ok := name.(*string); ok {
		return *name
	}
	return filepath.Base(filename[0 : len(filename)-3])
}

func buildPluginInstance(p *plugin.Plugin) (interface{}, error) {
	sym, err := p.Lookup("Instance")
	if err != nil {
		return nil, fmt.Errorf(locales.Get(locales.ERROR_INCORRECT_PLUGIN), err)
	}
	ins, ok := sym.(func() interface{})
	if !ok {
		return nil, fmt.Errorf(locales.Get(locales.ERROR_INCORRECT_PLUGIN), "plugin Instance func must like func() interface{}")
	}
	return ins(), nil
}
