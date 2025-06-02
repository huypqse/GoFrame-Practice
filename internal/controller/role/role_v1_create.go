package role

import (
	"context"
	v1 "demo/api/role/v1"
	"demo/internal/model"
	"demo/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := service.Role().Create(ctx, model.RoleCreateInput{Name: req.Name})
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{Id: id}
	return
}
