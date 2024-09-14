/*
 * @Date: 2024-08-08 09:13:53
 * @LastEditTime: 2024-08-28 09:29:59
 * @FilePath: \library_room\server\handler\task_create_post.go
 * @description: 注释
 */
package handler

import (
	"library_room/internal/core"
	"library_room/internal/entity"
	"library_room/server/common"
	"time"

	"github.com/gofiber/fiber/v2"
)

// type RequestTaskCrate struct {
// 	entity.Task
// }

func CreateTask(app *core.App, r fiber.Router) {
	var p entity.Task
	r.Post("/task/create", func(c *fiber.Ctx) error {

		ok, resp := common.ParamsDecode(c, &p)
		if !ok {
			return resp
		}

		p.CreateTime = time.Now().Format("2006/01/02 15:04:05")
		p.UserID = 526064888234115472

		app.Dao().FindTaskList(&p)

		return nil

	})
}

// func ValidateParams(p entity.Task, c *fiber.Ctx) {
// 	// query := c.Queries()
// 	// for i := range query {
// 	// 	fmt.Printf("当前key%v,value为%v\n", i, query[i])
// 	// }
// 	if err := c.BodyParser(p); err != nil {
// 		fmt.Println("打印", p)
// 	}

// }
