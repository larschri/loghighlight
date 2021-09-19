package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestHighlight(t *testing.T) {
	buf := lineBuffer{
		wordCounts: make(map[string]int),
		lines:      make([]*line, 2),
	}

	r := strings.NewReader(`2021-09-19T18:05:44Z foo bar
2021-09-19T18:05:45Z bar
2021-09-19T18:05:45Z foobar
2021-09-19T18:05:45Z foo bar
`)
	w := new(bytes.Buffer)

	if err := buf.highlight(r, w); err != nil {
		t.Fatal(err)
	}

	expected := `2021-09-19T18:05:44Z [[foo]] [[bar]]
2021-09-19T18:05:45Z bar
2021-09-19T18:05:45Z [[foobar]]
2021-09-19T18:05:45Z [[foo]] bar
`

	expected = strings.ReplaceAll(expected, "[[", "\033[1;4m")
	expected = strings.ReplaceAll(expected, "]]", "\033[0m")

	if w.String() != expected {
		t.Errorf("got %s, expected %v", w.String(), expected)
	}
}
