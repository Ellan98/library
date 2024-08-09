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
	dao.DB().Model(&entity.Task{}).Where("user_id = ?", userId).Find(&taskList)

	return taskList

}

// 根据task 的 id 进行查询
func (dao *Dao) FindTaskById(id string) []entity.Task {
	var list []entity.Task
	dao.DB().Model(&entity.Task{}).Where("id = ?", id).First(&list)
	return list

}

// 编辑 任务
func (dao *Dao) EditTaskById(p *entity.Task) {
	dao.DB().Model(&entity.Task{}).Where("id = ?", p.ID).Updates(&p)
}

// 删除
func (dao *Dao) DeleteTask(taskId string) {

	dao.DB().Model(&entity.Task{}).Where("id = ?", taskId).Delete(&entity.Task{})

	fmt.Println(taskId)

}
