package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"http-win-notice/model"
	"http-win-notice/utils/constant"
	"http-win-notice/utils/notice"
	"net/http"
	"strconv"
)

type Button struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func ToastApi(c *gin.Context) {
	msg := c.DefaultQuery("msg", constant.DefaultMsg)
	title := c.DefaultQuery("title", constant.DefaultTitle)
	icon := c.DefaultQuery("icon", constant.DefaultIcon)
	duration := c.DefaultQuery("duration", string(constant.DefaultDuration))
	audio := c.DefaultQuery("audio", string(constant.DefaultAudio))
	button := c.DefaultQuery("button", "")

	remark := c.Query("remark")

	var buttonData []Button
	var actions []notice.Action
	err := json.Unmarshal([]byte(button), &buttonData)
	if err == nil {
		for _, bt := range buttonData {
			action := notice.Action{
				Type:      "protocol",
				Label:     bt.Name,
				Arguments: bt.Url,
			}
			actions = append(actions, action)
		}
	}
	data := model.Msg{
		Title:  title,
		Msg:    msg,
		Remark: remark,
	}
	res := model.CreateMsg(&data)
	if !res {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fail, db error",
		})
		return
	}
	go func() {
		err := notice.Notice(msg, title, icon, duration, audio, actions)
		if err != nil {
			model.UpdateMsgStatus(data.ID, 9)
		} else {
			model.UpdateMsgStatus(data.ID, 1)
		}
	}()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func HomePage(c *gin.Context) {
	page := c.Query("page")
	pageInt64, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		pageInt64 = 1
	}
	pageInt := int(pageInt64)
	sizeInt := constant.DefaultWebPageSize
	msgList := model.SearchMsg(pageInt, sizeInt)
	formatMsgList := make([]gin.H, len(msgList))
	total := model.TotalMsg()
	for i, msg := range msgList {
		var statusShow string
		switch msg.Status {
		case 1:
			statusShow = "已提醒"
		case 9:
			statusShow = "提醒失败"
		default:
			statusShow = "等待"
		}
		formatMsg := gin.H{
			"id":     msg.ID,
			"index":  (pageInt-1)*10 + i + 1,
			"title":  msg.Title,
			"msg":    msg.Msg,
			"remark": msg.Remark,
			"time":   msg.CreatedAt.Format("2006-01-02 15:04:05"),
			"status": statusShow,
		}
		formatMsgList[i] = formatMsg
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msgList":     formatMsgList,
		"total":       total,
		"currentPage": pageInt,
		"pageSize":    sizeInt,
	})
}
