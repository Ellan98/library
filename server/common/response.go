package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type JSONResult struct {
	Msg  string      `json:"msg,omitempty"`
	Code int32       `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// RespData just response data
func RespData(c *fiber.Ctx, code int32, msg string, data interface{}) error {
	respData := &JSONResult{
		Msg:  msg,
		Code: code,
		Data: data,
	}
	// 设置响应的 HTTP 状态。
	return c.Status(http.StatusOK).JSON(respData)
}
