// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package server

import (
	"bufio"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherland/target_labs/grpc/internal/generated"
	"github.com/gopherland/target_labs/grpc/internal/grep"
)

type Grepper struct {
	generated.UnimplementedGrepperServer
	assets string
}

func NewGrepper(dir string) *Grepper {
	return &Grepper{assets: dir}
}

// Grep counts occurrences of a given word in a book.
func (g *Grepper) Grep(ctx context.Context, in *generated.BookInfo) (*generated.Occurrences, error) {
	<<!!YOUR_CODE!!>> -- using the count method below, compute the grep count and issue a response.
}

func (g *Grepper) count(book, word string) (int64, error) {
	if len(book) == 0 || len(word) == 0 {
		return 0, errors.New("you must specify a book name and a word")
	}

	file, err := os.Open(filepath.Join(g.assets, book+".txt"))
	if err != nil {
		return 0, err
	}

	var count int64
	scanner := bufio.NewScanner(file)
	w := strings.ToLower(word)
	for scanner.Scan() {
		count += grep.WordCount(w, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
