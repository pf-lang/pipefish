package filesystem

type FileSystem interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
	DeleteFile(path string) error
	FileExists(path string) bool
	GetFilenames(directory string, recursive bool) ([]string, error)
    GetDirectoryNames(directory string, recursive bool) ([]string, error)
}

type FileInfo interface {
	Name() string
	IsDir() bool
}

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

