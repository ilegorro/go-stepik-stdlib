package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func execute(cancel <-chan struct{}, fn func() int) (int, error) {
	ch := make(chan int, 1)

	go func() {
		ch <- fn()
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-cancel:
		return 0, errors.New("cancelled")
	}
}

func worker() int {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("work is done")
	return 42
}

func main() {
	cancel := make(chan struct{})

	go func() {
		time.Sleep(50 * time.Millisecond)
		if rand.Float64() < 0.5 {
			close(cancel)
		}
	}()

	res, err := execute(cancel, worker)
	fmt.Println(res, err)
}
