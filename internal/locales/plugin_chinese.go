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
		ERROR_OPEN_PLUGIN_FAILED:         "载入插件失败: %s",
		ERROR_INCORRECT_PLUGIN:           "无效的插件: %s",
		ERROR_INITIALIZE_PLUGIN_FAILED:   "初始化插件失败: %s",
		ERROR_INCORRECT_PLUGIN_FOLDER:    "无效的插件目录: %s",
		ERROR_START_APPLICATION_FAILED:   "启动应用失败: %s",
		ERROR_ALREADY_EXISTS:             "名称 %s 已存在",
		ERROR_SAVE_FAILED:                "保存失败: %s",
		ERROR_REMOVE_FAILED:              "删除失败: %s",
		ERROR_MARSHAL_FAILED:             "序列化失败",
		ERROR_UNMARSHAL_FAILED:           "反序列化失败",
		ERROR_ENDPOINTS_REQUIRED:         "redis节点地址必须设置",
		ERROR_INCORRECT_ENDPOINT:         "无效的节点: %s",
		LOG_INFO_PLUGIN_DISCOVERY:        "发现插件 %s",
		LOG_INFO_PLUGIN_LANGUAGE_LOADED:  "语言包 %s 载入",
		LOG_INFO_APPLICATION_STARTED:     "应用已启动",
		LOG_INFO_APPLICATION_STOPED:      "应用已停止",
		LOG_INFO_LANGUAGE_CHANGED:        "语言变化 %s => %s",
		LOG_INFO_THEME_CHANGED:           "主题变化 %s => %s",
		LOG_INFO_RENEW_RENDER_WINDOWS:    "重新渲染窗口",
		FLAG_PLUGIN_FOLDER_DESCRIPTION:   "配置插件目录",
		LABEL_SELECT_LANGUAGE:            "语言",
		LABEL_SELECT_THEME:               "主题",
		LABEL_THEME_LIGHT:                "浅色",
		LABEL_THEME_DARK:                 "深色",
		LABEL_RECORD_SIZE:                "记录窗口大小",
		LABEL_NEW_REDIS_NAME:             "Redis名称",
		LABEL_NEW_REDIS_ENDPOINT:         "Redis节点",
		LABEL_NEW_REDIS_PASSWORD:         "Redis密码",
		LABEL_NEW_REDIS_PLACEHOLDER:      "采用 host:port 格式,多个用\",\"分隔",
		LABEL_NEW_REDIS_NAME_PLACEHOLDER: "请输入一个唯一名称做标识",
		TITLE_SETTING_WINDOWS:            "配置",
		TITLE_REDIS_CREATE_WINDOWS:       "新建Redis",
		BUTTON_CONFIRM:                   "确认",
		BUTTON_SAVE:                      "保存",
		BUTTON_CANCEL:                    "取消",
		BUTTON_EDIT_SETTING:              "修改配置",
		BUTTON_CONNECT:                   "连接",
		BUTTON_DISCONNECT:                "断开",
	}
}

func init() {
	InjectionPlugin(new(simplifiedChinesePlugin))
}
