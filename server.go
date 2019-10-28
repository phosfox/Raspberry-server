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
	ID        int    `json:"ID"`
	IP        string `json:"IP"`
	Name      string `json:"Name"`
	Timestamp string `json:"Timestamp"`
}

type dbrows []dbrow

func (s *server) ipHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		switch req.Header.Get("content-type") {
		case "application/json":
			var obj struct {
				IP   string `json:"IP"`
				Name string `json:"Name"`
			}
			err := json.NewDecoder(req.Body).Decode(&obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			InsertIP(s.db, obj.IP, obj.Name)

			fmt.Fprintf(w, "Added IP: %s Name: %s\n", obj.IP, obj.Name)
		case "application/x-www-form-urlencoded":
			if err := req.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ip := req.FormValue("ip")
			name := req.FormValue("name")

			InsertIP(s.db, ip, name)

			fmt.Fprintf(w, "Added IP: %s Name: %s\n", ip, name)
		default:
			fmt.Fprintf(w, "Allowed content-type: application/x-www-form-urlencoded, application/json\n")
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *server) allIpsHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		rows := SelectAllIps(s.db)
		var ips dbrows

		for rows.Next() {
			var (
				id        int
				ip        string
				name      string
				timestamp time.Time
			)
			if err := rows.Scan(&id, &ip, &name, &timestamp); err != nil {
				log.Fatal(err)
			}
			ips = append(ips, dbrow{id, ip, name, timestamp.Local().String()})
		}

		mIps, err := json.Marshal(ips)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(mIps)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *server) start() {
	fs := http.FileServer(http.Dir("frontend/dist/"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("api/ips", s.allIpsHandle)
	http.HandleFunc("api/ip", s.ipHandle)

	http.ListenAndServe(":8090", nil)
}

func main() {
	var s server
	db := GetDB()
	s.db = db
	s.start()
}
