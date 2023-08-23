package comm

import (
	"http-win-notice/utils/constant"
	"http-win-notice/utils/setting"
	"os/exec"
	"strconv"
)

func openUrl(url string) {
	exec.Command(`cmd`, `/c`, `start`, url).Start()
}

func OpenHomePage() {
	port := setting.Config.Port
	url := constant.BaseUrl + ":" + strconv.Itoa(port)
	openUrl(url)
}
