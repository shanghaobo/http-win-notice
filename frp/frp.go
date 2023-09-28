package frp

import (
	"http-win-notice/utils"
	"http-win-notice/utils/log"
	"http-win-notice/utils/setting"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const FrpConfigTemplate = `[common]
server_addr = ${ServerAddr}
server_port = ${ServerPort}
token = ${Token}

[ssh]
type = tcp
local_ip = 127.0.0.1
local_port = ${LocalPort}
remote_port = ${RemotePort}`

func initFrpConfig() {
	var config = map[string]string{
		"ServerAddr": setting.Config.Frp.ServerAddr,
		"ServerPort": strconv.Itoa(setting.Config.Frp.ServerPort),
		"Token":      setting.Config.Frp.Token,
		"LocalPort":  strconv.Itoa(setting.Config.Port),
		"RemotePort": strconv.Itoa(setting.Config.Frp.RemotePort),
	}
	s := os.Expand(FrpConfigTemplate, func(k string) string {
		return config[k]
	})
	configPath := path.Join(utils.RootDir, "frpc.ini")
	err := os.WriteFile(configPath, []byte(s), 0644)
	if err != nil {
		log.Fatalln("frp配置生成失败", err)
	}
}

func StartFrp() {
	go func() {
		if setting.Config.Frp.Enable != 1 {
			return
		}
		initFrpConfig()
		frpExePath := path.Join(utils.RootDir, "frpc.exe")
		if _, err := os.Stat(frpExePath); os.IsNotExist(err) {
			log.Println("启动frpc失败: " + frpExePath + "不存在")
			return
		}
		configPath := path.Join(utils.RootDir, "frpc.ini")
		cmd := exec.Command(frpExePath, "-c", configPath)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		log.Println("start launch frpc")
		err := cmd.Start()
		if err != nil {
			log.Fatalln("frp启动失败", err)
		}
	}()

}
