package pf_database_2fsql

import (
	"database/sql"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"SqlDb": func(t uint32, v any) any { return v },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"SqlDb": (**sql.DB)(nil),
}

func GoMakeSqlConnection(driver string, connectionString string) any {
	var driversMap = map[string]string{"COCKROACHDB": "postgres", "FIREBIRD_SQL": "firebirdsql", "MARIADB": "mysql", "MICROSOFT_SQL_SERVER": "sqlserver", "MYSQL": "mysql",
		"ORACLE": "oracle", "POSTGRESQL": "postgres", "SNOWFLAKE": "snowflake", "SQLITE": "sqlite", "TIDB": "mysql"}
	sqlObj, connectionError := sql.Open(driversMap[driver], connectionString)
	if connectionError != nil {
		return connectionError
	}
	pingError := sqlObj.Ping()
	if pingError != nil {
		return pingError
	}
	// If this is an in-memory Sqlite DB, force a single open connection
	// so the same connection (and thus same in-memory DB) is reused.
	if driversMap[driver] == "sqlite" && connectionString == ":memory:" {
		sqlObj.SetMaxOpenConns(1)
		sqlObj.SetMaxIdleConns(1)
	}
	return sqlObj
}
