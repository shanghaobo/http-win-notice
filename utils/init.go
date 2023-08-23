package utils

import (
	"log"
	"os"
	"os/user"
	"path"
)

var (
	UserDir string
	RootDir string
)

func init() {
	userInfo, err := user.Current()
	if nil != err {
		log.Fatalln(err)
	}
	UserDir = userInfo.HomeDir
	RootDir = path.Join(UserDir, "http-win-notice")
	if _, err := os.Stat(RootDir); os.IsNotExist(err) {
		err = os.Mkdir(RootDir, os.ModeDir)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
