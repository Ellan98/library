package handler

// handler 模块 编写接口

import (
	"library_room/internal/core"
	"library_room/server/common"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name     string `json:"account"`
	Password string `json`
}
type RequestAuthAccountLogin struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"optional"`
	Code     string `json:"code" validate:"optional"`
}

func AuthAccountLogin(app *core.App, router fiber.Router) {

	router.Post("/auth/account/login", func(c *fiber.Ctx) error {

		var p RequestAuthAccountLogin

		if ok, resp := common.ParamsDecode(c, p); !ok {
			return resp
		}
		// fmt.Print()

		data, error := app.Dao().FindAuthByAccount(p.Account, p.Password)
		if error != nil {

		}

		return c.JSON(data)

		// p := Person{Name: "hello"}

		// return c.Send(p)
	})
}
