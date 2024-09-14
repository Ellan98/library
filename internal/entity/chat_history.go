package entity

import (
	"time"
)

type ChatRecord struct {
	Status        uint      `json:"status" gorm:"comment:'消息状态'"`
	SendUserId    uint64    `json:"send_user_id" gorm:"comment:'发送者user_id'"`
	CreatedAt     time.Time `json:"created_at"gorm:"comment:'发送时间'"`
	ReceiverModel string    `json:"receiver_model" gorm:"comment:'模型类型'"`
	ModelReply    string    `json:"receiver_model" gorm:"comment:'模型回复'"`
	Content       string    `json:"message" gorm:"comment:'消息内容'"`
	Type          string    `json:"type" gorm:"comment:'消息类型  text | img | file'"`
}

func (c *ChatRecord) TableName() string {
	return "chat_records"
}
