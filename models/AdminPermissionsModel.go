package models

type AdminPermissions struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	HttpMethod string `json:"http_method"`
	HttpPath   string `json:"http_path"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
