// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gopherland/target_labs/webservice/internal/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	port           = ":5000"
	defaultTimeout = 10 * time.Second
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc(`/api/v1/grep/{book:\w+}/{word:\w+}`, handler.CountHandler)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
	})

	svc := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: defaultTimeout,
		ReadTimeout:  defaultTimeout,
	}
	log.Printf("[BookGrep] Service listening on port %s", port)
	log.Panic(svc.ListenAndServe())
}
