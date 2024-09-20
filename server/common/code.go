/*
 * @Date: 2024-09-20 16:48:10
 * @LastEditTime: 2024-09-20 17:14:34
 * @FilePath: \library_room\server\common\code.go
 * @description: 注释
 */
package common

const (
	// 10代表服务器相关错误
	ServerError = 10101
	// 请求次数 过多
	TooManyRequests = 10102
	// 参数绑定 错误
	ParamBindError = 10103
	// 授权失败
	AuthorizationError = 10104
	// Url 签名错误
	UrlSignError = 10105
	// 缓存设置错误
	CacheSetError = 10106
	// 缓存获取错误
	CacheGetError = 10107
	// 缓存删除错误
	CacheDelError = 10108
	// 缓存不存在
	CacheNotExist = 10109
	// 重新提交错误
	ResubmitError = 10110
	// 哈希Id编码错误
	HashIdsEncodeError = 10111
	// 哈希Id删除错误
	HashIdsDecodeError = 10112
	// RBAC 错误
	RBACError = 10113
	// redis连接错误
	RedisConnectError = 10114
	// mysql连接错误
	MySQLConnectError = 10115
	// 写入配置错误
	WriteConfigError = 10116
	// 发送邮箱错误
	SendEmailError = 10117
	// mysql执行错误
	MySQLExecError = 10118
	// go 版本错误
	GoVersionError = 10119
	// socket连接错误
	SocketConnectError = 10120
	// socket发送错误
	SocketSendError = 10121
)
