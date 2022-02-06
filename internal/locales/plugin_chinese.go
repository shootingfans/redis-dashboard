package locales

import "golang.org/x/text/language"

type simplifiedChinesePlugin struct{}

func (c simplifiedChinesePlugin) Name() string {
	return "简体中文"
}

func (c simplifiedChinesePlugin) Language() language.Tag {
	return language.SimplifiedChinese
}

func (c simplifiedChinesePlugin) GetTagList() map[Tag]string {
	return map[Tag]string{
		ERROR_OPEN_PLUGIN_FAILED:        "载入插件失败: %s",
		ERROR_INCORRECT_PLUGIN:          "无效的插件: %s",
		ERROR_INITIALIZE_PLUGIN_FAILED:  "初始化插件失败: %s",
		FLAG_PLUGIN_FOLDER_DESCRIPTION:  "配置插件目录",
		LOG_INFO_PLUGIN_DISCOVERY:       "发现插件 %s",
		ERROR_INCORRECT_PLUGIN_FOLDER:   "无效的插件目录: %s",
		LOG_INFO_PLUGIN_LANGUAGE_LOADED: "语言包 %s 载入",
		ERROR_START_APPLICATION_FAILED:  "启动应用失败: %s",
		LOG_INFO_APPLICATION_STARTED:    "应用已启动",
		LOG_INFO_APPLICATION_STOPED:     "应用已停止",
	}
}

func init() {
	InjectionPlugin(new(simplifiedChinesePlugin))
}
