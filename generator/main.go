// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"fmt"
	"log"

	"github.com/gookit/color"
	<<!!YOUR_CODE!!>> -- set up your stacks import
)

<<!!YOUR_CODE!!>> Setup your generate annotation for float64 and int32
func main() {
	fmt.Print("\033[H\033[2J")

	tryFloat()
	tryInt()
}

func tryFloat() {
	s := stacks.Float64{}
	for _, v := range []float64{10.5, 20.2, 42.25} {
		s.Push(v)
	}
	v, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}

	cyan, green := color.FgCyan.Render, color.FgGreen.Render
	log.Printf("🥞 %-10s Pop:%v -- Top:%s -- Peek:%v", "Float64", cyan(v), green(s.Top()), &s)
}

func tryInt() {
	s := stacks.Int32{}
	for _, v := range []int32{200, 100, 300} {
		s.Push(v)
	}
	v, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}

	cyan, green := color.FgCyan.Render, color.FgGreen.Render
	log.Printf("📚 %-10s Pop:%v   -- Top:%s  -- Peek:%v", "Int32", cyan(v), green(s.Top()), &s)
}
