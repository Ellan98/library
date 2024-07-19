package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type JSONResult struct {
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// RespData just response data
func RespData(c *fiber.Ctx, data interface{}) error {
	// 设置响应的 HTTP 状态。
	return c.Status(http.StatusOK).JSON(data)
}

func RespSuccess(c *fiber.Ctx, msg ...string) error {
	respData := &JSONResult{
		Msg: "Success",
	}

	if len(msg) > 0 {
		respData.Msg = msg[0]
	}
	return c.Status(http.StatusOK).JSON(respData)
}
