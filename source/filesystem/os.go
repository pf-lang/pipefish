package filesystem

import (
	"os"
)

type OSFileSystem struct{}

func (OSFileSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (OSFileSystem) WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}

func (OSFileSystem) DeleteFile(path string) error {
	return os.Remove(path)
}

func (OSFileSystem) FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (OSFileSystem) ReadDir(path string) ([]FileInfo, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]FileInfo, len(entries))
	for i, entry := range entries {
		result[i] = entry
	}

	return result, nil
}