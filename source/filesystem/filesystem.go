package filesystem

import "time"

type VFS struct {
	files map[string][]byte
	dirs  map[string]bool
}

type FileSystem interface {
    ReadFile(path string) ([]byte, error)
    WriteFile(path string, data []byte) error
    DeleteFile(path string) error
    FileExists(path string) bool

	Rename(oldPath, newPath string) error

    GetFilenames(directory string, recursive bool) ([]string, error)
    GetDirectoryNames(directory string, recursive bool) ([]string, error)
    CreateDirectory(path string) error
    DeleteDirectory(path string) error

    ModTime(path string) time.Time
}

type FileInfo interface {
	Name() string
	IsDir() bool
}

func NewVFS() *VFS {
	return &VFS{
		files: make(map[string][]byte),
		dirs:  map[string]bool{".": true},
	}
}

