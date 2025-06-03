package hello

import (
	"context"

	v1 "demo/api/hello/v1"

	"github.com/gogf/gf/v2/frame/g"
)

// Hello handles the /hello endpoint for API version 1.
// It writes "Hello World!" to the HTTP response.
func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	// token, _ := service.JWTAuth().LoginHandler(ctx)
	// g.RequestFromCtx(ctx).Response.Writeln(token)
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
