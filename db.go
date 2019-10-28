package main

import (
	"database/sql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InsertIP(db *sql.DB, ip string, name string) {
	stmt, err := db.Prepare("INSERT INTO raspberrypis (ip, name) VALUES (?, ?)")
	checkErr(err)
	
	stmt.Exec(ip, name)
}

func createTable(db *sql.DB, s string) {
	stmt, err := db.Prepare(s)
	checkErr(err)

	stmt.Exec()
}

func SelectAllIps(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT id, ip, name, timestamp FROM raspberrypis ORDER BY timestamp DESC")
	checkErr(err)

	return rows
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./raspberrypi-ips.db")
	checkErr(err)

	createTable(db, "CREATE TABLE IF NOT EXISTS raspberrypis (id INTEGER PRIMARY KEY, ip TEXT, name TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")

	return db
}
