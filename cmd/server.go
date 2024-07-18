/*
 * @Date: 2024-07-05 10:00:24
 * @LastEditTime: 2024-07-16 17:36:20
 * @FilePath: \library_room\cmd\server.go
 * @description: 注释
 */
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"library_room/server"
	"log"

	"github.com/spf13/cobra"
)

func NewServeCommand(library *LibraryRoomCmd) *cobra.Command {
	// serverCmd represents the server command
	var serverCmd = &cobra.Command{
		Use:   "server",            //Use: "start" 指定了命令名称为 start。
		Short: "Start the service", // 提供了简短的描述。
		Long:  `run server`,
		//函数定义了执行命令时的具体逻辑。
		Run: func(cmd *cobra.Command, args []string) {
			server.StartServer(library.App)
			_, err := server.StartServer(library.App)
			if err != nil {
				log.Fatalln(err)
			}

		},
	}

	return serverCmd

}
