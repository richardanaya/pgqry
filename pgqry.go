package pgqry

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// CreateConnectionString makes a connection string
func CreateConnectionString(host string, port int, user, password, dbname string) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}

// Execute executes sql a database with a connection string
func Execute(connectionString, query string) (sql.Result, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return result, nil
}

// Query executes sql a database with a connection string
func Query(connectionString, query string) (*sql.Rows, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	result, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return result, nil
}
