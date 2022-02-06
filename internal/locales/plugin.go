package locales

import (
	"github.com/shootingfans/redis-dashboard/internal/logger"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Plugin interface {
	Language() language.Tag
	GetTagList() map[Tag]string
	Name() string
}

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
