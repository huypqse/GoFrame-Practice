package model

type RoleCreateInput struct {
	Name string
}

type RoleUpdateInput struct {
	Id   int64  `json:"id" in:"path"` // Lấy id từ path
	Name string `json:"name"`         // Lấy name từ body hoặc query
}
type DeleteReq struct {
	Id int64 `json:"id" in:"path"`
}
