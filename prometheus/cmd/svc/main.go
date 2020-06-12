// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gopherland/target_labs/prometheus/hangman"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const port = ":5000"

func main() {
	r := mux.NewRouter()
	m := handlers.LoggingHandler(os.Stdout, r)

	h, err := hangman.NewHandler("assets/words.txt", hangman.NewMetrics())
	if err != nil {
		panic(err)
	}
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	r.HandleFunc("/api/v1/new_game", h.NewGame).Methods("GET")
	r.HandleFunc("/api/v1/guess", h.Guess).Methods("POST")

	log.Printf("Hangman Listening on port %s...\n", port)
	log.Panic(http.ListenAndServe(port, m))
}
