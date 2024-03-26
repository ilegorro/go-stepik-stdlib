package main

import (
	"fmt"
	"io"
	"strings"
)

type TokenReader interface {
	ReadToken() (string, error)
}

type TokenWriter interface {
	WriteToken(s string) error
	Words() []string
}

type WordReader struct {
	s []string
	i int
}

func (r *WordReader) ReadToken() (string, error) {
	if r.i >= len(r.s) {
		return "", io.EOF
	}
	s := r.s[r.i]
	r.i++

	return s, nil
}

func NewWordReader(s string) TokenReader {
	return &WordReader{strings.Fields(s), 0}
}

type WordWriter struct {
	s []string
}

func (w *WordWriter) WriteToken(t string) error {
	w.s = append(w.s, t)
	return nil
}

func (w *WordWriter) Words() []string {
	return w.s
}

func NewWordWriter() TokenWriter {
	return &WordWriter{}
}

func FilterTokens(dst TokenWriter, src TokenReader, predicate func(s string) bool) (int, error) {
	n := 0
	for {
		s, err := src.ReadToken()
		if err == io.EOF {
			return n, nil
		} else if err != nil {
			return n, err
		}
		if predicate(s) {
			err = dst.WriteToken(s)
			if err != nil {
				return n, err
			}
			n += 1
		}
	}
}

func main() {
	r := NewWordReader("go is awesome")
	w := NewWordWriter()
	predicate := func(s string) bool {
		return s != "is"
	}
	n, err := FilterTokens(w, r, predicate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d tokens: %v\n", n, w.Words())
	// 2 tokens: [go awesome]
}
