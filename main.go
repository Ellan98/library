/*
 * @Date: 2024-06-28 10:33:07
 * @LastEditTime: 2024-07-12 14:45:36
 * @FilePath: \library_room\main.go
 * @description: 注释
 */
package main

import (
	"fmt"
	"library_room/internal/config"
	"library_room/internal/dao"
	"library_room/internal/db"
	"library_room/internal/models"
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

	var list []*models.UserBasic

	dao.DB.Find(&list)
	// Print the actual data content of the list
	for _, user := range list {
		fmt.Printf("User: %+v\n", user)
	}

	server.StartServer()
	// app := cmd.New()
	// app.Launch()

}
