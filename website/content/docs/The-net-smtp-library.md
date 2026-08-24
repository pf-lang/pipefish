## Overview 

The `smtp` library implements the Simple Mail Transfer Protocol (SMTP).  

## Types 

### `Crammd_5_Auth = struct (username string, password string)`

An `Auth` that implements the CRAM-MD5 authentication mechanism as defined in RFC 2195. The returned Auth uses the given username and secret to authenticate to the server using the challenge-response mechanism.  

### `PlainAuth = struct (identity string, username string, password string, host string)`

An `Auth` that implements the PLAIN authentication mechanism as defined in RFC 4616. The returned `Auth` uses the given username and password to authenticate to host. Usually `identity` should be the empty string, to act as username.  

### `Mailer = struct (addr string, auth Auth, sender string)`

Data structure containing information about the sender of the email.  

### `Auth = abstract Crammd_5_Auth/PlainAuth`

Authorization to use a mail server.  
## Commands 

### `post from (m Mailer) to (recipient string, msg snippet)`

Emails a message to the given recipient.  

### `post from (m Mailer) to (recipients list, msg snippet) -> ok / error`

Emails a message to the given list of recipients.  

