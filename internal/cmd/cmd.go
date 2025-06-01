package cmd

import (
	"context"

	"demo/internal/controller/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp" //HTTP sever in GoFrame
	"github.com/gogf/gf/v2/os/gcmd"   // CLI command in GoFrame

	"demo/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()                                        //init HTTP sever in GoFrame
			s.AddStaticPath("/swagger", "resource/public/swagger") //Swager UI
			//Define group of route
			s.Group("/", func(group *ghttp.RouterGroup) {
				//Middleware resolve reponse in GoFrame
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				//Blind controller hello and user into router
				group.Bind(
					hello.NewV1(),
					user.NewV1(),
				)
			})
			s.Run() //init sever
			return nil
		},
	}
)
