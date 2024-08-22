package ioc

import (
	"geekweb/geek"
	"geekweb/geek/middleware/accesslog"
	"geekweb/internal/web"
)

func InitWebServer(midls []geek.Middleware, userHdl *web.UserHandler) *geek.HTTPServer {
	server := geek.NewHTTPServer()
	server.Use(midls...)
	userHdl.RegisterRoutes(server)
	return server
}

func InitMiddlewares() []geek.Middleware {
	return []geek.Middleware{
		accesslog.NewBuilder().Build(),
	}
}
