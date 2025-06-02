// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
		"demo/internal/model/entity"
	    "github.com/gogf/gf/v2/frame/g"

)
//Generate controller when u already have define api in package api
// Status marks role status.
type Status bool

const (
    StatusOK       Status = false // Role is OK.
    StatusDisabled Status = true // Role is disabled.
)

type CreateReq struct { //@RequestBody DTO in Spring
    g.Meta `path:"/role" method:"post" tags:"Role" summary:"Create role"`
    Name   string `v:"required|length:3,30" dc:"Role name"`
}

type CreateRes struct {// response body like DTO
    Id int64 `json:"id" dc:"Role id"`
}

type DeleteReq struct {
    g.Meta `path:"/role/{id}" method:"delete" tags:"Role" summary:"Delete role"`
    Id     int64 `v:"required" dc:"role id"`
}
type DeleteRes struct{}

type UpdateReq struct {
    g.Meta `path:"/role/{id}" method:"put" tags:"Role" summary:"Update role"`
    Id     int64   `v:"required" dc:"role id"`
    Name   *string `v:"length:3,10" dc:"role name"`
}
type UpdateRes struct{}

type GetOneReq struct {
    g.Meta `path:"/role/{id}" method:"get" tags:"Role" summary:"Get one role"`
    Id     int64 `v:"required" dc:"role id"`
}
type GetOneRes struct {
    *entity.Role `dc:"role"`
}
type GetListReq struct {
    g.Meta `path:"/roles" method:"get" tags:"Role" summary:"Get roles"`
}
type GetListRes struct {
    List []*entity.Role `json:"list" dc:"role list"`
}
// Get list user in this role
type GetUsersByRoleReq struct {
    g.Meta `path:"/role/{id}/users" method:"get" tags:"Role" summary:"Get users by role"`
    Id     int64 `v:"required" dc:"Role id"`
}
type GetUsersByRoleRes struct {
    List []*entity.User `json:"list" dc:"User list"`
}

//Get List Role from User
type GetRolesByUserReq struct {
    g.Meta `path:"/user/{id}/roles" method:"get" tags:"Role" summary:"Get roles by user"`
    Id     int64 `v:"required" dc:"User id"`
}
type GetRolesByUserRes struct {
    List []*entity.Role `json:"list" dc:"Role list"`
}
