## Overview 

The `path` library supplies utility routines for manipulating slash-separated paths.  

## Modules 

### ` ("path")`

### ` ("errors")`
## Functions 

### `base(p string) -> string`

`base` returns the last element of `path`. Trailing slashes are removed before extracting the last element. If the path is empty, `base` returns ".". If the path consists entirely of slashes, `base` returns "/".  

### `clean(p string) -> string`

`clean` returns the shortest path name equivalent to path by purely lexical processing. It applies the following rules iteratively until no further processing can be done: 


- Replace multiple slashes with a single slash.
- Eliminate each . path name element (the current directory).
- Eliminate each inner .. path name element (the parent directory) along with the non-.. element that precedes it.
- Eliminate .. elements that begin a rooted path: that is, replace "/.." by "/" at the beginning of a path.


The returned path ends in a slash only if it is the root "/". 

If the result of this process is an empty string, `clean` returns the string "."  

### `dir(p string) -> string`

`dir` returns all but the last element of `path`, typically the path's directory. After dropping the final element using `split`, the path is `clean`ed and trailing slashes are removed. If the path is empty, `dir` returns ".". If the path consists entirely of slashes followed by non-slash bytes, Dir returns a single slash. In any other case, the returned path does not end in a slash.  

### `ext(p string) -> string`

Returns the file name extension used by `path`. The extension is the suffix beginning at the final dot in the final slash-separated element of path; it is empty if there is no dot.  

### `isAbs(p string) -> bool`

Tests whether the path is absolute.  

### `join(t tuple)`

`join` joins any number of path elements into a single path, separating them with slashes. Empty elements are ignored. The result is `clean`ed. However, if the argument list is empty or all its elements are empty, `join` returns an empty string.  

### `match(pattern string, name string) -> bool / error`

Tests whether the given name matches the pattern, where the patters is specified according to the usual shell matching patterns, e.g. `*` is a wildcard.  

### `split(p string) -> string, string`

`split` splits `path` immediately following the final slash, separating it into a directory and filename component. If there is no slash in path, Split returns "" as the directory and the filename the same as `path`. The returned values have the property that if we set `dir, filename = split(path)`, then `path == dir+filename`.  

