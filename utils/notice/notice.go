package notice

import (
	"github.com/go-toast/toast"
	"http-win-notice/utils/constant"
	"http-win-notice/utils/setting"
	"path"
	"strconv"
)

func notificationSetAttr(n *toast.Notification, duration string, audio string) {
	switch duration {
	case "long":
		n.Duration = toast.Long
	case "short":
		n.Duration = toast.Short
	default:
		n.Duration = toast.Short
	}

	switch audio {
	case "default":
		n.Audio = toast.Default
	case "im":
		n.Audio = toast.IM
	case "mail":
		n.Audio = toast.Mail
	case "reminder":
		n.Audio = toast.Reminder
	case "sms":
		n.Audio = toast.SMS
	case "loopingalarm":
		n.Audio = toast.LoopingAlarm
	case "loopingalarm2":
		n.Audio = toast.LoopingAlarm2
	case "loopingalarm3":
		n.Audio = toast.LoopingAlarm3
	case "loopingalarm4":
		n.Audio = toast.LoopingAlarm4
	case "loopingalarm5":
		n.Audio = toast.LoopingAlarm5
	case "loopingalarm6":
		n.Audio = toast.LoopingAlarm6
	case "loopingalarm7":
		n.Audio = toast.LoopingAlarm7
	case "loopingalarm8":
		n.Audio = toast.LoopingAlarm8
	case "loopingalarm9":
		n.Audio = toast.LoopingAlarm9
	case "loopingalarm10":
		n.Audio = toast.LoopingAlarm10
	case "loopingcall":
		n.Audio = toast.LoopingCall
	case "loopingcall2":
		n.Audio = toast.LoopingCall2
	case "loopingcall3":
		n.Audio = toast.LoopingCall3
	case "loopingcall4":
		n.Audio = toast.LoopingCall4
	case "loopingcall5":
		n.Audio = toast.LoopingCall5
	case "loopingcall6":
		n.Audio = toast.LoopingCall6
	case "loopingcall7":
		n.Audio = toast.LoopingCall7
	case "loopingcall8":
		n.Audio = toast.LoopingCall8
	case "loopingcall9":
		n.Audio = toast.LoopingCall9
	case "loopingcall10":
		n.Audio = toast.LoopingCall10
	case "silent":
		n.Audio = toast.Silent
	default:
		n.Audio = toast.Default
	}

}

func Notice(msg, title, icon, duration, audio string) error {
	iconIndex, err := strconv.Atoi(icon)
	if err != nil {
		iconIndex = 0
	}

	var iconPath string
	if len(setting.Config.Toast.Icons) <= iconIndex {
		iconPath = path.Join(setting.ImagesDir, "logo.png")
	} else {
		iconPath = path.Join(setting.ImagesDir, setting.Config.Toast.Icons[iconIndex])
	}
	notification := toast.Notification{
		AppID:   constant.AppID,
		Title:   title,
		Message: msg,
		Icon:    iconPath,
		Actions: []toast.Action{
			{"protocol", "查看", setting.HomeUrl()},
		},
	}
	notificationSetAttr(&notification, duration, audio)
	//fmt.Println("n=", notification)
	err = notification.Push()
	return err
	//if err != nil {
	//	panic(err)
	//}
}
