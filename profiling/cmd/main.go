// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gopherland/labs2/profiling/internal"
)

const httpPort = ":4500"

func main() {
	http.HandleFunc("/fib", internal.FibHandler)
	log.Printf("[Fib] service is listening on [%s]", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
