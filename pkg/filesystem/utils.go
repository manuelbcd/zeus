package filesystem

import (
	"os"
)

// Exists checks if a file or directory exists using and absolutePath.
func Exists(absolutePath string) bool {
	_, err := os.Stat(absolutePath)
	return !os.IsNotExist(err)
}
