//go:build !js && !wasm

package filesystem

import (
    "fmt"
	"os"
	"path/filepath"
    "time"
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

func (fs OSFileSystem) ModTime(path string) time.Time {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return time.Time{}
    }
    return fileInfo.ModTime()
}

// This is here for testing purposes.
func NewVFSFromDirectory(path string) (*VFS, error) {
	vfs := NewVFS()
	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relativePath, err := filepath.Rel(path, currentPath)
		if err != nil {
			return err
		}
		if relativePath == "." {
			return nil
		}
		relativePath = filepath.ToSlash(relativePath)
		if info.IsDir() {
			vfs.dirs[relativePath] = true
			return nil
		}
		data, err := os.ReadFile(currentPath)
		if err != nil {
			return err
		}
		vfs.files[relativePath] = data
		return nil
	})
	if err != nil {
		return nil, err
	}
	return vfs, nil
}

func (fs OSFileSystem) CreateDirectory(path string) error {
    return os.MkdirAll(path, 0755)
}

func (fs OSFileSystem) DeleteDirectory(path string) error {
    info, err := os.Stat(path)
    if err != nil {
        return err
    }

    if !info.IsDir() {
        return fmt.Errorf("%q is not a directory", path)
    }

    return os.RemoveAll(path)
}

func (fs OSFileSystem) Rename(oldPath, newPath string) error {
    if oldPath == newPath {
        return nil
    }

    if _, err := os.Stat(oldPath); err != nil {
        return err
    }

    // Don't silently replace an existing file or directory.
    if _, err := os.Lstat(newPath); err == nil {
        return fmt.Errorf("%q already exists", newPath)
    } else if !os.IsNotExist(err) {
        return err
    }

    return os.Rename(oldPath, newPath)
}
