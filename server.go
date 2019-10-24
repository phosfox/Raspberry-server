package main

import (
	"fmt"
	"net/http"
)

type server struct {
	ips []string
}

func (s *server) ipHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "Parseform() err: %v", err)
			return
		}

		ip := req.FormValue("ip")
		s.ips = append(s.ips, ip)
		
		fmt.Fprintf(w, "Post: %v\n", req.PostForm)
		fmt.Fprintf(w, "IP: %s\n", ip)
		fmt.Fprintf(w, "IPs: %v\n", s.ips)
	default:
		fmt.Fprintf(w, "Not implemented yet\n")
	}
}

func (s *server) start(){
	http.HandleFunc("/ip", s.ipHandle)

	http.ListenAndServe(":8090", nil)
}

func main() {
	var s server
	s.start()
}
