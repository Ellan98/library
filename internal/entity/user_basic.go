/*
 * @Date: 2024-07-12 09:32:01
 * @LastEditTime: 2024-08-02 11:15:52
 * @FilePath: \library_room\internal\entity\user_basic.go
 * @description: 注释
 */
package entity

import (
	"time"

	"gorm.io/gorm"
)

// type Model struct {
// 	ID        uint           `gorm:"prinmaryKey`
// 	CreateAt  time.Time      `gorm:"create_at"`
// 	Update    time.Time      `gorm:"type:datetime"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Avatar        string
	Account       string
	Gender        string `gorm:"type:varchar(6); default:'male'; comment:'male表示男，female表示女'"` //gorm为数据库字段约束
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string `valid:"ipv4"`
	ClientPort    string
	Salt          string     //盐值
	LoginTime     *time.Time `gorm:"column:login_time"`
	HeartBeatTime *time.Time `gorm:"column:heart_beat_time"`
	LoginOutTime  *time.Time `gorm:"column:login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string //登录设备
}

func (table *UserBasic) UserTableName() string {
	return "user_basic"
}
