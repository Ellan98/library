/*
 * @Date: 2024-07-04 14:20:25
 * @LastEditTime: 2024-07-16 16:53:04
 * @FilePath: \library_room\cmd\config.go
 * @description: 注释
 */
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "implement Config Information",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("当前文件名run 因为没有赋值 所以为空")
			// Get config filename from cmd flags
			filename, err := cmd.Flags().GetString("config")
			if err != nil {
				log.Fatal("Config  错误: ", err)
			}
			fmt.Printf("当前文件名%v\n 因为没有赋值 所以为空", filename)
			// Get new config instance
			// config, err := getConfig(filename)
			// if err != nil {
			// 	log.Fatal("Config fail: ", err)
			// }

			// Output JSON of config
			// Go 语言标准库 encoding/json 中的一个函数，数据结构（如结构体、数组、切片、映射等）序列化为格式化的 JSON 字符串。与 json.Marshal 不同的是，json.MarshalIndent 会对输出的 JSON 进行缩进和格式化，使其更易读。
			// buf, _ := json.MarshalIndent(config, "", "    ")
		},
	}

	return configCmd
}
