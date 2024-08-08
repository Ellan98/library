package dao

import (
	"fmt"
	"library_room/internal/entity"
)

func (dao *Dao) FindTaskList(task *entity.Task) {

	var taskList []entity.Task

	dao.DB().Model(&entity.Task{}).Where("user_id = ?", task.UserID).Find(&taskList)
	fmt.Println("implements...........", taskList)

		dao.DB().Model(&entity.Task{}).Create(&task)
	
}
