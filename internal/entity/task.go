package entity

type Task struct {
	CreateTime string `json:"create_time" gorm:"comment:'创建时间'"`
	StartTime  string `json:"start_time" gorm:"comment:'开始日期'"`
	EndTime    string `json:"end_time" gorm:"comment:'结束日期'"`
	Title      string `json:"title" gorm:"comment:'任务标题'"`
	Tag        string `gorm:"comment:'任务标签，用于分类'"`
	Descript   string `json:descript gorm:"comment:'任务主体内容 简要描述'"`
	Details    string `json:details gorm:"comment:'任务主体内容 详细描述'"`
	Level      int64  `json:"level"gorm:"comment:'任务等级 0红色，1橙色，2黄色，3蓝色 '"`
	UserID     uint64 `json:"level"gorm:"comment:user_id 雪花算法'`
}
