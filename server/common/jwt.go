/*
 * @Date: 2024-07-17 16:26:11
 * @LastEditTime: 2024-09-13 10:55:29
 * @FilePath: \library_room\server\common\jwt.go
 * @description: 注释
 */
package common

import (
	"fmt"
	"library_room/internal/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

// 定义 JWT 的秘钥
var jwtKey = []byte("your_secret_key")

type jwtCustomClaims struct {
	UserID   uint64 `json:"user_id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func LoginGetUserToken(user *entity.UserBasic) (string, error) {
	//设置有效时间
	expirationTime := time.Now().Add(24 * time.Hour)
	// 为JWT 添加声明信息
	claims := &jwtCustomClaims{
		UserID:   user.UserID,
		Account:  user.Account,
		Password: user.Password,
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

func GetTokenByReq(tokenString string) (*jwt.Token, error) {
	// token := c.Get("token")

	// 使用相同的密钥来验证和解码
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		// 验证方法名
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtKey, nil

	})

	if err != nil {
		return nil, err
	}

	// fmt.Println("当前token", token.Claims)
	return token, nil

}
