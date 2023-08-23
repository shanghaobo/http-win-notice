package api

import (
	"github.com/gin-gonic/gin"
	"http-win-notice/model"
	"http-win-notice/utils/constant"
	"http-win-notice/utils/notice"
	"net/http"
	"strconv"
)

func ToastApi(c *gin.Context) {
	msg := c.Query("msg")
	title := c.Query("title")
	remark := c.Query("remark")
	if msg == "" {
		msg = constant.DefaultMsg
	}
	if title == "" {
		title = constant.DefaultTitle
	}

	data := model.Msg{
		Title:  title,
		Msg:    msg,
		Remark: remark,
	}
	res := model.CreateMsg(&data)
	if !res {
		c.JSON(400, gin.H{
			"message": "fail, db error",
		})
		return
	}
	go func() {
		err := notice.Notice(msg, title)
		if err != nil {
			model.UpdateMsgStatus(data.ID, 9)
		} else {
			model.UpdateMsgStatus(data.ID, 1)
		}
	}()
	c.JSON(200, gin.H{
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
