package dao

import (
	"fmt"
	"library_room/internal/entity"
)

func (dao *Dao) FindTaskList(task *entity.Task) {

	dao.DB().Model(&entity.Task{}).Create(&task)
}

func (dao *Dao) FindUserTaskList(userId string) []entity.Task {

	var taskList []entity.Task
	fmt.Println("id", userId)
	dao.DB().Model(&entity.Task{}).Where("user_id = ?", userId).Find(&taskList)

	return taskList

}
