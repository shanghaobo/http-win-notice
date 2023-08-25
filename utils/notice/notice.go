package notice

import (
	"github.com/go-toast/toast"
	"http-win-notice/utils/constant"
	"http-win-notice/utils/setting"
)

func Notice(msg, title string) error {
	notification := toast.Notification{
		AppID:   constant.AppID,
		Title:   title,
		Message: msg,
		Icon:    setting.LogoPath,
		Actions: []toast.Action{
			{"protocol", "查看", setting.HomeUrl()},
		},
	}
	err := notification.Push()
	return err
	//if err != nil {
	//	panic(err)
	//}
}
