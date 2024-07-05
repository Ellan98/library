/*
 * @Date: 2024-06-28 10:35:27
 * @LastEditTime: 2024-07-05 15:23:00
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
			Short: "Library: A self-hosted comment system",
			Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:
		
		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("base called")
			},
		},
	}
	cmd.RootCmd.SetVersionTemplate("Library ({{printf \"%s\" .Version}})\n")

	// Parse base flags
	cmd.eagerParseFlags()

	return cmd

}

func (library *LibraryRoomCmd) eagerParseFlags() error {
	library.RootCmd.PersistentFlags().StringVarP(&library.cfgFile, "config", "c", "", "config file path (defaults are 'configs/library.yml')")
	// library.RootCmd.PersistentFlags().StringVarP(&library.workDir, "workdir", "w", "", "program working directory (defaults are './')")
	return library.RootCmd.ParseFlags(os.Args[1:])
}

func (library *LibraryRoomCmd) addCommand(cmd *cobra.Command) {
	// originalPreRunFunc := cmd.PreRun

	// cmd.PreRun = func(cmd *cobra.Command, args []string) {
	// 	// ================
	// 	//  APP Bootstrap
	// 	// ================
	if cmd.Annotations[BootModeKey] != string(MODE_MINI_BOOT) {
		// Load config
		config, err := getConfig(library.cfgFile)
		if err != nil {
			log.Fatal("Config fail: ", err)
		}

		fmt.Printf("文件路径:%v\n", config)
		// 		// Create new instance
		library.App = core.NewApp(config)

	}

	// 	if originalPreRunFunc != nil {
	// 		originalPreRunFunc(cmd, args) // extends original pre-run logic
	// 	}
	// }
	library.RootCmd.AddCommand(cmd)
	fmt.Printf("library:type%+v\n", library)
}

func (library *LibraryRoomCmd) mountCommands() {

	library.addCommand(NewServeCommand(library))
	library.addCommand(NewConfigCommand())
}

// // 初始化配置
// cfg, err := getConfig("configs/library_room.yml")
// if err != nil {
// 	log.Fatalf("Error getting config: %v", err)
// }
// fmt.Printf("New implement....%+v\n", cfg)

// // 初始化数据库，传递 cfg.DB 给 db.Newdb 函数

// db, err := db.Newdb(cfg.DB)
// if err != nil {
// 	log.Fatalf("Error connecting to database: %v", err)
// }
// // configCmd := NewConfigCommand()
// // fmt.Printf("configCmd%+v\n", configCmd)
// server.StartServer()

// 	// Create new config instance by specific config filename
// func getConfig(cfgFile string) (*config.Config, error) {
// 	// Retrieve config file by default names when specific filename is empty

// 	// Create new config instance and return
// 	return config.NewFromFile(cfgFile)
// }

// Create new config instance by specific config filename
func getConfig(cfgFile string) (*config.Config, error) {
	// Retrieve config file by default names when specific filename is empty
	// if cfgFile == "" {
	// 	cfgFile = config.RetrieveConfigFile()
	// }

	// // Generate new config file when retrieve failed
	// if cfgFile == "" {
	// 	cfgFile = config.CONF_DEFAULT_FILENAMES[0]
	// 	core.Gen("config", cfgFile, false)
	// }

	// Create new config instance and return
	return config.NewFromFile(cfgFile)
}

func (library *LibraryRoomCmd) Launch() {
	// ===================
	//  1. Prepare Commands
	// ===================
	library.mountCommands()

	// 打印当前所有子命令
	library.printCommands()
	if err := library.RootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}

}

func (library *LibraryRoomCmd) printCommands() {
	commands := library.RootCmd.Commands()
	fmt.Println("Current commands:")
	for _, cmd := range commands {
		fmt.Printf("- %s: %s\n", cmd.Use, cmd.Short)
	}
}
