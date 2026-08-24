## Overview 

The `filepath` library implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths. 

The `filepath` package uses either forward slashes or backslashes, depending on the operating system. To process paths such as URLs that always use forward slashes regardless of the operating system, see the `path` package. 

Features of the corresponding Golang standard library that require IO to work have been moved to the `files` package ifthey have been implemented at all.  

## Modules 

### ` ("path/filepath")`
## Functions 

### `abs(path string)`

`abs` returns an absolute representation of `path`. If the path is not absolute it will be joined with the current working directory to turn it into an absolute path. The absolute path name for a given file is not guaranteed to be unique. `abs` calls `clean` on the result.  

### `base(path string) -> string`

`base` returns the last element of `path`. Trailing path separators are removed before extracting the last element. If the path is empty, Base returns ".". If the path consists entirely of separators, `base` returns a single separator.  

### `clean(path string) -> string`

Returns the shortest path name equivalent to `path` by purely lexical processing.  

### `dir(path string) -> string`

`dir` returns all but the last element of `path`, typically the path's directory. After dropping the final element, `dir` calls `clean` on the path and trailing slashes are removed. If the path is empty, `dir` returns ".". If the path consists entirely of separators, Dir returns a single separator. The returned path does not end in a separator unless it is the root directory.  

### `evalSymlinks(path string)`

`evalSymlinks` returns the path name after the evaluation of any symbolic links. If `path` is relative, the result will be relative to the current directory, unless one of the components is an absolute symbolic link. EvalSymlinks calls Clean on the result.  

### `ext(path string) -> string`

Returns the filename extension used by path. The extension is the suffix beginning at the final dot in the final element of path; it is empty if there is no dot.  

### `fromSlash(path string) -> string`

Returns the result of replacing each slash ('/') character in path with a separator character. Multiple slashes are replaced by multiple separators.  

### `isAbs(path string) -> bool`

Tests whether the path is absolute.  

### `isLocal(path string) -> bool`

`isLocal` reports whether `path`, using lexical analysis only, has all of these properties: 


- is within the subtree rooted at the directory in which path is evaluated
- is not an absolute path
- is not empty


`IsLocal` is a purely lexical operation. In particular, it does not account for the effect of any symbolic links that may exist in the filesystem.  

### `join(elem ... string) -> string`

Join joins any number of path elements into a single path, separating them with an OS- specific Separator. Empty elements are ignored. The result is `clean`ed. However, if the argument list is empty or all its elements are empty, Join returns an empty string.  

### `localize(path string)`

`localize` converts a slash-separated path into an operating system path. The input path must be a valid path. It returns an error if the path cannot be represented by the operating system. 

The path returned by Localize will always be local, as reported by IsLocal.  

### `match(pattern string, name string)`

`match` reports whether `name` matches `pattern`, using the usual pattern-matching for the shell, e.g. `*` as a wildcard character.  

### `rel(basepath string, targpath string)`

`rel` returns a relative path that is lexically equivalent to `targPath` when joined to `basePath` with an intervening separator. That is, `join(basePath, rel(basePath, targPath))` is equivalent to `targPath` itself. 

The returned path will always be relative to basePath, even if basePath and targPath share no elements. `rel` calls `clean` on the result. 

An error is returned if targPath can't be made relative to basePath or if knowing the current working directory would be necessary to compute it.  

### `split(path string) -> string, string`

`split` splits `path` immediately following the final separator, separating it into a directory and file name component. If there is no separator in `path`, `split` returns `""` for the directory and `path` for the file.  

### `splitList(path string) -> list`

splitList splits a list of paths joined by the OS-specific list separator. Unlike `strings.split`, `splitList` returns an empty list when passed an empty string.  

### `toSlash(path string) -> string`

`toSlash` returns the result of replacing each separator character in path with a slash ('/') character. Multiple separators are replaced by multiple slashes.  

