package main

import (
	"database/sql"
	"fmt"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InsertIP(db *sql.DB, ip string) {
	stmt, err := db.Prepare("INSERT INTO raspberrypis (ip) VALUES (?)")
	checkErr(err)

	stmt.Exec(ip)
}

func createTable(db *sql.DB, s string) {
	stmt, err := db.Prepare(s)
	checkErr(err)

	stmt.Exec()
}

func SelectAllIps(db *sql.DB) *sql.Rows{
	rows, err := db.Query("SELECT id, ip, timestamp FROM raspberrypis")
	checkErr(err)

	return rows
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./raspberrypi-ips.db")
	checkErr(err)

	createTable(db, "CREATE TABLE IF NOT EXISTS raspberrypis (id INTEGER PRIMARY KEY, ip TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")

	InsertIP(db, "9.9.9.9")

	rows := SelectAllIps(db)

	var id int
	var ip string
	var timestamp time.Time

	for rows.Next() {
		rows.Scan(&id, &ip, &timestamp)
		fmt.Printf("Id:%d IP:%s Timestamp:%s\n", id, ip, timestamp.Local())
	}

	return db
}
