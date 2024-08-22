//go:build wireinject

package main

import (
	"geekweb/geek"
	"geekweb/internal/repository"
	"geekweb/internal/repository/cache"
	"geekweb/internal/repository/dao"
	"geekweb/internal/service"
	"geekweb/internal/web"
	"geekweb/ioc"
	"github.com/google/wire"
)

func initWebServer() *geek.HTTPServer {
	wire.Build(
		// 最基础的第三方依赖
		ioc.InitDB, ioc.InitCache,

		dao.NewUserDAO,

		cache.NewRedisUserCache,

		repository.NewUserRepository,

		service.NewUserService,

		web.NewUserHandler,

		ioc.InitMiddlewares,
		ioc.InitWebServer,
	)

	return new(geek.HTTPServer)
}
