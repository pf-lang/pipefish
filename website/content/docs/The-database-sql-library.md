## Overview 

The `sql` library provides access to SQL databases.  

## Modules 

### ` ("database/sql")`
## Types 

### `SqlDriver = enum COCKROACHDB, FIREBIRD_SQL, MARIADB, MICROSOFT_SQL_SERVER, MYSQL, ORACLE, POSTGRESQL, SNOWFLAKE, SQLITE, TIDB`

The available drivers for creating a SQL connection.  

### `Varchar{i int} = clone string`

A type templace corresponding to SQL's `VARCHAR(i)`, and automatically converted to and from it.  

### `SqlDb = wrapper *DB`

A connection to a SQL database. This is a non-standard type, a wrapper around Golang's `*sql.DB` type, and does not implement `==`, `!=`, or serialization.  
## Commands 

### `get(x ref) as (t type) from (y SqlDb, query snippet)`

Populates the reference variable with the result of the SQL query represented as a value of the given type.  

### `get(x ref) like (t type) from (y SqlDb, query snippet)`

Populates the reference variable with the result of the SQL query represented as a value like that of the given type, e.g. if the type is `list{MyStructType}` then the value returned will be of type `list`, with its elements of type `StructType`.  

### `get(r ref) from (y SqlDb, query snippet)`

Populates the reference variable with the result of the SQL query represented as a value of type `list` with elements of type `map`.  

### `ping(db SqlDb)`

Pings the database to test the connection.  

### `post to (sql SqlDb, command snippet)`

Posts the given SQL command to the SQL database.  
## Functions 

### `SqlDb(driver SqlDriver, hostpath string, port int, hostname string, username string, password string) -> SqlDb / error`

Returns a SQL connection.  

### `SqlDb(driver SqlDriver)`

Supplies a transient in-memory version of the given database. Currently only works for `SQLITE`, until someone writes another one.  

