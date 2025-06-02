package role

import (
	"context"

	v1 "demo/api/role/v1"
	"demo/internal/model"
	"demo/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.Role().Update(ctx, model.RoleUpdateInput{
		Id:   req.Id,
		Name: derefString(req.Name),
	})
	if err != nil {
		return nil, err
	}
	res = &v1.UpdateRes{}
	return
}
func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
