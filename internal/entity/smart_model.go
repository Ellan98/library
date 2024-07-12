/*
 * @Date: 2024-07-12 14:09:14
 * @LastEditTime: 2024-07-12 14:11:13
 * @FilePath: \library_room\internal\entity\smart_model.go
 * @description: 注释
 */
package entity

import "gorm.io/gorm"

type SmartModel struct {
	gorm.Model
	Name    string
	ApiKey  string
	AlModel string
}

func (table *SmartModel) ModelTableName() string {
	return "smart_model"
}
