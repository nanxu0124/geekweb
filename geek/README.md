### geek web框架

geek 是一个轻量级的、受 Gin 启发的 Web 框架。虽然功能没有 Gin 强大，但它提供了一个简单高效的方式来构建 Web 应用程序，具备基本的功能，如路由、上下文管理和中间件支持。

### 功能特性
- **路由树**：基于 / 分割的高效路由，支持快速且有序的请求处理。支持通配符、参数路径。
- **上下文(Context)**:为每个请求提供一个独立的上下文，用于在请求的生命周期内传递数据和处理响应，支持请求、响应、参数等信息的管理。
- **中间件**：支持中间件机制，当前内置了访问日志中间件，帮助开发者轻松记录和监控请求信息。

### 使用示例

~~~go
package main

import (
	"fmt"
	"geek"
)

func main() {
	s := geek.NewHTTPServer()
	s.Get("/", func(ctx *geek.Context) {
		ctx.Resp.Write([]byte("hello, world"))
	})
	s.Get("/user", func(ctx *geek.Context) {
		ctx.Resp.Write([]byte("hello, user"))
	})

	s.Post("/form", func(ctx *geek.Context) {
		err := ctx.Req.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
	})

	s.Start(":8081")
}
~~~

