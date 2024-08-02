package handler

// handler 模块 编写接口

import (
	"fmt"
	"library_room/internal/core"
	"library_room/internal/entity"
	"library_room/server/common"

	"github.com/gofiber/fiber/v2"
)

type RequestAuthAccount struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"optional"`
}

var p RequestAuthAccount

// 登录
func AuthAccountLogin(app *core.App, router fiber.Router) {

	router.Post("/auth/login", func(c *fiber.Ctx) error {

		if ok, resp := common.ParamsDecode(c, &p); !ok {
			return resp
		}

		findAuth := func(account string, password string) (auth entity.UserBasic) {
			auths := app.Dao().FindAuthByAccount(account, password)

			if len(auths) == 0 {
				return entity.UserBasic{}
			}
			user := entity.UserBasic{}
			for _, auth := range auths {
				if auth.Password == password {
					user = auth
				}
			}
			return user
		}

		auth := findAuth(p.Account, p.Password)
		if auth.ID == 0 {
			mapAuth := map[string]interface{}{}
			return common.RespData(c, 500, fmt.Sprintf("用户%v不存在", p.Account), mapAuth)
		}

		// 构建返回的 map
		response := map[string]interface{}{
			"id":         auth.ID,
			"name":       auth.Name,
			"email":      auth.Email,
			"avatar":     auth.Avatar,
			"gender":     auth.Gender,
			"phone":      auth.Phone,
			"deviceInfo": auth.DeviceInfo,
			"createAt":   auth.CreatedAt,
			// 根据需要添加更多字段
		}
		return common.RespData(c, 200, "操作成功", response)
	})
}

// 注册
func AuthAccountSignup(app *core.App, router fiber.Router) {
	router.Post("/auth/signup", func(c *fiber.Ctx) error {
		if ok, resp := common.ParamsDecode(c, &p); !ok {
			return resp
		}
		user := &entity.UserBasic{
			Account:  p.Account,
			Password: p.Password,
		}

		err := app.Dao().CreateAuth(user)

		if err != nil {
			fmt.Errorf("错误", err)
			return common.RespData(c, 500, fmt.Sprintf("用户%v创建失败", p.Account), err)
		}

		mapData := map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"avatar":     user.Avatar,
			"gender":     user.Gender,
			"phone":      user.Phone,
			"deviceInfo": user.DeviceInfo,
			"createAt":   user.CreatedAt,
		}
		fmt.Println("解析后的参数为：", p)
		return common.RespData(c, 200, fmt.Sprintf("用户%v创建成功", p.Account), mapData)
	})
}

// 编辑
func AuthAccountEdit(app *core.App, router fiber.Router) {
	router.Post("/auth/edit", func(c *fiber.Ctx) error {
		if ok, resp := common.ParamsDecode(c, &p); !ok {
			return resp
		}
		user := &entity.UserBasic{
			Account:  p.Account,
			Password: p.Password,
		}
		err := app.Dao().CreateAuth(user)

		if err != nil {
			return common.RespData(c, 500, fmt.Sprintf("用户%v创建失败", p.Account), err)
		}

		mapData := map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"avatar":     user.Avatar,
			"gender":     user.Gender,
			"phone":      user.Phone,
			"deviceInfo": user.DeviceInfo,
			"createAt":   user.CreatedAt,
		}
		fmt.Println("解析后的参数为：", p)
		return common.RespData(c, 200, fmt.Sprintf("用户%v创建成功", p.Account), mapData)
	})
}

func getAuthsList() {

}
