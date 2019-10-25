package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type server struct {
	db *sql.DB
}

type dbrow struct {
	id        int    `json:"id"`
	ip        string `json:"ip"`
	timestamp string `json:"timestamp"`
}

type dbrows []dbrow

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

func (s *server) rootHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		rows := SelectAllIps(s.db)
		ips := dbrows{dbrow{99, "255.255.255.255", "24:24"}}

		for rows.Next() {
			var (
				id        int
				ip        string
				timestamp time.Time
			)
			if err := rows.Scan(&id, &ip, &timestamp); err != nil {
				log.Fatal(err)
			}
			ips = append(ips, dbrow{id, ip, timestamp.Local().String()})
		}
		mIps, err := json.Marshal(ips)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(mIps)
	default:
		fmt.Fprintf(w, "ROOT Default\n")
	}
}

func (s *server) start() {
	http.HandleFunc("/", s.rootHandle)
	http.HandleFunc("/ip", s.ipHandle)

	http.ListenAndServe(":8090", nil)
}

func main() {
	var s server
	db := GetDB()
	s.db = db
	s.start()
}
