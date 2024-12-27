package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func Logger(handler HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		handler(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(now))
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func World(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("World"))
}


func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", Logger(Hello))
	mux.HandleFunc("GET /world", Logger(World))

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Server is running at 8080")
	srv.ListenAndServe()
}