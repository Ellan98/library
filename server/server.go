/*
 * @Date: 2024-06-28 10:35:01
 * @LastEditTime: 2024-08-02 14:44:43
 * @FilePath: \library_room\server\server.go
 * @description: 注释
 */
package server

import (
	"fmt"
	"library_room/internal/core"
	h "library_room/server/handler"
	"library_room/server/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name"`
}

func StartServer(app *core.App) (*fiber.App, error) {
	loc, _ := time.LoadLocation("Local")
	time.Local = loc
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

	fb.Get("/modellist", func(c *fiber.Ctx) error {
		return c.JSON(p)
	})

	api := fb.Group("/api/v2")

	h.AuthAccountLogin(app, api)
	h.AuthAccountSignup(app, api)
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(p)
		// return c.Send(p)
	})
	fmt.Printf("%v", app)
	return fb, fb.Listen(":3030")

}

func cors(app *core.App, fiber fiber.Router) {
	fiber.Use(middleware.CorsMiddleware(app))
}
