package mysqltest

import (
	"database/sql"
	"fmt"
	"testing"
)

type rowQueryer interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

// RowExists checks if a given query returns any rows.
//
// The function accepts a rowQueryer interface instead of *sql.DB so that it works with sqlx as well.
//
// Source: https://snippets.aktagon.com/snippets/756-checking-if-a-row-exists-in-go-database-sql-and-sqlx-
func RowExists(t *testing.T, db rowQueryer, query string, args ...interface{}) bool {
	t.Helper()

	var exists bool

	query = fmt.Sprintf("SELECT exists (%s)", query) // nolint

	err := db.QueryRow(query, args...).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		t.Fatalf("error checking if row exists '%s' %v", args, err)
	}

	return exists
}
