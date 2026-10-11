package filesystem

import (
	"time"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type VFS struct {
	files *orderedmap.OrderedMap[string, []byte]
	dirs  *orderedmap.OrderedMap[string, bool]
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
	dirs := orderedmap.New[string, bool]()
	dirs.Set(".", true)
	return &VFS{
		files: orderedmap.New[string, []byte](),
		dirs:  dirs,
	}
}

