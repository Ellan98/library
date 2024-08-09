package handler

import (
	"library_room/internal/core"
	"library_room/internal/entity"
	"library_room/server/common"

	"github.com/gofiber/fiber/v2"
)

func FindTaskById(app *core.App, r fiber.Router) {
	r.Get("/task/:id", func(c *fiber.Ctx) error {
		list := app.Dao().FindTaskById(c.Params("id"))
		common.RespData(c, 200, "查询成功", list)
		return nil
	})
}

func EditTaskById(app *core.App, r fiber.Router) {
	var p entity.Task
	r.Put("/task/edit", func(c *fiber.Ctx) error {

		c.BodyParser(&p)
		app.Dao().EditTaskById(&p)
		common.RespData(c, 200, "当前参数", p)

		return nil
	})
}

func DeleteTask(app *core.App, r fiber.Router) {

	r.Delete("/task/delete/:id", func(c *fiber.Ctx) error {
		app.Dao().DeleteTask(c.Params("id"))
		common.RespData(c, 200, "删除成功", []interface{}{})
		return nil
	})

}
