package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "sqlite-database.db")
	if err != nil {
		panic("pau na migration")
	}
	db.Exec("CREATE TABLE exchanges_rates ( id            INTEGER PRIMARY KEY, code          TEXT, code_un       TEXT, name          TEXT, high          TEXT, low           TEXT, var_bid       TEXT, pct_change    TEXT, bid           TEXT, ask           TEXT, timestamp     TEXT, create_date   TEXT );")
}
