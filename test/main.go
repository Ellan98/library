package main

import (
	"library_room/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=Hhu123456& dbname=library_room port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
	err = db.AutoMigrate(&entity.SmartModel{})
	if err != nil {
		panic(err)
	}

}
