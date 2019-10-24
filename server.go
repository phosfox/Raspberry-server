package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type server struct {
	db  *sql.DB
}

func (s *server) ipHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "Parseform() err: %v", err)
			return
		}

		ip := req.FormValue("ip")

		InsertIP(s.db, ip)

		fmt.Fprintf(w, "IP: %s\n", ip)
	default:
		fmt.Fprintf(w, "Not implemented yet\n")
	}
}

func (s *server) start() {
	http.HandleFunc("/ip", s.ipHandle)

	http.ListenAndServe(":8090", nil)
}

func main() {
	var s server
	db := GetDB()
	s.db = db

	rows := SelectAllIps(db)
	var id int
	var ip string
	var timestamp time.Time

	for rows.Next() {
		rows.Scan(&id, &ip, &timestamp)
		fmt.Printf("Id:%d IP:%s Timestamp:%s\n", id, ip, timestamp.Local())
	}

	s.start()
}
