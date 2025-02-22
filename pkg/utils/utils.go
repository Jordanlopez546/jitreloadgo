package utils

import (
	"os"
	"path/filepath"
)

// Contains checks if a string exists in a slice of strings
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}

// GetTempBinaryPath returns the path for the temporary binary
func GetTempBinaryPath() string {
	return filepath.Join(os.TempDir(), "jitreloadgo-bin")
}

// IsDirectory checks if a path is a directory
func IsDirectory(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return info.IsDir()
}

// EnsureDirectory ensures a directory exists, creates if it doesn't
func EnsureDirectory(path string) error {
	if !IsDirectory(path) {
		return os.MkdirAll(path, 0755)
	}

	return nil
}
