// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"log"
	"net/http"
	<<!!YOUR_CODE!!>> -- make sure pprof is available

	"github.com/gopherland/labs2/profiling/internal"
)

const httpPort = ":4500"

func main() {
	log.Printf("[Fib] service is listening on [%s]", httpPort)

	http.HandleFunc("/fib", internal.FibHandler)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
