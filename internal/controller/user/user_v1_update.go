package user

import (
    "context"

    "demo/api/user/v1"
    "demo/internal/dao"
    "demo/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
    _, err = dao.User.Ctx(ctx).Data(do.User{
        Name:   req.Name,
        Status: req.Status,
        Age:    req.Age,
    }).WherePri(req.Id).Update()
    return
}
