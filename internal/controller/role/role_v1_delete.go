package role

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "demo/api/role/v1"
	"demo/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	g.Log().Infof(ctx, "Delete role with id: %v", req.Id)
	if req.Id == 0 {
		return nil, gerror.New("Invalid id")
	}
	err = service.Role().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.DeleteRes{}
	return
}
