package main

import (
	"flag"
	"os"
)

type format int

const (
	normal    format = iota
	bold             = iota
	maxformat        = iota
)

var formats [maxformat]string

func init() {
	formats[normal] = "%s%s"
	// ANSI color code that works with both dark and bright background
	// 1 - bold
	// 4 - underline
	formats[bold] = "%s\033[1;4m%s\033[0m"
}

func main() {
	bufSize := flag.Int("numLines", 100, "the capacity of the buffer")
	flag.Parse()

	buf := lineBuffer{
		wordCounts: make(map[string]int),
		lines:      make([]*line, *bufSize),
	}

	if err := buf.highlight(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
