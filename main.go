package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	copyHeader := func(rw *http.ResponseWriter, r *http.Request) {
		for key, val := range r.Header {
			for _, item := range val {
				(*rw).Header().Add(key, item)
			}
		}
		(*rw).Header().Set("VERSION", os.Getenv("VERSION"))
	}

	http.HandleFunc("/req_header", func(rw http.ResponseWriter, r *http.Request) {
		copyHeader(&rw, r)
		fmt.Printf("client host: %s\n", r.RemoteAddr)
		fmt.Printf("return http code: %d\n", http.StatusOK)
	})

	http.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		copyHeader(&rw, r)
		rw.WriteHeader(http.StatusOK)
	})

	server.ListenAndServe()
}

