package handler

// handler 模块 编写接口

import (
	"fmt"
	"library_room/internal/core"
	"library_room/internal/entity"
	"library_room/internal/utils"
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

			fmt.Println("待匹配账户", account)
			auths := app.Dao().FindAuthByAccount(account)

			if len(auths) == 0 {
				fmt.Println("未匹配到用户", auths)
				return entity.UserBasic{}
			}
			user := entity.UserBasic{}
			for _, auth := range auths {
				if auth.Password == password {
					user = auth
				}
			}
			fmt.Println("匹配到用户", user)
			return user
		}

		auth := findAuth(p.Account, p.Password)

		if auth.ID != 0 {
			token, err := common.LoginGetUserToken(&auth)
			if err != nil {
				return common.RespError(c, 500, "token生成失败")
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
				"userId":     auth.UserID,
				// 根据需要添加更多字段
				"token": token,
			}
			common.GetTokenByReq(token)
			return common.RespData(c, 200, "登录成功", response)
		} else {
			return common.RespError(c, 500, "账号或密码错误", map[string]interface{}{})
		}

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
			Name:     p.Account,
			Password: p.Password,
			UserID:   utils.NewSonyflake(),
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
			"userID":     user.UserID,
		}
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
		return common.RespData(c, 200, fmt.Sprintf("用户%v创建成功", p.Account), mapData)
	})
}
