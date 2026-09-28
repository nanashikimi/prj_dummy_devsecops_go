package main

import (
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/", helloHandler)
	addr := ":8080"
	log.Printf("starting server on %s", addr)
	// SAST works, now exclude this single failure
	// Configure timeouts for http.Server
	// Counterexample with ErrServerClosed check:
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

	/*   Previous ver:
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
	*/
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
	// gosec found XSS risk
	skipName := html.EscapeString(name)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("hello, " + skipName)); err != nil {
		log.Printf("write response successfully failed: %v", err)
	}
}
