package main

import (
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"http-win-notice/forward/client"
	"http-win-notice/frp"
	"http-win-notice/makelnk"
	"http-win-notice/model"
	"http-win-notice/utils"
	"http-win-notice/utils/comm"
	"http-win-notice/utils/icon"
	"http-win-notice/utils/log"
	"http-win-notice/utils/setting"
	"http-win-notice/web"
)

func init() {
	model.InitDb()
	makelnk.RegisterAppUserModeId()
	client.StartClient()
	frp.StartFrp()
}

func onReady() {
	go func() {
		serverCh := make(chan bool)
		go web.ServerManage(serverCh)
		systray.SetIcon(icon.Data)
		systray.SetTitle("HttpWinNotice")
		systray.SetTooltip("Http消息通知")
		versionMenu := systray.AddMenuItem("版本 "+setting.Version, "版本号")
		versionMenu.Disable()
		systray.AddSeparator()
		restartMenu := systray.AddMenuItem("重启程序", "重启程序")
		serverCh <- true
		//clientRunMenu := systray.AddMenuItemCheckbox("服务启动", "服务启动", true)
		webMenu := systray.AddMenuItem("查看记录", "点击打开web页面")
		settingMenu := systray.AddMenuItem("配置文件", "配置文件")
		settingDirMenu := systray.AddMenuItem("配置目录", "配置目录")
		autoLunchFlag := false
		if comm.ShortcutExists() {
			autoLunchFlag = true
		}
		autoLunchMenu := systray.AddMenuItemCheckbox("开机启动", "开机启动", autoLunchFlag)
		systray.AddSeparator()

		quitMenu := systray.AddMenuItem("退出", "Quit the whole app")

		for {
			select {
			case <-webMenu.ClickedCh:
				comm.OpenHomePage()
			case <-autoLunchMenu.ClickedCh:
				if autoLunchMenu.Checked() {
					autoLunchMenu.Uncheck()
					comm.RemoveShortcut()
				} else {
					autoLunchMenu.Check()
					comm.MakeShortcut()
				}
			case <-settingMenu.ClickedCh:
				open.Run(setting.ConfigPath)
			case <-settingDirMenu.ClickedCh:
				open.Run(utils.RootDir)
			//case <-clientRunMenu.ClickedCh:
			//	if clientRunMenu.Checked() {
			//		clientRunMenu.Uncheck()
			//		serverCh <- false
			//	} else {
			//		clientRunMenu.Check()
			//		serverCh <- true
			//	}
			case <-restartMenu.ClickedCh:
				comm.RestartApp(serverCh)
			case <-quitMenu.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

}

func onExit() {
	// clean up here
	log.Debug("app exit success")
}

func main() {
	systray.Run(onReady, onExit)
}
