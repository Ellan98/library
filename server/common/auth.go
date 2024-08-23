/*
 * @Date: 2024-07-17 16:26:11
 * @LastEditTime: 2024-08-19 09:49:13
 * @FilePath: \library_room\server\common\auth.go
 * @description: 注释
 */
package common

import (
	"library_room/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// 定义 JWT 的秘钥
var jwtKey = []byte("your_secret_key")

type jwtCustomClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

func LoginGetUserToken(user *entity.UserBasic) (string, error) {
	//设置有效时间
	expirationTime := time.Now().Add(24 * time.Hour)
	// 为JWT 添加声明信息
	claims := &jwtCustomClaims{
		UserID: user.UserID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),     // 签发时间
			ExpiresAt: expirationTime.Unix(), // 过期时间
		},
	}
	//创建jwt 并 签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetTokenByReq(c *fiber.Ctx) {
	// token := c.Get("token")

	// ParseWithClaims

}
