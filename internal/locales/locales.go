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

	LOG_INFO_PLUGIN_DISCOVERY       Tag = "plugin %s discovery"
	LOG_INFO_PLUGIN_LANGUAGE_LOADED Tag = "language %s loaded"
	LOG_INFO_APPLICATION_STARTED    Tag = "application started"
	LOG_INFO_APPLICATION_STOPED     Tag = "application stopped"

	FLAG_PLUGIN_FOLDER_DESCRIPTION Tag = "Set the plugin folder"

	LABEL_SELECT_LANGUAGE Tag = "Language"

	TITLE_SETTING_WINDOWS Tag = "Setting"

	BUTTON_CONFIRM      Tag = "Confirm"
	BUTTON_SAVE         Tag = "Save"
	BUTTON_CANCEL       Tag = "Cancel"
	BUTTON_EDIT_SETTING Tag = "Edit Setting"
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

// Get is get translate message of tag and args
func Get(tag Tag, v ...interface{}) string {
	return currentPrinter.Sprintf(string(tag), v...)
}
