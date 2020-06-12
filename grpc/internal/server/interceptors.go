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
	<<!!YOUR_CODE!!>> -- Define a logger interceptor to print the book and word from the request
}

// Measure intercepts and times a request.
func Measure(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	<<!!YOUR_CODE!!>> -- Define your measure interceptor to compute and print the call duration
}
