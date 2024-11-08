package utils

func ValidateFile(ext string) bool {
	return ALLOWED_EXTENSIONS[ext]
}
