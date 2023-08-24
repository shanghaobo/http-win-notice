package frp

import (
	"embed"
	"fmt"
	"http-win-notice/utils"
	"http-win-notice/utils/setting"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
)

//go:embed frpc.exe
var frpClient embed.FS

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

func initFrp() {
	frpPath := path.Join(utils.RootDir, "frpc.exe")
	if _, err := os.Stat(frpPath); os.IsNotExist(err) {
		frpExeData, _ := frpClient.ReadFile("frpc.exe")
		err = os.WriteFile(frpPath, frpExeData, 0755)
		if err != nil {
			log.Fatalln("frp创建失败", err)
		}
	}

}

func StartFrp() {
	go func() {
		if setting.Config.Frp.Enable != 1 {
			return
		}
		initFrp()
		initFrpConfig()
		frpExePath := path.Join(utils.RootDir, "frpc.exe")
		configPath := path.Join(utils.RootDir, "frpc.ini")
		fmt.Println("frp path=", frpExePath)
		cmd := exec.Command(frpExePath, "-c", configPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalln("frp启动失败", err)
		}
	}()

}
