package handler

import (
	"fmt"
	"library_room/internal/core"
	"library_room/internal/models"

	"github.com/gofiber/fiber/v2"
)

type RequestMsgInfo struct {
	Message string `json:"messgae"`
}

func SendMessage(app *core.App, r fiber.Router) {

	var m models.Message
	r.Get("/spark/conversation", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		if err := c.QueryParser(&m); err != nil {
			return err
		}

		fmt.Println("当前映射信息", m)
		createSomething(c, m.Content)

		// data := models.RequestModelQuestion(m.Content)
		// common.RespData(c, 200, "操作成功", data)
		return nil
	})

}

// Post something
func createSomething(c *fiber.Ctx, content string) (err error) {

	// Authorization: Bearer 123456

	data := make([]map[string]string, 1)

	data = append(data, map[string]string{
		"role":    "user",
		"content": content,
	})

	agent := fiber.Post("https://spark-api-open.xf-yun.com/v1/chat/completions")
	agent.ContentType("application/json")
	// CobzANrvxDAcxZIgYamO:orIzGgUFSOZybMOdOwpl   4.0Ultra
	agent.Set("Authorization", "Bearer gsLHbjIbIclcQnWGDMPd:SLIHjpgTaKoFzNslIjKF")
	agent.Body(c.Body()) // set body received by request
	agent.JSON(fiber.Map{
		"model":    "Lite",
		"messages": data,
		"stream":   true,
	})
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	// pass status code and body received by the proxy
	return c.Status(statusCode).Send(body)
}
