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

// 返回错误
func RespError(c *fiber.Ctx, code int, msg string, data ...Map) error {

	respData := Map{}
	if len(data) > 0 {
		respData = data[0]
	}

	respData["msg"] = msg
	respData["code"] = code
	c.Status(code)
	return c.JSON(respData)

}
