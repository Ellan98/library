/*
 * @Date: 2024-08-07 17:34:09
 * @LastEditTime: 2024-08-08 09:12:55
 * @FilePath: \library_room\server\handler\task_query_get.go
 * @description: 注释
 */
package handler

import (
	"fmt"
	"library_room/internal/core"

	"github.com/gofiber/fiber/v2"
)

type RequestTask struct {
	UserId string `json:"userId"` // 后期 是否 需要  换为 token ？
	Level  string `json:"level"`
}

func QueryTaskList(app *core.App, r fiber.Router) {

	r.Get("/task/list", func(c *fiber.Ctx) error {
		query := c.Queries()
		fmt.Println("fiber.Ctx的context", query)

		return nil
	})

}
