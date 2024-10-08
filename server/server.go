/*
 * @Date: 2024-06-28 10:35:01
 * @LastEditTime: 2024-08-21 14:51:32
 * @FilePath: \library_room\server\server.go
 * @description: 注释
 */
package server

import (
	"fmt"
	"library_room/internal/core"
	h "library_room/server/handler"
	"library_room/server/middleware"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name"`
}

func StartServer(app *core.App) (*fiber.App, error) {

	fmt.Println("app 相关配置", app.Conf())
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
	h.QueryTaskList(app, api)
	h.CreateTask(app, api)
	h.FindTaskById(app, api)
	h.EditTaskById(app, api)
	h.DeleteTask(app, api)
	h.SendMessage(app, api)
	fmt.Printf("%v", app)
	return fb, fb.Listen(":3030")

}

func cors(app *core.App, fiber fiber.Router) {
	fiber.Use(middleware.CorsMiddleware(app))

}

// func socket(app *core.App, fiber fiber.Router) error {
// 	fiber.Use(func(c *fiber.Ctx) error {
// 		// IsWebSocketUpgrade returns true if the client
// 		// requested upgrade to the WebSocket protocol.
// 		if websocket.IsWebSocketUpgrade(c) {
// 			c.Locals("allowed", true)
// 			return c.Next()
// 		}
// 		return fiber.ErrUpgradeRequired
// 	})
// }
