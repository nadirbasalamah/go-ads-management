package utils

type Role string

const (
	ROLE_USER  Role = "user"
	ROLE_ADMIN Role = "admin"
)

var ALLOWED_EXTENSIONS map[string]bool = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
}
