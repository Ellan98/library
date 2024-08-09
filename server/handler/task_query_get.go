/*
 * @Date: 2024-08-07 17:34:09
 * @LastEditTime: 2024-08-09 08:46:55
 * @FilePath: \library_room\server\handler\task_query_get.go
 * @description: 注释
 */
package handler

import (
	"fmt"
	"library_room/internal/core"
	"library_room/server/common"

	"github.com/gofiber/fiber/v2"
)

type RequestTask struct {
	UserId string `json:"userId"` // 后期 是否 需要  换为 token ？
	Level  string `json:"level"`
}

func QueryTaskList(app *core.App, r fiber.Router) {

	r.Get("/task/list", func(c *fiber.Ctx) error {
		//分配内存 并进行初始化 返回一个指向分配内存的指针，指针的类型为 存储在变量 p 中：变量 p 将持有这个指向新分配的 RequestTask 实例的指针。
		p := new(RequestTask)
		c.QueryParser(p)

		list := app.Dao().FindUserTaskList(p.UserId)
		fmt.Printf("查询数据%v\n,当前反射数据%v\n", list, p.UserId)
		common.RespData(c, 200, "操作成功", list)

		return nil
	})

}
