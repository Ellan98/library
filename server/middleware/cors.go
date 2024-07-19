/*
 * @Date: 2024-07-12 15:11:30
 * @LastEditTime: 2024-07-19 08:27:32
 * @FilePath: \library_room\server\middleware\cors.go
 * @description: 注释
 */
package middleware

import (
	"fmt"
	"library_room/internal/core"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getCorsAllowOrigins(app *core.App) []string {
	allowURLs := []string{}
	// allowURLs = append(allowURLs, app.Conf().TruestedDomains...) //导入配置中可信域名
	return allowURLs
}

func CheckOriginTrusted(app *core.App, origin string) bool {
	// 用于检查切片中是否包含特定元素。 返回布尔值 slices.Contains(getCorsAllowOrigins(app), origin)
	return true
}

// fb.Use(cors.New(cors.Config{
// 	AllowOrigins: "*",
// 	AllowHeaders: "Origin, Content-Type, Accept",
// }))

func CorsMiddleware(app *core.App) func(*fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {

			fmt.Printf("app.Conf的配置为%v\n", app.Conf())
			return CheckOriginTrusted(app, origin)
		},
	})
}
