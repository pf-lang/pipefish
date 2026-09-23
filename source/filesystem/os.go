package filesystem

import (
	"os"
	"path/filepath"
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

func (fs OSFileSystem) GetFilenames(directory string, recursive bool) ([]string, error) {
    if !recursive {
        entries, err := os.ReadDir(directory)
        if err != nil {
            return nil, err
        }

        result := []string{}
        for _, entry := range entries {
            if !entry.IsDir() {
                result = append(result, filepath.Join(directory, entry.Name()))
            }
        }
        return result, nil
    }

    result := []string{}
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            result = append(result, path)
        }
        return nil
    })

    return result, err
}

func (fs OSFileSystem) GetDirectoryNames(directory string, recursive bool) ([]string, error) {
    if !recursive {
        entries, err := os.ReadDir(directory)
        if err != nil {
            return nil, err
        }

        result := []string{}
        for _, entry := range entries {
            if entry.IsDir() {
                result = append(result, filepath.Join(directory, entry.Name()))
            }
        }
        return result, nil
    }

    result := []string{}
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() && path != directory {
            result = append(result, path)
        }
        return nil
    })

    return result, err
}