/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"library_room/server"

	"github.com/spf13/cobra"
)

func NewServeCommand(library *LibraryRoomCmd) *cobra.Command {
	// serverCmd represents the server command
	var serverCmd = &cobra.Command{
		Use:   "server",            //Use: "start" 指定了命令名称为 start。
		Short: "Start the service", // 提供了简短的描述。
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		//函数定义了执行命令时的具体逻辑。
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server called")
			server.StartServer(library.App)

			// _, err := server.StartServer(library.App)
			// if err != nil {
			// 	log.Fatalln(err)
			// }

		},
	}

	return serverCmd

}
