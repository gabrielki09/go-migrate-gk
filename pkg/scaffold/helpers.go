package scaffold

import (
	"os"
	"strings"
)

func createInformedPath(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func globalTrimSpace(value string) string {
	return strings.TrimSpace(value)
}

func globalToLower(value string) string {
	return strings.ToLower(value)
}

func invertBarPath(value string) string {
	return strings.ReplaceAll(value, "\\", "/")
}
