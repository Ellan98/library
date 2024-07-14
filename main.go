/*
 * @Date: 2024-06-28 10:33:07
 * @LastEditTime: 2024-07-14 19:32:41
 * @FilePath: \library_room\main.go
 * @description: 注释
 */
package main

import (
	"fmt"
	"library_room/internal/config"
	"library_room/internal/dao"
	"library_room/internal/db"
	"library_room/internal/entity"
	"library_room/server"
	"log"
)

func main() {
	conf, err := config.NewFromFile("configs/library_room.yml")
	if err != nil {
		// 发生错误时的处理逻辑，比如日志记录或返回错误信息给调用者
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("conf:%v\n", conf)

	DB, err := db.Newdb(conf.DB)
	if err != nil {
		// 发生错误时的处理逻辑，比如日志记录或返回错误信息给调用者
		log.Fatalf("\nFailed to load gormDB: %v", err)
	}
	dao := dao.Dao{
		DB: DB,
	}

	var list []*entity.SmartModel

	// Gemini := &entity.SmartModel{
	// 	Name:    "Gemini",
	// 	AlModel: "gemini-1.5-pro",
	// }

	// OpenAl := &entity.SmartModel{
	// 	Name:    "OpenAI",
	// 	AlModel: "gpt-3.5-turbo",
	// }

	// // 定义一个包含两个结构体的结构体类型
	// type Combined struct {
	// 	Struct1 entity.SmartModel
	// 	Struct2 entity.SmartModel
	// }

	// // 创建一个 Combined 类型的切片并初始化
	// slices := make([]Combined, 0)
	// slices = append(slices, Combined{Struct1: *Gemini, Struct2: *OpenAl})

	// // 遍历切片中的元素
	// for _, el := range slices {
	// 	dao.DB.Create(&el.Struct1)
	// 	dao.DB.Create(&el.Struct2)
	// }
	dao.DB.Find(&list)
	// // Print the actual data content of the list
	for _, user := range list {
		fmt.Printf("User: %+v\n", user)
	}

	server.StartServer(list)
	// app := cmd.New()
	// app.Launch()

}
