package client

import (
	"github.com/shanghaobo/go-http-forward/client"
	"http-win-notice/utils/setting"
	"strconv"
)

func StartClient() {
	if setting.Config.Forward.Enable != 1 {
		return
	}
	go func() {
		forward := setting.Config.Forward
		client.Start(forward.ServerAddr, strconv.Itoa(forward.ServerPort), forward.Token, "http://127.0.0.1:"+strconv.Itoa(setting.Config.Port)+"/api/toast")
	}()
}
