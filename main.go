/*
 * @Date: 2024-06-28 10:33:07
 * @LastEditTime: 2024-07-05 10:56:58
 * @FilePath: \library_room\main.go
 * @description: 注释
 */
package main

import (
	"library_room/cmd"
)

func main() {

	app := cmd.New()
	app.Launch()

}
