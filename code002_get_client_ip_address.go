package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// obtain the client's IP address and port number by Remote Addr
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatalf("Error extracting IP: %v", err)
		}
		fmt.Fprintf(w, "Client IP: %s", ip)
	})

	// start the HTTP server and listen on port 8080
	http.ListenAndServe(":8080", nil)
}
