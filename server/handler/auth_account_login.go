package handler

// handler 模块 编写接口

import (
	"library_room/internal/core"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name"`
}
type RequestAuthAccountLogin struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"optional"`
	Code     string `json:"code" validate:"optional"`
}

func AuthAccountLogin(app *core.App, router fiber.Router) {
	router.Post("/auth/account/login", func(c *fiber.Ctx) error {

		p := Person{Name: "hello"}
		return c.JSON(p)
		// return c.Send(p)
	})
}
