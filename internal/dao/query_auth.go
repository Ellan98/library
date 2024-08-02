/*
 * @Date: 2024-07-18 11:38:31
 * @LastEditTime: 2024-08-02 14:44:07
 * @FilePath: \library_room\internal\dao\query_auth.go
 * @description: 注释 定义对数据库操作的方法
 */
package dao

import (
	"errors"
	"library_room/internal/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 查找用户 (通过 account)
func (dao *Dao) FindAuthByAccount(account string, password string) []entity.UserBasic {
	// dsn := "host=localhost user=postgres password=Hhu123456& dbname=library_room port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dsn := "root:xu20010809@tcp(127.0.0.1:3306)/library_room?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// sqlDB, err := sql.Open("mysql", "mydb_dsn")
	// gormDB, err := gorm.Open(mysql.New(mysql.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var users []entity.UserBasic
	// Get matched records  password
	db.Where("account = ? ", account).Find(&users)

	return users
}

//创建用户

func (dao *Dao) CreateAuth(auth *entity.UserBasic) error {

	// dsn := "host=localhost user=postgres password=Hhu123456& dbname=library_room port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dsn := "root:xu20010809@tcp(127.0.0.1:3306)/library_room?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var users []entity.UserBasic

	db.Where("account = ?", auth.Account).Find(&users)

	if len(users) != 0 {
		return errors.New("用户已经存在")
	}
	db.Create(&auth)

	return nil

}
