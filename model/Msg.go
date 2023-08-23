package model

import "gorm.io/gorm"

type Msg struct {
	gorm.Model
	Title  string `gorm:"type:varchar(50);not null"`
	Msg    string `gorm:"type:varchar(500);not null"`
	Remark string `gorm:"type:text"`
	Status int    `gorm:"type:tinyint;not null;default:0"`
}

func CreateMsg(data *Msg) bool {
	err = db.Create(&data).Error
	if err != nil {
		return false
	}
	return true
}

func UpdateMsgStatus(id uint, status int) bool {
	err = db.Model(&Msg{}).Where("id = ? ", id).Update("status", status).Error
	if err != nil {
		return false
	}
	return true
}

func SearchMsg(page, size int) []Msg {
	var msgList []Msg
	offset := (page - 1) * size
	db.Order("id desc").Limit(size).Offset(offset).Find(&msgList)
	return msgList
}

func TotalMsg() int64 {
	var total int64
	err := db.Model(&Msg{}).Count(&total).Error
	if err != nil {
		return 0
	}
	return total
}
