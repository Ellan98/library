/*
 * @Date: 2024-07-04 09:45:48
 * @LastEditTime: 2024-07-04 15:32:23
 * @FilePath: \library_room\internal\config\config.go
 * @description: 注释
 */
package config

type Config struct {
	DB      DBConf `koanf:"db" json:"db"` // 数据库配置
	cfgFile string
}

type DBConf struct {
	Type DBType `koanf:"type" json:"type"`
	Dsn  string `koanf:"dsn" json:"dsn"` // 最高优先级

	File string `koanf:"file" json:"file"`
	Name string `koanf:"name" json:"name"`

	Host     string `koanf:"host" json:"host"`
	Port     int    `koanf:"port" json:"port"`
	User     string `koanf:"user" json:"user"`
	Password string `koanf:"password" json:"password"`

	TablePrefix string `koanf:"table_prefix" json:"table_prefix"`
	Charset     string `koanf:"charset" json:"charset"`
	SSL         bool   `koanf:"ssl" json:"ssl"`
	PrepareStmt *bool  `koanf:"prepare_stmt" json:"prepare_stmt"`
}

type DBType string

const (
	TypeMySql      DBType = "mysql"
	TypeSQLite     DBType = "sqlite"
	TypePostgreSQL DBType = "pgsql"
	TypeMSSQL      DBType = "mssql"
)
