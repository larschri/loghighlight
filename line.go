package main

import (
	"fmt"
	"strings"
	"unicode"
)

type tokenType int

const (
	space  tokenType = iota
	word             = iota
	wordN            = iota
	number           = iota
	other            = iota
)

type token struct {
	pos  int
	end  int
	typ  tokenType
	fmat format
}

type line struct {
	line   string
	tokens []token
}

func (l *line) String() string {
	var bld strings.Builder
	bs := []byte(l.line)
	var i int
	for _, t := range l.tokens {
		_, _ = fmt.Fprintf(&bld, formats[t.fmat], bs[i:t.pos], bs[t.pos:t.end])
		i = t.end
	}
	_, _ = fmt.Fprintf(&bld, "%s", bs[i:])

	return bld.String()
}

func tokenize(s string) line {
	var tokens []token
	var t token

	for i, c := range s {
		if unicode.IsLetter(c) {
			switch t.typ {
			case other:
				t.end = i
				tokens = append(tokens, t)
				fallthrough
			case space:
				t.typ = word
				t.pos = i
			case word:
				// NOP
			case wordN:
				// NOP
			case number:
				t.typ = wordN
			}
		} else if unicode.IsNumber(c) {
			switch t.typ {
			case other:
				t.end = i
				tokens = append(tokens, t)
				fallthrough
			case space:
				t.typ = number
				t.pos = i
			case word:
				t.typ = wordN
			case wordN:
				// NOP
			case number:
				// NOP
			}
		} else if unicode.IsSpace(c) {
			switch t.typ {
			case space:
				// NOP
			case other:
				fallthrough
			case word:
				fallthrough
			case wordN:
				fallthrough
			case number:
				t.end = i
				tokens = append(tokens, t)
			}
			t.typ = space
		} else {
			switch t.typ {
			case other:
				// NOP
			case word:
				fallthrough
			case wordN:
				fallthrough
			case number:
				t.end = i
				tokens = append(tokens, t)
				fallthrough
			case space:
				t.typ = other
				t.pos = i
			}
		}
	}
	return line{
		line:   s,
		tokens: tokens,
	}
}
