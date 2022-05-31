package models

type AdminRoleUsers struct {
	RoleId    int    `json:"role_id"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
