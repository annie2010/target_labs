// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gopherland/target_labs/webservice/internal/grep"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const (
	assetDir = "assets"
	maxBuff  = 10_000
)

type Response struct {
	Book        string `json:"book"`
	Word        string `json:"word"`
	Occurrences int64  `json:"count"`
}

type Book struct{}

func NewBook() Book {
	return Book{}
}

// Count returns the number of occurrence of a word in a book.
func (b Book) Count(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	count, err := count(params["book"], params["word"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	res := Response{
		Book:        params["book"],
		Word:        params["word"],
		Occurrences: count,
	}
	raw, err := json.Marshal(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(raw); err != nil {
		http.Error(w, "unable to write response", http.StatusInternalServerError)
	}
}

// Helpers...

func count(book, word string) (int64, error) {
	if len(book) == 0 || len(word) == 0 {
		return 0, errors.New("you must specify a book name and a word")
	}

	file, err := os.Open(assetDir + "/" + book + ".txt")
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
