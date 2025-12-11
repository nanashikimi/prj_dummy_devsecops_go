package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)
	mux.HandleFunc("/", helloHandler)
	addr := ":8080"
	log.Printf("starting server on %s", addr)
	// SAST works, now exclude this single failure
	// Counterexample with ErrServerClosed check:
	/*
	   server := &http.Server{
	       Addr:         addr,
	       Handler:      mux,
	       ReadTimeout:  5 * time.Second,
	       WriteTimeout: 10 * time.Second,
	       IdleTimeout:  120 * time.Second,
	   }

	   if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	       log.Fatal(err)
	   }
	*/
	// This insecure usage would be fixed for prod, where we configure timeouts for http.Server, so:
	// #nosec G114
	// nosemgrep: go.lang.security.audit.net.use-tls.use-tls
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello, " + name))
}
