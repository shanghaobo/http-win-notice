package main

import (
	"fmt"
	"github.com/shanghaobo/go-http-forward/server"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

var Config ConfigType

func init() {
	Config = ConfigType{}
	exePath, _ := os.Executable()
	configPath := path.Join(filepath.Dir(exePath), "config.yml")
	dataBytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln("读取配置文件失败")
	}
	fmt.Println("yaml 文件的内容: \n", string(dataBytes))

	err = yaml.Unmarshal(dataBytes, &Config)
	if err != nil {
		fmt.Println("解析config失败")
		log.Fatalln("解析config失败")
	}
}

func main() {
	server.Start(strconv.Itoa(Config.ServerPort), Config.Token, Config.ApiToken, strconv.Itoa(Config.ApiPort))
}
