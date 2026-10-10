//go:build js && wasm

package filesystem

import (
	"errors"
	"path/filepath"
	"strings"
	"time"
    "fmt"
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
    if err := vfs.CreateDirectory(filepath.Dir(path)); err != nil {
        return err
    }
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

func (vfs *VFS) CreateDirectory(path string) error {
    path = vfs.cleanPath(path)

    if path == "." {
        return nil
    }

    // Create every directory in the path, starting with its parents.
    current := ""
    for _, part := range strings.Split(path, "/") {
        if current == "" {
            current = part
        } else {
            current += "/" + part
        }

        if isDir, exists := vfs.dirs[current]; exists {
            if !isDir {
                return fmt.Errorf("%q is a file, not a directory", current)
            }
            continue
        }

        vfs.dirs[current] = true
    }

    return nil
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

func (fs *VFS) GetFilenames(directory string, recursive bool) ([]string, error) {
    result := []string{}

    for path := range fs.files {
        if recursive {
			if filepath.Dir(path) == directory ||
				directory == "." ||
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

    for path, isDir := range fs.dirs {
        if !isDir || path == directory {
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

// TODO --- will need to do this properly to implement the hub.
func (fs VFS) ModTime(path string) time.Time {
    return time.Time{}
}

func (vfs *VFS) DeleteDirectory(path string) error {
    path = vfs.cleanPath(path)

    if path == "." {
        return errors.New("cannot delete the root directory")
    }

    isDir, exists := vfs.dirs[path]
    if !exists || !isDir {
        return fmt.Errorf("directory %q does not exist", path)
    }

    prefix := path + "/"

    // Remove files in the directory and all its descendants.
    for name := range vfs.files {
        if name == path || strings.HasPrefix(name, prefix) {
            delete(vfs.files, name)
        }
    }

    // Remove the directory and all descendant entries.
    for name := range vfs.dirs {
        if name == path || strings.HasPrefix(name, prefix) {
            delete(vfs.dirs, name)
        }
    }

    return nil
}

func (vfs *VFS) Rename(oldPath, newPath string) error {
    oldPath = vfs.cleanPath(oldPath)
    newPath = vfs.cleanPath(newPath)

    if oldPath == "." || newPath == "." {
        return errors.New("cannot rename the root directory")
    }

    if oldPath == newPath {
        return nil
    }

    isDir, exists := vfs.dirs[oldPath]
    if !exists {
        return fmt.Errorf("%q does not exist", oldPath)
    }

    if _, exists := vfs.dirs[newPath]; exists {
        return fmt.Errorf("%q already exists", newPath)
    }

    // A directory cannot be moved inside itself.
    if isDir && strings.HasPrefix(newPath, oldPath+"/") {
        return errors.New("cannot move a directory inside itself")
    }

    if isDir {
        oldPrefix := oldPath + "/"
        newPrefix := newPath + "/"

        // Collect directory entries before changing the map.
        dirs := make(map[string]bool)
        for name, value := range vfs.dirs {
            if strings.HasPrefix(name, oldPrefix) {
                dirs[newPrefix+strings.TrimPrefix(name, oldPrefix)] = value
            }
        }

        // Collect file contents before changing the map.
        files := make(map[string][]byte)
        for name, data := range vfs.files {
            if strings.HasPrefix(name, oldPrefix) {
                files[newPrefix+strings.TrimPrefix(name, oldPrefix)] = data
            }
        }

        // Remove the old directory tree.
        for name := range vfs.dirs {
            if name == oldPath || strings.HasPrefix(name, oldPrefix) {
                delete(vfs.dirs, name)
            }
        }
        for name := range vfs.files {
            if strings.HasPrefix(name, oldPrefix) {
                delete(vfs.files, name)
            }
        }

        // Install the renamed directory tree.
        vfs.dirs[newPath] = true
        for name, value := range dirs {
            vfs.dirs[name] = value
        }
        for name, data := range files {
            vfs.files[name] = data
        }

        return nil
    }

    // Rename a file.
    data := vfs.files[oldPath]
    delete(vfs.files, oldPath)
    delete(vfs.dirs, oldPath)

    vfs.files[newPath] = data
    vfs.dirs[newPath] = false

    return nil
}

