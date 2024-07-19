/*
 * @Date: 2024-07-19 08:38:50
 * @LastEditTime: 2024-07-19 14:28:31
 * @FilePath: \library_room\server\common\params.go
 * @description: 注释
 */
package common

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Map = map[string]interface{}

// 参数解码
func ParamsDecode(c *fiber.Ctx, destParams interface{}) (isContinue bool, resp error) {
	reqMethod := c.Method()
	fmt.Printf("当前请求方法:%v", reqMethod)
	// parse body
	if reqMethod == "POST" || reqMethod == "PUT" {
		if err := c.BodyParser(destParams); err != nil {

			fmt.Printf("当前解析:%v", destParams)
			return false, err
		}
	}
	return true, nil
}
