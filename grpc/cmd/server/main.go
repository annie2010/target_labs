// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"log"
	"net"

	"github.com/gopherland/target_labs/grpc/internal/generated"
	"github.com/gopherland/target_labs/grpc/internal/server"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = "localhost:50052"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(server.Measure, server.Logger))
	generated.RegisterGrepperServer(s, server.NewGrepper("assets"))
	reflection.Register(s)

	log.Printf("[Grepper] Listening on %q...", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Boom! Server Failed: %v", err)
	}
}
