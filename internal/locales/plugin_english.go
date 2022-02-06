package locales

import "golang.org/x/text/language"

type englishPlugin struct{}

func (e englishPlugin) Name() string {
	return "English"
}

func (e englishPlugin) Language() language.Tag {
	return language.English
}

func (e englishPlugin) GetTagList() map[Tag]string {
	return map[Tag]string{
		ERROR_OPEN_PLUGIN_FAILED:        string(ERROR_OPEN_PLUGIN_FAILED),
		ERROR_INCORRECT_PLUGIN:          string(ERROR_INCORRECT_PLUGIN),
		ERROR_INITIALIZE_PLUGIN_FAILED:  string(ERROR_INITIALIZE_PLUGIN_FAILED),
		FLAG_PLUGIN_FOLDER_DESCRIPTION:  string(FLAG_PLUGIN_FOLDER_DESCRIPTION),
		LOG_INFO_PLUGIN_DISCOVERY:       string(LOG_INFO_PLUGIN_DISCOVERY),
		ERROR_INCORRECT_PLUGIN_FOLDER:   string(ERROR_INCORRECT_PLUGIN_FOLDER),
		LOG_INFO_PLUGIN_LANGUAGE_LOADED: string(LOG_INFO_PLUGIN_LANGUAGE_LOADED),
		ERROR_START_APPLICATION_FAILED:  string(ERROR_START_APPLICATION_FAILED),
	}
}
