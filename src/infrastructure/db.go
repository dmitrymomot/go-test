package infrastructure

import (
	"database/sql"
	"go-test/src/interfaces/repositories"

	_ "github.com/go-sql-driver/mysql" // init mysql driver
)

// NewMySQLHandler is a factory function,
// returns a new MySQLHandler struct instance
func NewMySQLHandler(dbSourcePath string) *MySQLHandler {
	conn, err := sql.Open("mysql", dbSourcePath)
	if err != nil {
		panic(err)
	}
	if err = conn.Ping(); err != nil {
		panic(err.Error())
	}
	return &MySQLHandler{conn}
}

// MySQLHandler struct
type MySQLHandler struct {
	conn *sql.DB
}

// Execute a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (h *MySQLHandler) Execute(query string, args ...interface{}) (repositories.DbResult, error) {
	smth, err := h.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	return smth.Exec(args...)
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (h *MySQLHandler) Query(query string, args ...interface{}) (repositories.DbRow, error) {
	rows, err := h.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &MySQLRow{rows}, nil
}

// Close closes the database, releasing any open resources.
func (h *MySQLHandler) Close() error {
	return h.conn.Close()
}

// MySQLRow structure
type MySQLRow struct {
	rows *sql.Rows
}

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in Rows.
func (r *MySQLRow) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}

// Next prepares the next result row for reading with the Scan method
func (r *MySQLRow) Next() bool {
	return r.rows.Next()
}
