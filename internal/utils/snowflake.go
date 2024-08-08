/*
 * @Date: 2024-08-06 17:19:02
 * @LastEditTime: 2024-08-06 17:43:47
 * @FilePath: \library_room\internal\utils\snowflake.go
 * @description: 注释
 */
package utils

import (
	"fmt"

	"github.com/sony/sonyflake"
)

func NewSonyflake() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		fmt.Printf("生成雪花 ID 时出错: %s\n", err)
	}
	return id
}
