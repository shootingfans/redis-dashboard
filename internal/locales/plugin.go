package locales

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/shootingfans/redis-dashboard/internal/logger"
)

// Plugin is language Plugin
type Plugin interface {
	// Language return the plugin language tag
	Language() language.Tag
	// GetTagList return all tag and it's translate message
	GetTagList() map[Tag]string
	// Name return plugin name
	Name() string
}

// InjectionPlugin is load language plugin
func InjectionPlugin(plugs ...Plugin) error {
	for _, plug := range plugs {
		lang := plug.Language()
		for tag, msg := range plug.GetTagList() {
			if err := message.SetString(lang, string(tag), msg); err != nil {
				return err
			}
		}
		logger.Info(Get(LOG_INFO_PLUGIN_LANGUAGE_LOADED, plug.Name()))
		supportLanguages = append(supportLanguages, [2]string{plug.Name(), lang.String()})
	}
	return nil
}
