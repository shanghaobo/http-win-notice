package comm

import (
	"http-win-notice/utils/constant"
	"http-win-notice/utils/setting"
	"log"
	"os/exec"
	"strconv"
	"syscall"
)

func openUrl(url string) {
	cmd := exec.Command(`cmd`, `/c`, `start`, url)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Start()
	if err != nil {
		log.Println("url打开失败：", url)
	}
}

func OpenHomePage() {
	port := setting.Config.Port
	url := constant.BaseUrl + ":" + strconv.Itoa(port)
	openUrl(url)
}
