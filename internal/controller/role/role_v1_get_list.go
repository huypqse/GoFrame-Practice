package role

import (
	"context"

	v1 "demo/api/role/v1"
	"demo/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	roles, err := service.Role().GetList(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetListRes{List: roles}
	return

}
