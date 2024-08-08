/*
 * @Date: 2024-06-28 10:35:27
 * @LastEditTime: 2024-08-07 13:55:48
 * @FilePath: \library_room\cmd\base.go
 * @description: 注释
 */
/*
 * @Date: 2024-06-28 10:35:27
 * @LastEditTime: 2024-07-16 17:37:18
 * @FilePath: \library_room\cmd\base.go
 * @description: 注释
 */
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"library_room/internal/config"
	"library_room/internal/core"
	"library_room/server"

	"github.com/spf13/cobra"
)

type LibraryRoomCmd struct {
	*core.App
	cfgFile string
	workDir string
	RootCmd *cobra.Command
}

const BootModeKey = "BootMode"

const (
	MODE_FULL_BOOT = "FULL_BOOT"
	MODE_MINI_BOOT = "MINI_BOOT"
)

func New() *LibraryRoomCmd {

	cmd := &LibraryRoomCmd{
		RootCmd: &cobra.Command{
			Use:   "library",
			Short: "Library: 为主commond",
			Long: `Cobra是Go的CLI库，可为应用程序提供功能。
此应用程序是生成所需文件的工具
以快速创建Cobra应用程序。`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Print("-------------------------------\n\n")

			},
		},
	}

	return cmd

}

// Create new config instance by specific config filename
func getConfig(cfgFile string) (*config.Config, error) {

	// Create new config instance and return
	return config.NewFromFile(cfgFile)
}

func (library *LibraryRoomCmd) Launch() error {
	library.cfgFile = "configs/library_room.yml"

	// Get new config instance
	config, _ := getConfig(library.cfgFile)

	library.App = core.NewApp(config)
	_, err := server.StartServer(library.App)

	if err != nil {
		return err
	}
	return nil

}

func (library *LibraryRoomCmd) printCommands() {
	commands := library.RootCmd.Commands()
	for _, cmd := range commands {
		fmt.Printf("- %s: %s\n", cmd.Use, cmd.Short)
	}
}
