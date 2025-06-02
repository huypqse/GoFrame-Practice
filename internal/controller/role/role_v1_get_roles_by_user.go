package role

import (
	"context"

	v1 "demo/api/role/v1"
	"demo/internal/service"
)

func (c *ControllerV1) GetRolesByUser(ctx context.Context, req *v1.GetRolesByUserReq) (res *v1.GetRolesByUserRes, err error) {
	roles, err := service.Role().GetRolesByUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetRolesByUserRes{List: roles}
	return
}
