package comm

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"http-win-notice/utils/constant"
	"http-win-notice/utils/setting"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
)

var (
	linkPath string
	appPath  string
)

func init() {
	userInfo, err := user.Current()
	if nil != err {
		log.Fatalln(err)
	}
	linkPath = path.Join(userInfo.HomeDir, constant.LinkSuffix)

	appPath, err = os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
}

func InitLog() {
	logFile, _ := os.OpenFile(setting.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if setting.Config.Debug == 1 {
		os.Stdout = logFile
		os.Stderr = logFile
	}
}

// 创建快捷方式
func createShortcut(source string, target string) error {
	var err error
	err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	if err != nil {
		return err
	}
	defer ole.CoUninitialize()
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}
	defer oleShellObject.Release()
	wShell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer wShell.Release()
	cs, err := oleutil.CallMethod(wShell, "CreateShortcut", target)
	if err != nil {
		return err
	}
	iDispatch := cs.ToIDispatch()
	_, err = oleutil.PutProperty(iDispatch, "TargetPath", source)
	if err != nil {
		return err
	}
	_, err = oleutil.CallMethod(iDispatch, "Save")
	if err != nil {
		return err
	}
	return nil
}

// 创建快捷方式
func MakeShortcut() {
	fmt.Println("appPath=", appPath)
	fmt.Println("linkPath=", linkPath)
	err := createShortcut(appPath, linkPath)
	if err != nil {
		log.Fatalln(err)
	}
}

// 删除快捷方式
func RemoveShortcut() {
	err := os.Remove(linkPath)
	if err != nil {
		log.Fatalln(err)
	}
}

func ShortcutExists() bool {
	if _, err := os.Stat(linkPath); !os.IsNotExist(err) {
		return true
	}
	return false
}

func RestartApp(serverCh chan bool) {
	log.Println("restart start")
	serverCh <- false
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	cmd := exec.Command(exePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("restart success")
	os.Exit(0)
}
