/*
 * @Date: 2024-06-28 10:41:02
 * @LastEditTime: 2024-07-12 11:39:50
 * @FilePath: \library_room\internal\db\db.go
 * @description: 注释
 */
/*
 * @Date: 2024-06-28 10:41:02
 * @LastEditTime: 2024-06-28 10:52:50
 * @FilePath: \project-layout\internal\db\base.go
 * @description: 注释
 */
package db

import (
	"fmt"
	"library_room/internal/config"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Newdb(conf config.DBConf) (*gorm.DB, error) {
	// fmt.Printf("当前键值%v", conf)
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: conf.TablePrefix,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	// Enable Prepared Statement by default
	// PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，
	if prepareStmt := conf.PrepareStmt; prepareStmt != nil {
		gormConfig.PrepareStmt = *prepareStmt
	} else {
		gormConfig.PrepareStmt = true
	}

	var dsn string
	if conf.Dsn != "" {
		dsn = conf.Dsn
	} else {
		dsn = getDsnByConf(conf)
	}

	switch conf.Type {
	case config.TypeSQLite:
		return OpenSQLite(dsn, gormConfig)
	case config.TypeMySql:
		return OpenMySql(dsn, gormConfig)
	case config.TypePostgreSQL:
		return OpenPostgreSQL(dsn, gormConfig)

		// return db
	case config.TypeMSSQL:
		return OpenSqlServer(dsn, gormConfig)
	}
	return nil, fmt.Errorf("unsupported database type %s", conf.Type)

}

func CloseDB(db *gorm.DB) error {
	ddb, err := db.DB()
	if err != nil {
		return err
	}
	return ddb.Close()
}
