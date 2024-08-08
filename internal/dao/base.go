/*
 * @Date: 2024-07-04 16:41:45
 * @LastEditTime: 2024-08-07 15:15:41
 * @FilePath: \library_room\internal\dao\base.go
 * @description: 注释
 */
package dao

import (
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{
		db: db,
	}

	return dao
}

func (dao *Dao) DB() *gorm.DB {
	return dao.db
}
