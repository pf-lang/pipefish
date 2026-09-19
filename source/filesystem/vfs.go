package filesystem

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type VFS struct {
	files map[string][]byte
	dirs  map[string]bool
}

func NewVFS() *VFS {
	return &VFS{
		files: make(map[string][]byte),
		dirs:  map[string]bool{".": true},
	}
}

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

func (vfs *VFS) ReadFile(path string) ([]byte, error) {
	path = vfs.cleanPath(path)

	data, ok := vfs.files[path]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	// We don't let callers modify the VFS's stored copy.
	return append([]byte(nil), data...), nil
}

func (vfs *VFS) CreateFile(path string) error {
	path = vfs.cleanPath(path)
	if _, exists := vfs.files[path]; exists {
		return errors.New("file already exists")
	}
	vfs.files[path] = []byte{}
	return nil
}

func (vfs *VFS) WriteFile(path string, data []byte) error {
	path = vfs.cleanPath(path)
	if _, exists := vfs.files[path]; !exists {
		return errors.New("file does not exist")
	}
	vfs.files[path] = append([]byte(nil), data...)
	return nil
}

func (vfs *VFS) DeleteFile(path string) error {
	path = vfs.cleanPath(path)
	if _, exists := vfs.files[path]; !exists {
		return errors.New("file does not exist")
	}
	delete(vfs.files, path)
	return nil
}

func (vfs *VFS) FileExists(path string) bool {
	path = vfs.cleanPath(path)
	_, exists := vfs.files[path]
	return exists
}

func (vfs *VFS) ReadDir(path string) ([]FileInfo, error) {
	path = vfs.cleanPath(path)
	if !vfs.dirs[path] {
		return nil, errors.New("directory does not exist")
	}
	var result []FileInfo
	prefix := path
	if prefix != "." {
		prefix += "/"
	} else {
		prefix = ""
	}

	for filePath := range vfs.files {
		if strings.HasPrefix(filePath, prefix) {
			remainder := strings.TrimPrefix(filePath, prefix)

			if !strings.Contains(remainder, "/") {
				result = append(result, vfsFileInfo{
					name:  remainder,
					isDir: false,
				})
			}
		}
	}

	for dirPath := range vfs.dirs {
		if dirPath == path {
			continue
		}

		if strings.HasPrefix(dirPath, prefix) {
			remainder := strings.TrimPrefix(dirPath, prefix)

			if !strings.Contains(remainder, "/") {
				result = append(result, vfsFileInfo{
					name:  remainder,
					isDir: true,
				})
			}
		}
	}

	return result, nil
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