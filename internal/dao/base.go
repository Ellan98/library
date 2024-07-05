package dao

import "gorm.io/gorm"

type Dao struct {
	*gorm.DB
}
