// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package handler

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherland/target_labs/webservice/internal/grep"
	"github.com/gorilla/mux"
)

const assetDir = "assets"

type Response struct {
	Book        string `json:"book"`
	Word        string `json:"word"`
	Occurrences int64  `json:"count"`
}

// CountHandler returns the number of occurrence of a word in a book.
func CountHandler(w http.ResponseWriter, r *http.Request) {
	<<!!YOUR_CODE!!>> -- fetch the vars from gorilla for the book and word
	<<!!YOUR_CODE!!>> -- compute the number of occurrences given the count helper function.
	<<!!YOUR_CODE!!>> -- Return a json response and make sure the header Content-Type is set correctly!
}

// Helpers...

func count(book, word string) (int64, error) {
	if len(book) == 0 || len(word) == 0 {
		return 0, errors.New("you must specify a book name and a word")
	}

	file, err := os.Open(filepath.Join(assetDir, book+".txt"))
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
