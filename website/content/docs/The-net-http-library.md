## Overview 

The `http` library supplies an HTTP client implementation. 

Acting as a server is built into Pipefish services and there is no good reason to duplicate this functionality.  

## Modules 

### ` ("io")`

### ` ("net/http")`

### ` ("strings")`
## Types 

### `Http = struct (url string?, contentType string?)`

Data structure saying where to post and what to post to.  

### `Response = struct (body string)`

Data type containing an HTTP response.  
## Commands 

### `post to (h Http, body snippet)`

Posts `body` to the url specified in the `url` field of `h`.  

### `post to (h Http, body any)`

Posts `body` to the url specified in the `url` field of `h`.  

### `get(x ref) from (h Http)`

Gets an HTTP response from the url specified in the `url` field of `h`.  
## Functions 

### `Http(url string)`

Overloads the `Http` constructor to default to a `NULL` value for the `contentType` field when only the url is supplied.  

