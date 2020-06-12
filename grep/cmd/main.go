// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gopherland/target_labs/grep"
)

func main() {
	var word, fileName string
	flag.StringVar(&word, "w", "moby", "Specify the word to find")
	flag.StringVar(&fileName, "f", "assets/moby.txt", "Specify a text file to grep from")
	flag.Parse()

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open file %s", fileName)
	}

	var count int64
	scanner := bufio.NewScanner(file)
	w := strings.ToLower(word)
	for scanner.Scan() {
		occ := grep.WordCount(w, scanner.Text())
		count += occ
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner failed %v", err)
	}

	fmt.Printf("Found %d occurrences of %q\n", count, word)
}
