package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func GenerateFilename(file *multipart.FileHeader) string {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	return fmt.Sprintf("%s%s", uuid.New().String(), ext)
}
