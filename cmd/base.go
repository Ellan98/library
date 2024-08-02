/*
 * @Date: 2024-06-28 10:35:27
 * @LastEditTime: 2024-08-02 11:01:28
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
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
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
				fmt.Println("NOTE: add `-h` flag to show help about any command.")
			},
		},
	}
	cmd.RootCmd.SetVersionTemplate("Library ({{printf \"%s\" .Version}})\n")
	// Parse base flags
	cmd.eagerParseFlags()

	return cmd

}

func (library *LibraryRoomCmd) eagerParseFlags() error {
	// library.RootCmd.PersistentFlags()：访问命令的全局标志（persistent flags）。 这些标志是持久标志（persistent flags），它们对命令及其所有子命令都有效
	// StringVarP：定义一个字符串类型的标志
	//&library.cfgFile：将标志的值存储在 library 实例的 cfgFile 字段中。
	// "config"：标志的全名。
	// "c"：标志的短名（单字符）""：标志的默认值。 "config file path (defaults are 'configs/library.yml')"：标志的描述。
	library.RootCmd.PersistentFlags().StringVarP(&library.cfgFile, "config", "c", "", "config file path (defaults are 'configs/library_room.yml')")
	library.RootCmd.PersistentFlags().StringVarP(&library.workDir, "workdir", "w", "", "program working directory (defaults are './')")
	// 从命令行参数中解析标志。os.Args 是一个包含所有命令行参数的字符串数组，os.Args[0] 通常是程序名称，os.Args[1:] 包含实际的命令行参数。ParseFlags 方法会解析这些参数并将结果存储在相关的标志变量中。
	return library.RootCmd.ParseFlags(os.Args[1:])
}

func (library *LibraryRoomCmd) addCommand(cmd *cobra.Command) {
	// Load config
	// config, err := getConfig(library.cfgFile)

	library.RootCmd.Execute()
	config, err := getConfig("configs/library_room.yml")
	if err != nil {
		log.Fatal("Config fail: ", err)
	}
	// Create new instance
	library.App = core.NewApp(config)

	library.RootCmd.AddCommand(cmd)
	//运行子命令
	library.RootCmd.SetArgs([]string{"server", "config"})
}

func (library *LibraryRoomCmd) mountCommands() {
	library.addCommand(NewServeCommand(library))
	library.addCommand(NewConfigCommand())
}

// Create new config instance by specific config filename
func getConfig(cfgFile string) (*config.Config, error) {

	// Create new config instance and return
	return config.NewFromFile(cfgFile)
}

func (library *LibraryRoomCmd) Launch() error {
	// ===================
	//  1. Prepare Commands
	// ===================
	library.mountCommands()

	done := make(chan bool, 1) // shutdown signal

	// listen for interrupt signal to gracefully shutdown the application
	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		<-sigch

		done <- true
	}()

	// ===================
	//  2. Command Execute
	// ===================
	go func() {
		if err := library.RootCmd.Execute(); err != nil {
			color.Red(err.Error())
		}

		done <- true
	}()

	<-done

	// ===================
	//  3. App Cleanups
	// ===================
	// if library.App != nil {
	// 	return library.App.OnTerminate().Trigger(&core.TerminateEvent{
	// 		App: library.App,
	// 	})
	// }

	return nil

}

func (library *LibraryRoomCmd) printCommands() {
	commands := library.RootCmd.Commands()
	for _, cmd := range commands {
		fmt.Printf("- %s: %s\n", cmd.Use, cmd.Short)
	}
}
