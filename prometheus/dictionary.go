// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package hangman

import (
	"bufio"
	"os"
)

// WordList a collection of strings
type WordList []string

// Dictionary tracks all available words
type Dictionary struct {
	words WordList
}

// NewDictionary creates a new dictionary
func NewDictionary(file string) (*Dictionary, error) {
	wl, err := load(file)
	if err != nil {
		return nil, err
	}
	return &Dictionary{words: wl}, nil
}

// Words list all loaded dictionary words
func (d *Dictionary) Words() []string {
	return d.words
}

func load(file string) (WordList, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var wl WordList
	s := bufio.NewScanner(f)
	for s.Scan() {
		wl = append(wl, s.Text())
	}
	return wl, err
}
