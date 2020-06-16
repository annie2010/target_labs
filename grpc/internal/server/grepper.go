// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package server

import (
	"context"
	"errors"
	"os"

	"github.com/gopherland/target_labs/grpc/internal/generated"
	"github.com/gopherland/target_labs/grpc/internal/grep"
	"github.com/rs/zerolog/log"
)

const maxBuff = 10_000

type Grepper struct {
	generated.UnimplementedGrepperServer
	assets string
}

func NewGrepper(dir string) *Grepper {
	return &Grepper{assets: dir}
}

// Grep counts occurrences of a given word in a book.
func (g *Grepper) Grep(ctx context.Context, in *generated.BookInfo) (*generated.Occurrences, error) {
	total, err := g.count(in.Book, in.Word)
	if err != nil {
		return nil, err
	}
	resp := generated.Occurrences{
		Book:  in.Book,
		Word:  in.Word,
		Total: total,
	}

	return &resp, nil
}

func (g *Grepper) count(book, word string) (int64, error) {
	if len(book) == 0 || len(word) == 0 {
		return 0, errors.New("you must specify a book name and a word")
	}

	file, err := os.Open(g.assets + "/" + book + ".txt")
	if err != nil {
		return 0, err
	}
	defer func() {
		if e := file.Close(); e != nil {
			log.Error().Err(e).Msg("closing file")
		}
	}()
	bb := make([]byte, maxBuff)
	if _, err = file.Read(bb); err != nil {
		return 0, err
	}

	return grep.WordCount([]byte(word), bb), nil
}
