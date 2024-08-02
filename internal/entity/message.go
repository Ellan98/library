/*
 * @Date: 2024-07-17 16:39:32
 * @LastEditTime: 2024-08-02 10:44:24
 * @FilePath: \library_room\internal\entity\message.go
 * @description: 注释
 */
package entity

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FormId   int64  `json:"userId"`   //信息发送者
	TargetId int64  `json:"targetId"` //信息接收者
	Type     int    //聊天类型：群聊 私聊 广播
	Media    int    //信息类型：文字 图片 音频
	Content  string //消息内容
	Pic      string `json:"url"` //图片相关
	Url      string //文件相关
	Desc     string //文件描述
	Amount   int    //其他数据大小

}

func (table *Message) MsgTableName() string {
	return "message"
}
