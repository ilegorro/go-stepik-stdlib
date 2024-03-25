package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	words := make([]string, 0)
	for scanner.Scan() {
		t := []rune(strings.ToLower(scanner.Text()))
		t[0] = []rune(strings.ToUpper(string(t[0])))[0]
		words = append(words, string(t))
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println(strings.Join(words, " "))
}
