/*
 * @Date: 2024-07-04 11:23:35
 * @LastEditTime: 2024-08-16 08:59:40
 * @FilePath: \library_room\internal\core\base.go
 * @description: 注释
 */
package core

import (
	"fmt"
	"library_room/internal/config"
	"library_room/internal/dao"
	"library_room/internal/db"

	"github.com/patrickmn/go-cache"
)

type App struct {
	conf    *config.Config
	dao     *dao.Dao
	cache   *cache.Cache
	service *map[string]Service

	// onTerminate   *hook.Hook[*TerminateEvent]
	// onConfUpdated *hook.Hook[*ConfUpdatedEvent]
}

func NewApp(conf *config.Config) *App {

	app := &App{
		conf:    conf,
		service: &map[string]Service{},
		// onTerminate:   &hook.Hook[*TerminateEvent]{},
		// onConfUpdated: &hook.Hook[*ConfUpdatedEvent]{},
	}
	app.injectDefaultServices()
	app.registerDefaultHooks()

	return app
}

func (app *App) injectDefaultServices() error {
	// 请勿依赖注入顺序

	// DAO
	if app.dao == nil {
		if err := app.initDao(); err != nil {
			return err
		}
	}
	return nil

}

func (app *App) registerDefaultHooks() {

}

func (app *App) Conf() *config.Config {
	return app.conf
}

func (app *App) Dao() *dao.Dao {
	return app.dao
}

func (app *App) initDao() error {
	dbInstance, err := db.Newdb(app.conf.DB)
	if err != nil {
		return fmt.Errorf("db init err: %w", err)
	}
	app.SetDao(dao.NewDao(dbInstance))

	return nil
}

func (app *App) SetDao(dao *dao.Dao) {
	app.dao = dao
}
