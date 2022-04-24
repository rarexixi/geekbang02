package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"
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

	getClientIP := func(r *http.Request) string {
		xForwardedFor := r.Header.Get("X-Forwarded-For")
		ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
		if ip != "" {
			return ip
		}
		ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
		if ip != "" {
			return ip
		}
		if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
			return ip
		}
		return ""
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		copyHeader(&rw, r)
		log.Printf("client host: %s\n", getClientIP(r))
		log.Printf("return http code: %d\n", http.StatusOK)
	})

	http.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		copyHeader(&rw, r)
		rw.WriteHeader(http.StatusOK)
	})

	server.ListenAndServe()
}

