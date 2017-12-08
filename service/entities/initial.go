package entities

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath string = "./database.db"
)

var db *sql.DB

// An SQLExecutor encapsulates all functions that execute sql
// statements as an interface.
type SQLExecutor interface {
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string) (*sql.Stmt, error)
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	QueryRow(sql string, args ...interface{}) *sql.Row
}

// An DataAccessObject is a data access object containing an
// SQLExecutor interface.
type DataAccessObject struct {
	SQLExecutor
}

func init() {
	dbExist := true
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		dbExist = false
		os.OpenFile(dbPath, os.O_CREATE, os.ModePerm)
	}
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	panicIfErr(err)
	if !dbExist {
		initTables()
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initTables() {
	_, err := db.Exec(`
        CREATE TABLE users(
            id INTEGER PRIMARY KEY,
            key TEXT,
            username TEXT,
            password TEXT,
            email TEXT,
            phone TEXT
        );
        CREATE TABLE meetings(
            id INTEGER PRIMARY KEY,
            title TEXT,
            host TEXT,
            members TEXT,
            starttime TEXT,
            endtime TEXT
        );
    `)
	panicIfErr(err)
}
