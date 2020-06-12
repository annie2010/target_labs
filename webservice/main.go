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
	<<!!YOUR_CODE!!>> -- create a gorilla mux and define your route /v1/wc/book/word
	<<!!YOUR_CODE!!>> -- Define your logging middleware
  <<!!YOUR_CODE!!>> -- Initialize your web server using the constant above
}
