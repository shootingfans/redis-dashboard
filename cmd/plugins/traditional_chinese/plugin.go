package main

import (
	"golang.org/x/text/language"

	"github.com/shootingfans/redis-dashboard/internal/locales"
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
		locales.ERROR_INCORRECT_PLUGIN_FOLDER:   "無效的插件目錄: %s",
		locales.ERROR_START_APPLICATION_FAILED:  "啓動應用失敗: %s",
		locales.LOG_INFO_PLUGIN_DISCOVERY:       "發現插件 %s",
		locales.LOG_INFO_PLUGIN_LANGUAGE_LOADED: "語言包 %s 載入",
		locales.LOG_INFO_APPLICATION_STARTED:    "應用已啓動",
		locales.LOG_INFO_APPLICATION_STOPED:     "應用已停止",
		locales.LOG_INFO_LANGUAGE_CHANGED:       "語言變化 %s => %s",
		locales.LOG_INFO_THEME_CHANGED:          "主題變化 %s => %s",
		locales.LOG_INFO_RENEW_RENDER_WINDOWS:   "重新渲染窗口",
		locales.FLAG_PLUGIN_FOLDER_DESCRIPTION:  "配置插件目錄",
		locales.LABEL_SELECT_LANGUAGE:           "語言",
		locales.LABEL_SELECT_THEME:              "主題",
		locales.LABEL_THEME_LIGHT:               "淺色",
		locales.LABEL_THEME_DARK:                "深色",
		locales.TITLE_SETTING_WINDOWS:           "配置",
		locales.BUTTON_CONFIRM:                  "確認",
		locales.BUTTON_SAVE:                     "保存",
		locales.BUTTON_CANCEL:                   "取消",
		locales.BUTTON_EDIT_SETTING:             "修改配置",
	}
}

func (t traditionalChinesePlugin) Name() string {
	return "繁體中文"
}

// Name is export the plugin name
var Name = "繁體中文語言包"

// Instance is return a plugin instance
func Instance() interface{} {
	return new(traditionalChinesePlugin)
}
