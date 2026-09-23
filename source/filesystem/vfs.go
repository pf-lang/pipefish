//go:build js && wasm

package filesystem

import (
	"errors"
	"path/filepath"
	"strings"
)

func (vfs *VFS) ReadFile(path string) ([]byte, error) {
	path = vfs.cleanPath(path)

	data, ok := vfs.files[path]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	// We don't let callers modify the VFS's stored copy.
	return append([]byte(nil), data...), nil
}

func (vfs *VFS) WriteFile(path string, data []byte) error {
	path = vfs.cleanPath(path)
	vfs.files[path] = append([]byte(nil), data...)
	vfs.dirs[path] = false
	return nil
}

func (vfs *VFS) DeleteFile(path string) error {
	path = vfs.cleanPath(path)
	if _, exists := vfs.files[path]; !exists {
		return errors.New("file does not exist")
	}
	delete(vfs.files, path)
	delete(vfs.dirs, path)
	return nil
}

func (vfs *VFS) FileExists(path string) bool {
	path = vfs.cleanPath(path)
	_, exists := vfs.files[path]
	return exists
}

func (vfs *VFS) cleanPath(path string) string {
	path = filepath.ToSlash(path)
	path = strings.TrimPrefix(path, "./")
	path = strings.TrimPrefix(path, "/")
	if path == "" {
		return "."
	}
	return path
}

type vfsFileInfo struct {
	name  string
	isDir bool
}

func (info vfsFileInfo) Name() string {
	return info.name
}

func (info vfsFileInfo) IsDir() bool {
	return info.isDir
}

func (fs *VFS) GetFilenames(directory string, recursive bool) ([]string, error) {
    result := []string{}

    for path := range fs.files {
        if recursive {
            if filepath.Dir(path) == directory ||
                strings.HasPrefix(path, directory+string(filepath.Separator)) {
                result = append(result, path)
            }
        } else if filepath.Dir(path) == directory {
            result = append(result, path)
        }
    }

    return result, nil
}

func (fs *VFS) GetDirectoryNames(directory string, recursive bool) ([]string, error) {
    result := []string{}

    for path := range fs.dirs {
        if path == directory {
            continue
        }

        if recursive {
            if filepath.Dir(path) == directory ||
                strings.HasPrefix(path, directory+string(filepath.Separator)) {
                result = append(result, path)
            }
        } else if filepath.Dir(path) == directory {
            result = append(result, path)
        }
    }

    return result, nil
}