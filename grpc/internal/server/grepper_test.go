// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package server_test

import (
	"context"
	"testing"

	"github.com/gopherland/target_labs/grpc/internal/generated"
	"github.com/gopherland/target_labs/grpc/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestGrepper(t *testing.T) {
	svc := server.NewGrepper("testdata")

	resp, err := svc.Grep(context.Background(), &generated.BookInfo{Book: "fred", Word: "duh"})
	assert.Nil(t, err)
	assert.Equal(t, int64(3), resp.Total)
}

func BenchmarkGrepper(b *testing.B) {
	svc := server.NewGrepper("testdata")

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		svc.Grep(context.Background(), &generated.BookInfo{Book: "fred", Word: "duh"})
	}
}
