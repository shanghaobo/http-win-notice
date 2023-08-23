package setting

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"http-win-notice/utils"
	"http-win-notice/utils/b64"
	"http-win-notice/utils/constant"
	"log"
	"os"
	"path"
	"strconv"
)

var ConfigPath string
var Config ConfigType
var LogPath string
var DbPath string
var LogoPath string

type ConfigType struct {
	Mode int `yaml:"mode"`
	Port int `yaml:"port"`
}

func init() {
	ConfigPath = path.Join(utils.RootDir, constant.ConfigFile)
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		initConfigFile(ConfigPath)
	}
	InitConfig()
	LogPath = path.Join(utils.RootDir, "log.log")
	DbPath = path.Join(utils.RootDir, constant.DbFile)
	LogoPath = path.Join(utils.RootDir, "logo.png")
	if _, err := os.Stat(LogoPath); os.IsNotExist(err) {
		b64.WriteImgFileByBase64(constant.LogoImgBase64, LogoPath)
	}
}

func initConfigFile(ConfigPath string) {
	Config.Mode = 1
	Config.Port = 19000
	updatedData, err := yaml.Marshal(Config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("create ConfigPath", ConfigPath)
	err = os.WriteFile(ConfigPath, updatedData, 0644)
}

func InitConfig() {
	Config = ConfigType{}
	dataBytes, err := os.ReadFile(ConfigPath)
	if err != nil {
		log.Fatalln("读取配置文件失败")
	}
	fmt.Println("yaml 文件的内容: \n", string(dataBytes))

	err = yaml.Unmarshal(dataBytes, &Config)
	if err != nil {
		fmt.Println("解析config失败")
		log.Fatalln("解析config失败")
	}
	fmt.Println("config=")
	fmt.Println(Config)
	fmt.Println(Config.Mode)
	fmt.Println(Config.Port)
}

func PortStr() string {
	return strconv.Itoa(Config.Port)
}

func HomeUrl() string {
	return constant.BaseUrl + ":" + PortStr()
}
