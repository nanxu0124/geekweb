package accesslog

import (
	"geekweb/geek"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	b := NewBuilder()
	s := geek.NewHTTPServer()
	s.Get("/", func(ctx *geek.Context) {
		ctx.Resp.Write([]byte("hello, world"))
	})
	s.Get("/user", func(ctx *geek.Context) {
		time.Sleep(time.Second)
		ctx.RespData = []byte("hello, user")
	})
	s.Use(b.Build())
	s.Start(":8081")
}
