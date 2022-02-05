package main

import (
	"github.com/shootingfans/redis-dashboard/internal/locales"
	"golang.org/x/text/language"
)

type traditionalChinesePlugin struct{}

func (t traditionalChinesePlugin) Language() language.Tag {
	return language.TraditionalChinese
}

func (t traditionalChinesePlugin) GetTagList() map[locales.Tag]string {
	return map[locales.Tag]string{
		locales.ERROR_OPEN_PLUGIN_FAILED:        "載入插件失敗: %s",
		locales.ERROR_INCORRECT_PLUGIN:          "無效的插件: %s",
		locales.ERROR_INITIALIZE_PLUGIN_FAILED:  "初始化插件失敗: %s",
		locales.FLAG_PLUGIN_FOLDER_DESCRIPTION:  "配置插件目錄",
		locales.LOG_INFO_PLUGIN_DISCOVERY:       "發現插件 %s",
		locales.ERROR_INCORRECT_PLUGIN_FOLDER:   "無效的插件目錄: %s",
		locales.LOG_INFO_PLUGIN_LANGUAGE_LOADED: "語言包 %s 載入",
	}
}

func (t traditionalChinesePlugin) Name() string {
	return "繁體中文"
}

var Name = "繁體中文語言包"

func Instance() interface{} {
	return new(traditionalChinesePlugin)
}
