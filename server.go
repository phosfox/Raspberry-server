package main

	

import (
    "fmt"
    "net/http"
)


func ipHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%+v\n", req)
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post: %v\n", req.PostForm)
	fmt.Fprintf(w, "IP: %s\n", req.FormValue("ip"))
	
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/ip", ipHandle)

	http.ListenAndServe(":8090", nil)
}
