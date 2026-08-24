## Overview 

The `url` library parses URLs and implements query escaping. 

See RFC 3986. This package generally follows RFC 3986, except where it deviates for compatibility reasons. RFC 6874 is followed for IPv6 zone literals.  

## Modules 

### ` ("net/url")`

### ` ("errors")`
## Types 

### `UserInfo = struct (username string, password string)`

Data structure containing the name and password of a user.  

### `Url = struct (scheme string, opaque string, user UserInfo, host string, path string, rawPath string, omitHost bool, forceQuery bool, rawQuery string, fragment string, rawFragment string)`

Data structure containing a parsed URL.  
## Functions 

### `encode(values map) -> string / error`

Encodes the `map` into “URL encoded” form ("bar=baz&foo=quux") sorted by key.  

### `escapedFragment(u Url) -> string`

Returns the escaped form of `u[fragment]`.  

### `escapedPath(u Url) -> string`

Returns the escaped form of `u[path]`.  

### `hostname(u Url) -> string`

Returns the hostname of the `Url`.  

### `isAbs(u Url) -> bool`

Tests whether the `Url` is absolute.  

### `joinPath(u Url, elem ... string) -> Url`

`joinPath` with a `Url` as the first parameter returns a new URL with the provided path elements joined to any existing path and the resulting path cleaned of any ./ or ../ elements. Any sequences of multiple / characters will be reduced to a single /. Path elements must already be in escaped form, as produced by `pathEscape`.  

### `joinPath(base string, elem ... string) -> string / error`

`joinPath` with a `string` as the first parameter returns a URL string with the provided path elements joined to the existing path of base and the resulting path cleaned of any ./ or ../ elements. Path elements must already be in escaped form, as produced by `pathEscape`.  

### `parse(rawUrl string) -> Url / error`

Parses a raw url into a `URL` structure. 

The url may be relative (a path, without a host) or absolute (starting with a scheme). Trying to parse a hostname and path without a scheme is invalid but may not necessarily return an error, due to parsing ambiguities.  

### `parseInContext(v Url, rawUrl string) -> Url / error`

Parses a `string` into a `Url` in the context of `v`. The URL may be relative or absolute.  

### `parseQuery(query string) -> map / error`

ParseQuery parses the URL-encoded query string and returns a map listing the values specified for each key. `parseQuery` always returns a non-nil map containing all the valid query parameters found. If the string is ill-formed, the function will return an `error` describing the first decoding problem encountered. 

The `query` parameter is expected to be a list of key=value settings separated by ampersands. A setting without an equals sign is interpreted as a key set to an empty value. Settings containing a non-URL-encoded semicolon are considered invalid.  

### `parseRequestUri(rawUrl string) -> Url / error`

ParseRequestURI parses a raw url into a `Url` structure. It assumes that the url was received in an HTTP request, so the url is interpreted only as an absolute URI or an absolute path. The string url is assumed not to have a #fragment suffix. (Web browsers strip #fragment before sending the URL to a web server.)  

### `pathEscape(s string) -> string`

`pathEscape` escapes the string so it can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed.  

### `pathUnescape(s string) -> string / error`

`pathUnescape` does the inverse transformation of PathEscape, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits. 

`pathUnescape` is identical to `queryUnescape` except that it does not unescape '+' to ' ' (space).  

### `port(u Url) -> string`

Returns the port part of `u[host]`, without the leading colon. 

If `u[host]` doesn't contain a valid numeric port, `port` returns an empty string.  

### `query(u Url)`

`query` parses the raw query of the `Url` and returns the corresponding values. It silently discards malformed value pairs. To check errors use `parseQuery`.  

### `queryEscape(s string) -> string`

`queryEscape` escapes the string so it can be safely placed inside a URL query.  

### `queryUnescape(s string) -> string / error`

`queryUnescape` does the inverse transformation of `queryEscape`, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.  

### `redacted(u Url) -> string`

`redacted` is like `string` but replaces any password with "xxxxx". Only the password in `u[user]` is redacted.  

### `requestUri(u Url) -> string`

Returns the encoded path?query or opaque?query string that would be used in an HTTP request for u.  

### `resolveReference(u Url, v Url) -> Url`

Resolves a URI reference to an absolute URI from an absolute base URI `u`, per RFC 3986 Section 5.2. The URI reference may be relative or absolute.  

### `string(u Url) -> string`

Reassembles the URL into a valid URL string.  

