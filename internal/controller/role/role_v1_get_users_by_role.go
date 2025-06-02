package role

import (
	"context"

	v1 "demo/api/role/v1"
	"demo/internal/service"
)

func (c *ControllerV1) GetUsersByRole(ctx context.Context, req *v1.GetUsersByRoleReq) (res *v1.GetUsersByRoleRes, err error) {
	users, err := service.Role().GetUsersByRole(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUsersByRoleRes{List: users}
	return
}
