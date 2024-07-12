/*
 * @Date: 2024-07-04 13:40:07
 * @LastEditTime: 2024-07-12 10:26:37
 * @FilePath: \library_room\internal\config\base.go
 * @description: 注释
 */
package config

import (
	"fmt"
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func New() *Config {
	conf := &Config{}

	return conf
}

func NewFromFile(cfgFile string) (*Config, error) {
	// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
	kf := koanf.New(".")
	if err := kf.Load(file.Provider(cfgFile), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Printf("kf:%+v\n", kf.String("db.type"))
	conf := &Config{
		cfgFile: cfgFile,
	}

	// 将 koanf 实例中的值映射到 Config 结构体
	if err := kf.Unmarshal("", conf); err != nil {
		log.Fatalf("error unmarshalling config: %+v\n", err)
	}

	fmt.Printf("config file 当前实例化DBConf %v", conf)

	// fmt.Printf("当前结构体%+v\n", conf)

	return conf, nil
}
