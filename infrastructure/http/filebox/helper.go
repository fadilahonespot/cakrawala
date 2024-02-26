package filebox

import (
	"path/filepath"
	"strings"
)

func cutStringFile(nama string, fileURL string) string {
	extensionFile := filepath.Ext(nama)
	file := strings.ToLower(extensionFile)
	if file == ".jpg" || file == ".png" || file == ".jpeg" || file == ".gif" || file == ".mp4" {
		fileURL = fileURL[:len(fileURL)-4]
		return fileURL + "raw=1"
	}

	fileURL = fileURL[:len(fileURL)-1]
	return fileURL + "1"
}