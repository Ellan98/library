/*
 * @Date: 2024-07-12 14:09:14
 * @LastEditTime: 2024-07-13 12:03:44
 * @FilePath: \library_room\internal\entity\smart_model.go
 * @description: 注释
 */
package entity

import "gorm.io/gorm"

type SmartModel struct {
	gorm.Model
	Name    string `json:"name"`
	ApiKey  string `json:"api_key"`
	AlModel string `json:"ai_model"`
}

func (table *SmartModel) ModelTableName() string {
	return "smart_model"
}
