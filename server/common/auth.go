/*
 * @Date: 2024-07-17 16:26:11
 * @LastEditTime: 2024-08-08 08:33:17
 * @FilePath: \library_room\server\common\auth.go
 * @description: 注释
 */
package common

// type jwtCustomClaims struct {
// 	UserID uint64 `json:"user_id"`
// 	jwt.StandardClaims
// }

// func LoginGetUserToken(user entity.UserBasic, ttl int) (string, error) {
// 	// 为JWT 添加声明信息
// 	// claims := &jwtCustomClaims{
// 	// 	UserID: user.UserID,
// 	// 	StandardClaims: jwt.StandardClaims{
// 	// 		IssuedAt:  time.Now().Unix(),                                       // 签发时间
// 	// 		ExpiresAt: time.Now().Add(time.Second * time.Duration(ttl)).Unix(), // 过期时间
// 	// 	},
// 	// }

// 	// claims := jwt.MapClaims{
// 	// 	"iss":""
// 	// }
// }
