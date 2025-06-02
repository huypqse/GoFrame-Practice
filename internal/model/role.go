package model

type RoleCreateInput struct {
	Name string
}

type RoleUpdateInput struct {
	Id   int64
	Name string
}
