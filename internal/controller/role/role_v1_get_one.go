package role

import (
	"context"

	v1 "demo/api/role/v1"
	"demo/internal/service"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	role, err := service.Role().GetOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetOneRes{Role: role}
	return
}
