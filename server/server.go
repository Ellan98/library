/*
 * @Date: 2024-06-28 10:35:01
 * @LastEditTime: 2024-07-16 16:24:11
 * @FilePath: \library_room\server\server.go
 * @description: 注释
 */
package server

import (
		"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"library_room/internal/core"
)

type Person struct {
	Name string `json:"name"`
}


func StartServer(app *core.App) (*fiber.App, error) {
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
	fb.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	p := Person{Name: "hello"}
	fb.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(p)
		// return c.Send(p)
	})
	fb.Get("/modellist", func(c *fiber.Ctx) error {
		return c.JSON(p)
		// return c.Send(p)
	})

	fmt.Printf("%v",app)
	return fb, fb.Listen(":3000")
}
