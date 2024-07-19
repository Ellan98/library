/*
 * @Date: 2024-07-18 11:38:31
 * @LastEditTime: 2024-07-19 11:06:10
 * @FilePath: \library_room\internal\dao\query_auth.go
 * @description: 注释 定义对数据库操作的方法
 */
package dao

import (
	"fmt"
	"library_room/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 查找用户 (通过 account)
func (dao *Dao) FindAuthByAccount(account string, password string) (*entity.UserBasic, error) {
	dsn := "host=localhost user=postgres password=Hhu123456& dbname=library_room port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var user entity.UserBasic
	// Get first matched record
	result := db.Where("account = ?  AND password >= ?", account, password).Find(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Printf("未找到用户：%s\n", account)
			return nil, result.Error
		}
		fmt.Printf("查找用户时发生错误：%v\n", result.Error)
		return nil, result.Error
	}

	fmt.Printf("当前查找的用户为：%v\n", user)
	return &user, nil
}
