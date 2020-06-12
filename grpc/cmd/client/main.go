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
	var book, word string
	flag.StringVar(&book, "b", "3lpigs", "Specify a book")
	flag.StringVar(&word, "w", "pig", "Specify a word")
	flag.Parse()

	// Set up a connection to the server.
	log.Printf("Client Dialing %q...", port)
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	conn, err := grpc.DialContext(ctx, port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := generated.NewGrepperClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	r, err := c.Grep(ctx, &generated.BookInfo{Book: book, Word: word})
	if err != nil {
		log.Fatalf("Boom! Grep failed: %v", err)
	}
	log.Printf("Book: %s", r.Book)
	log.Printf("Word: %s", r.Word)
	log.Printf("Count: %d", r.Total)
}
