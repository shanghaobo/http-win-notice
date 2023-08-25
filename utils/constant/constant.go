package constant

const (
	LnkName      = "http-win-notice.lnk"
	LinkSuffix   = "AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/" + LnkName
	StartMenuLnk = "AppData/Roaming/Microsoft/Windows/Start Menu/Programs/" + LnkName
	BaseUrl      = "http://127.0.0.1"
	ConfigFile   = "config.yml"
	DbFile       = "data.db"
)

const AppID = "HttpWinNotice"

const (
	DefaultTitle = "消息通知"
	DefaultMsg   = "这是一条测试消息"
)

const DefaultWebPageSize = 10

const LogoImgBase64 = "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\n"
