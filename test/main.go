package main

import (
	"database/sql"
	"library_room/internal/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// dsn := "host=localhost user=postgres password=Hhu123456& dbname=library_room port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "root:root@tcp(127.0.0.1:3306)/library_room?charset=utf8mb4&parseTime=True&loc=Local"

	sqlDB, err := sql.Open("mysql", dsn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
	err = db.AutoMigrate(&entity.UserBasic{})
	if err != nil {
		panic(err)
	}

}
