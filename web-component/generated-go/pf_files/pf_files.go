package pf_files

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Filenames struct {
	Directory      string
	FileSearchMode FileSearchMode
}

type DirectoryNames struct {
	Directory      string
	FileSearchMode FileSearchMode
}

type FileSearchMode int

const (
	NONRECURSIVE FileSearchMode = iota
	RECURSIVE
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Filenames":      func(t uint32, v any) any { return Filenames{v.([]any)[0].(string), v.([]any)[1].(FileSearchMode)} },
	"DirectoryNames": func(t uint32, v any) any { return DirectoryNames{v.([]any)[0].(string), v.([]any)[1].(FileSearchMode)} },
	"FileSearchMode": func(t uint32, v any) any { return FileSearchMode(v.(int)) },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Filenames":      (*Filenames)(nil),
	"DirectoryNames": (*DirectoryNames)(nil),
	"FileSearchMode": (*FileSearchMode)(nil),
}

func GoFileExists(fname string) any {
	_, err := os.Stat(fname)
	return err == nil
}

func GoGetFile(fname string) any {
	fileContent, err := ioutil.ReadFile(fname)
	if err != nil {
		return errors.New("can't find file '" + fname + "'")
	}
	return string(fileContent)
}

func GoGetFilenames(filenames Filenames) any {
	resultList := []string{}
	if filenames.FileSearchMode == 0 {
		fileInfos, err := os.ReadDir(filenames.Directory)
		if err != nil {
			return err
		}
		for _, info := range fileInfos {
			if !info.IsDir() {
				resultList = append(resultList, filepath.Join(filenames.Directory, info.Name()))
			}
		}
	}
	err := filepath.Walk(filenames.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			resultList = append(resultList, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return resultList
}

func GoGetDirectoryNames(directoryNames DirectoryNames) any {
	resultList := []string{}
	if directoryNames.FileSearchMode == 0 {
		fileInfos, err := os.ReadDir(directoryNames.Directory)
		if err != nil {
			return err
		}
		for _, info := range fileInfos {
			if info.IsDir() {
				resultList = append(resultList, filepath.Join(directoryNames.Directory, info.Name()))
			}
		}
	}
	err := filepath.Walk(directoryNames.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			resultList = append(resultList, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return resultList
}

func GoDeleteFile(fname string) any {
	err := os.Remove(fname)
	if err != nil {
		return errors.New("can't delete file '" + fname + "'")
	} else {
		return struct{}{}
	}
}

func GoPutStringInFile(output string, fname string) any {
	f, err := os.Create(fname)
	if err != nil {
		return errors.New("can't access file '" + fname + "'")
	}
	defer f.Close()
	_, err2 := f.WriteString(output)

	if err2 != nil {
		return errors.New("can't write to file '" + fname + "'")
	}
	return struct{}{}
}
