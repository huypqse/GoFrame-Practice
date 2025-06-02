// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"demo/internal/model"
	"demo/internal/model/entity"
)

type (
	IRole interface {
		// Create role
		Create(ctx context.Context, in model.RoleCreateInput) (id int64, err error)
		// Delete role
		Delete(ctx context.Context, id int64) error
		// Update role
		Update(ctx context.Context, in model.RoleUpdateInput) error
		// Get one role
		GetOne(ctx context.Context, id int64) (*entity.Role, error)
		// Get list of roles
		GetList(ctx context.Context) ([]*entity.Role, error)
		// Get users by role (many-to-many)
		GetUsersByRole(ctx context.Context, roleId int64) ([]*entity.User, error)
		// Get roles by user (many-to-many)
		GetRolesByUser(ctx context.Context, userId int64) ([]*entity.Role, error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
