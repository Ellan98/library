/*
 * @Date: 2024-07-04 11:23:35
 * @LastEditTime: 2024-07-18 16:08:26
 * @FilePath: \library_room\internal\core\base.go
 * @description: 注释
 */
/*
 * @Date: 2024-07-04 11:23:35
 * @LastEditTime: 2024-07-05 08:43:31
 * @FilePath: \library_room\internal\core\base.go
 * @description: 注释
 */
package core

import (
	"library_room/internal/config"
	"library_room/internal/dao"

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

func (app *App) injectDefaultServices() {
	// 请勿依赖注入顺序

}

func (app *App) registerDefaultHooks() {

}

func (app *App) Conf() *config.Config {
	return app.conf
}
