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
)

func init() {
	currentPrinter = message.NewPrinter(language.English)
	InjectionPlugin(new(englishPlugin))
}

var currentPrinter *message.Printer
var supportLanguages [][2]string

func SetLanguage(lang string) {
	l, err := language.Parse(lang)
	if err != nil {
		return
	}
	currentPrinter = message.NewPrinter(l)
}

func AllLanguages() [][2]string {
	return supportLanguages
}

func Printer() *message.Printer {
	return currentPrinter
}

func Get(tag Tag, v ...interface{}) string {
	return currentPrinter.Sprintf(string(tag), v...)
}
