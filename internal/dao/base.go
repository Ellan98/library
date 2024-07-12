/*
 * @Date: 2024-07-04 16:41:45
 * @LastEditTime: 2024-07-12 11:01:49
 * @FilePath: \library_room\internal\dao\base.go
 * @description: 注释
 */
package dao

import (
	"gorm.io/gorm"
)

type Dao struct {
	DB *gorm.DB
}
