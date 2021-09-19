package main

import (
	"bufio"
	"fmt"
	"io"
)

type lineBuffer struct {
	// wordCounts contains counts for every word
	// TODO: consider map of []*line to highlight phrases
	wordCounts map[string]int
	i          int
	lines      []*line
}

func (b *lineBuffer) push(l *line) {
	// Format the line, and populate b.wordCount
	bs := []byte(l.line)
	for i, t := range l.tokens {
		if t.typ == word {
			w := string(bs[t.pos:t.end])
			b.wordCounts[w] += 1
			if b.wordCounts[w] == 1 {
				l.tokens[i].fmat = bold
			}
		}
	}

	popped := b.lines[b.i]
	b.lines[b.i] = l
	b.i = (b.i + 1) % len(b.lines)

	if popped != nil {
		// Unpopulate b.wordCount for popped line
		bs := []byte(popped.line)
		for _, t := range popped.tokens {
			if t.typ == word {
				w := string(bs[t.pos:t.end])
				if b.wordCounts[w] == 1 {
					delete(b.wordCounts, w)
				} else {
					b.wordCounts[w] -= 1
				}
			}
		}
	}
}

func (b *lineBuffer) highlight(in io.Reader, out io.Writer) error {
	bufIn := bufio.NewReader(in)
	for {
		s, err := bufIn.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("read failed: %w", err)
		}

		l := tokenize(s)

		b.push(&l)
		if _, err := out.Write([]byte(l.String())); err != nil {
			return fmt.Errorf("write failed: %w", err)
		}
	}
}
