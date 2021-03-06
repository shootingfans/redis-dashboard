package locales

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Tag string

const (
	ERROR_OPEN_PLUGIN_FAILED       Tag = "open plugin failed: %s"
	ERROR_INCORRECT_PLUGIN         Tag = "incorrect plugin: %s"
	ERROR_INITIALIZE_PLUGIN_FAILED Tag = "initialize plugin failed: %s"
	ERROR_INCORRECT_PLUGIN_FOLDER  Tag = "incorrect plugin folder: %s"
	ERROR_START_APPLICATION_FAILED Tag = "start application failed: %s"
	ERROR_ALREADY_EXISTS           Tag = "name %s already exists"
	ERROR_SAVE_FAILED              Tag = "save failed: %s"
	ERROR_REMOVE_FAILED            Tag = "delete failed: %s"
	ERROR_MARSHAL_FAILED           Tag = "marshal failed"
	ERROR_UNMARSHAL_FAILED         Tag = "unmarshal failed"
	ERROR_ENDPOINTS_REQUIRED       Tag = "endpoints required"
	ERROR_INCORRECT_ENDPOINT       Tag = "incorrect endpoint: %s"

	LOG_INFO_PLUGIN_DISCOVERY       Tag = "plugin %s discovery"
	LOG_INFO_PLUGIN_LANGUAGE_LOADED Tag = "language %s loaded"
	LOG_INFO_APPLICATION_STARTED    Tag = "application started"
	LOG_INFO_APPLICATION_STOPED     Tag = "application stopped"
	LOG_INFO_LANGUAGE_CHANGED       Tag = "language changed %s => %s"
	LOG_INFO_THEME_CHANGED          Tag = "theme changed %s => %s"
	LOG_INFO_RENEW_RENDER_WINDOWS   Tag = "renew render windows"

	FLAG_PLUGIN_FOLDER_DESCRIPTION Tag = "Set the plugin folder"

	LABEL_SELECT_LANGUAGE            Tag = "Language"
	LABEL_SELECT_THEME               Tag = "Theme"
	LABEL_THEME_DARK                 Tag = "Dark"
	LABEL_THEME_LIGHT                Tag = "Light"
	LABEL_RECORD_SIZE                Tag = "Record Window Size"
	LABEL_NEW_REDIS_NAME             Tag = "Redis Name"
	LABEL_NEW_REDIS_ENDPOINT         Tag = "Redis Endpoint"
	LABEL_NEW_REDIS_PASSWORD         Tag = "Redis Password"
	LABEL_NEW_REDIS_PLACEHOLDER      Tag = "like host:port, multi separate by \",\""
	LABEL_NEW_REDIS_NAME_PLACEHOLDER Tag = "please input a unique name"

	TITLE_SETTING_WINDOWS      Tag = "Setting"
	TITLE_REDIS_CREATE_WINDOWS Tag = "New Redis"

	BUTTON_CONFIRM      Tag = "Confirm"
	BUTTON_SAVE         Tag = "Save"
	BUTTON_CANCEL       Tag = "Cancel"
	BUTTON_EDIT_SETTING Tag = "Edit Setting"
	BUTTON_CONNECT      Tag = "Connect"
	BUTTON_DISCONNECT   Tag = "Disconnect"
)

func init() {
	currentLanguage = language.English
	currentPrinter = message.NewPrinter(currentLanguage)
	InjectionPlugin(new(englishPlugin))
}

var currentLanguage language.Tag
var currentPrinter *message.Printer
var supportLanguages [][2]string

// SetLanguage is change current language to lang
func SetLanguage(lang string) error {
	l, err := language.Parse(lang)
	if err != nil {
		return err
	}
	currentLanguage = l
	currentPrinter = message.NewPrinter(currentLanguage)
	return nil
}

// AllLanguages return all support language
// The [2]string array is like [2]string{pluginName, language.Tag.String()}
func AllLanguages() [][2]string {
	return supportLanguages
}

// Printer return current Printer
func Printer() *message.Printer {
	return currentPrinter
}

// CurrentLanguage return current language
func CurrentLanguage() language.Tag {
	return currentLanguage
}

func CurrentLanguageName() string {
	for _, lang := range supportLanguages {
		if lang[1] == currentLanguage.String() {
			return lang[0]
		}
	}
	return "English"
}

// Get is get translate message of tag and args
func Get(tag Tag, v ...interface{}) string {
	return currentPrinter.Sprintf(string(tag), v...)
}

// GetLanguageNameList return all language names
func GetLanguageNameList() []string {
	var names []string
	for _, lang := range supportLanguages {
		names = append(names, lang[0])
	}
	return names
}

// GetLanguageByName is query language by language name
func GetLanguageByName(name string) string {
	for _, lang := range supportLanguages {
		if lang[0] == name {
			return lang[1]
		}
	}
	return currentLanguage.String()
}
