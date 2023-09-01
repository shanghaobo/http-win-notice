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
var ImagesDir string

type ConfigType struct {
	Port    int         `yaml:"port"`
	Toast   ToastType   `yaml:"toast"`
	Frp     FrpType     `yaml:"frp"`
	Forward ForwardType `yaml:"forward"`
}

type ToastType struct {
	Icons []string `yaml:"icons"`
}

type FrpType struct {
	Enable     int    `yaml:"enable"`
	ServerAddr string `yaml:"server_addr"`
	ServerPort int    `yaml:"server_port"`
	Token      string `yaml:"token"`
	RemotePort int    `yaml:"remote_port"`
}

type ForwardType struct {
	Enable int    `yaml:"enable"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Token  string `yaml:"token"`
}

func init() {
	ConfigPath = path.Join(utils.RootDir, constant.ConfigFile)
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		initConfigFile(ConfigPath)
	}
	InitConfig()
	LogPath = path.Join(utils.RootDir, "log.log")
	DbPath = path.Join(utils.RootDir, constant.DbFile)

	ImagesDir = path.Join(utils.RootDir, "images")
	if _, err := os.Stat(ImagesDir); os.IsNotExist(err) {
		err = os.Mkdir(ImagesDir, os.ModeDir)
		if err != nil {
			log.Fatalln(err)
		}
	}

	LogoPath = path.Join(ImagesDir, "logo.png")
	if _, err := os.Stat(LogoPath); os.IsNotExist(err) {
		b64.WriteImgFileByBase64(constant.LogoImgBase64, LogoPath)
	}

}

func initConfigFile(ConfigPath string) {
	//默认配置
	Config.Port = 19000
	Frp := FrpType{
		Enable:     0,
		ServerAddr: "127.0.0.1",
		ServerPort: 7000,
		Token:      "httpwinnotice123456",
		RemotePort: 19001,
	}
	Config.Frp = Frp

	icons := []string{"logo.png"}
	Toast := ToastType{
		Icons: icons,
	}
	Config.Toast = Toast

	//默认转发配置
	Forward := ForwardType{
		Enable: 0,
		Host:   "127.0.0.1",
		Port:   9919,
		Token:  "httpwinnotice123456",
	}
	Config.Forward = Forward

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
}

func PortStr() string {
	return strconv.Itoa(Config.Port)
}

func HomeUrl() string {
	return constant.BaseUrl + ":" + PortStr()
}
