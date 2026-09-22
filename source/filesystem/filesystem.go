package filesystem

type FileSystem interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
	DeleteFile(path string) error
	FileExists(path string) bool
	ReadDir(path string) ([]FileInfo, error)
}

type FileInfo interface {
	Name() string
	IsDir() bool
}