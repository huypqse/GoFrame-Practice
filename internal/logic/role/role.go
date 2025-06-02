package role

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() service.IRole {
	return &sRole{}
}

// Create role
func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (id int64, err error) {
	err = dao.Role.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := dao.Role.Ctx(ctx).Data(do.Role{
			Name: in.Name,
		}).Insert()
		if err != nil {
			return err
		}
		id, err = result.LastInsertId()
		return err
	})
	return
}

// Delete role
func (s *sRole) Delete(ctx context.Context, id int64) error {
	result, err := dao.Role.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return gerror.Newf("Role id %d not found", id)
	}
	return nil
}

// Update role
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	_, err := dao.Role.Ctx(ctx).Where("id", in.Id).Data(do.Role{
		Name: in.Name,
	}).Update()
	return err
}

// Get one role
func (s *sRole) GetOne(ctx context.Context, id int64) (*entity.Role, error) {
	var role entity.Role
	err := dao.Role.Ctx(ctx).Where("id", id).Scan(&role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// Get list of roles
func (s *sRole) GetList(ctx context.Context) ([]*entity.Role, error) {
	var roles []*entity.Role
	err := dao.Role.Ctx(ctx).Scan(&roles)
	return roles, err
}

// Get users by role (many-to-many)
func (s *sRole) GetUsersByRole(ctx context.Context, roleId int64) ([]*entity.User, error) {
	var users []*entity.User
	err := dao.User.Ctx(ctx).
		LeftJoin("user_role", "user.id=user_role.user_id").
		Where("user_role.role_id", roleId).
		Scan(&users)
	return users, err
}

// Get roles by user (many-to-many)
func (s *sRole) GetRolesByUser(ctx context.Context, userId int64) ([]*entity.Role, error) {
	var roles []*entity.Role
	err := dao.Role.Ctx(ctx).
		LeftJoin("user_role", "role.id=user_role.role_id").
		Where("user_role.user_id", userId).
		Scan(&roles)
	return roles, err
}
