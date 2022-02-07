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
		ERROR_INCORRECT_PLUGIN_FOLDER:   string(ERROR_INCORRECT_PLUGIN_FOLDER),
		ERROR_START_APPLICATION_FAILED:  string(ERROR_START_APPLICATION_FAILED),
		LOG_INFO_PLUGIN_DISCOVERY:       string(LOG_INFO_PLUGIN_DISCOVERY),
		LOG_INFO_PLUGIN_LANGUAGE_LOADED: string(LOG_INFO_PLUGIN_LANGUAGE_LOADED),
		LOG_INFO_APPLICATION_STARTED:    string(LOG_INFO_APPLICATION_STARTED),
		LOG_INFO_APPLICATION_STOPED:     string(LOG_INFO_APPLICATION_STOPED),
		FLAG_PLUGIN_FOLDER_DESCRIPTION:  string(FLAG_PLUGIN_FOLDER_DESCRIPTION),
		LABEL_SELECT_LANGUAGE:           string(LABEL_SELECT_LANGUAGE),
		TITLE_SETTING_WINDOWS:           string(TITLE_SETTING_WINDOWS),
		BUTTON_CONFIRM:                  string(BUTTON_CONFIRM),
		BUTTON_SAVE:                     string(BUTTON_SAVE),
		BUTTON_CANCEL:                   string(BUTTON_CANCEL),
		BUTTON_EDIT_SETTING:             string(BUTTON_EDIT_SETTING),
	}
}
