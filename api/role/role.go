// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"demo/api/role/v1"
)

type IRoleV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetUsersByRole(ctx context.Context, req *v1.GetUsersByRoleReq) (res *v1.GetUsersByRoleRes, err error)
	GetRolesByUser(ctx context.Context, req *v1.GetRolesByUserReq) (res *v1.GetRolesByUserRes, err error)
}
