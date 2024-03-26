package main

import (
	"fmt"
	"io"
	"strings"
)

type AbyssWriter struct {
	n int
}

func (w *AbyssWriter) Write(p []byte) (int, error) {
	w.n += len(p)
	return len(p), nil
}

func (w *AbyssWriter) Total() int {
	return w.n
}

func NewAbyssWriter() *AbyssWriter {
	return &AbyssWriter{0}
}

func main() {
	r := strings.NewReader("go is awesome")
	w := NewAbyssWriter()
	written, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("written %d bytes\n", written)
	fmt.Println(written == int64(w.Total()))
}
