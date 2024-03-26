package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
)

type Reader struct {
	n int
}

func (r *Reader) Read(b []byte) (int, error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
	if len(b) > r.n {
		b = b[:r.n]
	}
	k, err := rand.Read(b)
	r.n -= k

	return k, err
}

func RandomReader(max int) io.Reader {
	return &Reader{max}
}

func main() {
	rnd := RandomReader(10)
	// rd := bufio.NewReader(rnd)
	rd := bufio.NewReaderSize(rnd, 3)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
}
