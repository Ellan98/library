/*
 * @Date: 2024-07-18 11:38:31
 * @LastEditTime: 2024-08-08 10:54:23
 * @FilePath: \library_room\internal\dao\query_auth.go
 * @description: 注释 定义对数据库操作的方法
 */
package dao

import (
	"errors"
	"library_room/internal/entity"
)

// 查找用户 (通过 account)
func (dao *Dao) FindAuthByAccount(account string, password string) []entity.UserBasic {

	var users []entity.UserBasic
	dao.DB().Model(&entity.UserBasic{}).Where("account = ? ", account).Find(&users)

	return users
}

//创建用户

func (dao *Dao) CreateAuth(auth *entity.UserBasic) error {
	var users []entity.UserBasic

	dao.DB().Model(&entity.UserBasic{}).Where("account = ? ", auth.Account).Find(&users)

	if len(users) != 0 {
		return errors.New("用户已经存在")
	}
	dao.DB().Model(&entity.UserBasic{}).Create(&auth)
	return nil

}
