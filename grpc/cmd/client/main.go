// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/gopherland/target_labs/grpc/internal/generated"
	"google.golang.org/grpc"
)

const (
	port    = "localhost:50052"
	timeOut = 500 * time.Millisecond
)

func main() {
	<<!!YOUR_CODE!!>> - Using the flag package parse the cli args: book, word
	<<!!YOUR_CODE!!>> - Establish server connection using contexts
	<<!!YOUR_CODE!!>> - Issue server side grep using your book and word

	log.Printf("Book: %s", resp.Book)
	log.Printf("Word: %s", resp.Word)
	log.Printf("Count: %d", resp.Total)
}
