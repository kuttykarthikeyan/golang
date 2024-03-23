package main

import (
	"log"
	"net/http"
	"time"
)


func main() {

	address := ":8080"
	mux := http.NewServeMux()
	srv := server.New()

	mux.HandleFunc("/", srv.HandleIndex)
	mux.HandleFunc("/info", srv.HandleInfo)

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("StartServer: %+v", address)
	log.Fatal(s.ListenAndServe())
}
