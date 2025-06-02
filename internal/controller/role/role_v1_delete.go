package role

import (
	"context"

	v1 "demo/api/role/v1"
	"demo/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.DeleteRes{}
	return
}
