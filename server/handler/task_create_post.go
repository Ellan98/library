package handler

import (
	"fmt"
	"library_room/internal/core"
	"library_room/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

// type RequestTaskCrate struct {
// 	entity.Task
// }

func CreateTask(app *core.App, r fiber.Router) {
	var p entity.Task
	r.Post("/task/create", func(c *fiber.Ctx) error {
		fmt.Println("打印")
		// common.ParamsDecode(c, &p)
		// ValidateParams(&p, c)
		c.BodyParser(&p)
		// task := &entity.Task{

		// }
		p.CreateTime = time.Now().Format("2006/01/02 15:04:05")
		p.UserID = 526064888234115472

		app.Dao().FindTaskList(&p)

		return nil

	})
}

func ValidateParams(p entity.Task, c *fiber.Ctx) {
	// query := c.Queries()
	// for i := range query {
	// 	fmt.Printf("当前key%v,value为%v\n", i, query[i])
	// }
	if err := c.BodyParser(p); err != nil {
		fmt.Println("打印", p)
	}

	// fmt.Println("test", p.Level)

}
