// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package server

import (
	"context"
	"log"
	"time"

	"github.com/gopherland/target_labs/grpc/internal/generated"
	"google.golang.org/grpc"
)

// Logger intercepts and logs a request.
func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info, ok := req.(*generated.BookInfo); ok {
		log.Printf("🧮-> Greping for %s//%s...", info.Book, info.Word)
	}

	return handler(ctx, req)
}

// Measure intercepts and times a request.
func Measure(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func(t time.Time) {
		log.Printf("Elapsed %v", time.Since(t))
	}(time.Now())

	return handler(ctx, req)
}
