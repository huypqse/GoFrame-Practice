package role

import (
	"context"

	v1 "demo/api/role/v1"
	"demo/internal/model"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	g.Log().Infof(ctx, "Update role: id=%v, name=%v", req.Id, req.Name)
	name := ""
	if req.Name != nil {
		name = *req.Name
	}
	err = service.Role().Update(ctx, model.RoleUpdateInput{
		Id:   req.Id,
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.UpdateRes{}
	return
}
