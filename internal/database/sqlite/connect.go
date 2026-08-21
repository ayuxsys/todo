package sqlite3

import (
	"database/sql"
	"todo/internal/utils/must"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	*sql.DB
}

func Connect(dsn string) *Connection {
	return &Connection{
		DB: must.Eval(sql.Open("sqlite3", dsn)),
	}
}

func (conn *Connection) InitTables() error {
	return conn.ExecStmts(tableStmts)
}

func Dsn(path string) string {
	return "file:" + path
}
