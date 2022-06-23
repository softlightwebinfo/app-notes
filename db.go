package main

import (
	"database/sql"
	"notes/helps"
	"notes/settings"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() {
	db, err := sql.Open("sqlite3", "./foo.db")
	helps.CheckErr(err)
	settings.DB = db
}
