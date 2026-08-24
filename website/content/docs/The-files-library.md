## Overview 

The `files` library implements some basic file-handling commands.  

## Modules 

### ` ("io/ioutil")`

### ` ("path/filepath")`

### ` ("os")`

### ` ("errors")`
## Types 

### `FileSearchMode = enum NONRECURSIVE, RECURSIVE`

The search mode to use when calling `get ... from Filenames` or `get ... from DirectoryNames`.  

### `File = struct (filepath string)`

Data structure containing the filepath of a file.  

### `FileExists = struct (filepath string)`

Another data structure containing the filepath of a file, for when we want `get ... from` to check for existence.  

### `Filenames = struct (directory string, fileSearchMode FileSearchMode)`

A struct to pass to `get ... from` to get the names of the files in a directory.  

### `DirectoryNames = struct (directory string, fileSearchMode FileSearchMode)`

A struct to pass to `get ... from` to get the names of the subdirectories of a directory.  
## Commands 

### `get(L ref) from (filenames Filenames)`

Populates the reference variable with filenames fitting the criteria in the `Filenames` struct.  

### `get(L ref) from (directoryNames DirectoryNames)`

Populates the reference variable with filenames fitting the criteria in the `Filenames` struct.  

### `get(contents ref) from (file File)`

Populates the reference variable with the contents of the file as a string.  

### `get(x ref) from (file FileExists)`

Populates the reference variable with a `bool` saying whether the file exists.  

### `put(s string) into (file File)`

Writes the string to the given file, creating it if it doesn't exist and overwriting it if it already exists.  

### `delete(file File)`

Deletes the given file.  

### `goDeleteFile(fname string)`

### `goPutStringInFile(output string, fname string)`
## Functions 

### `goFileExists(fname string)`

### `goGetFile(fname string)`

### `goGetFilenames(filenames Filenames)`

### `goGetDirectoryNames(directoryNames DirectoryNames)`

