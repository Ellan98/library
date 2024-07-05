/*
 * @Date: 2024-06-28 10:35:01
 * @LastEditTime: 2024-07-05 10:07:38
 * @FilePath: \library_room\server\server.go
 * @description: 注释
 */
package server

import (
	"fmt"
	"library_room/internal/core"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name"`
}

func StartServer(app *core.App) {
	fmt.Printf("执行app:%+v", app)
	fb := fiber.New(fiber.Config{
		// @see https://github.com/gofiber/fiber/issues/426
		// @see https://github.com/gofiber/fiber/issues/185
		Immutable: true,

		// ErrorHandler:       common.ErrorHandler,
		// ProxyHeader:        app.Conf().HTTP.GetProxyHeader(),
		// BodyLimit:          app.Conf().HTTP.BodyLimit * 1024 * 1024, // MB
		StreamRequestBody:  true,
		EnableIPValidation: true,
	})
	p := Person{Name: "hello"}
	fb.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(p)
		// return c.Send(p)
	})

	fb.Listen(":3000")
}
