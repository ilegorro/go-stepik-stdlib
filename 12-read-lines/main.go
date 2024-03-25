package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readLines(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	var r strings.Builder
	for _, v := range data {
		if v == '\n' {
			if r.Len() > 0 {
				lines = append(lines, r.String())
				r.Reset()
			}
		} else {
			r.WriteByte(v)
		}
	}

	return lines, nil
}

func readLines1(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			lines = append(lines, t)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func readLines2(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		s = strings.TrimRight(s, "\n")
		if s != "" {
			lines = append(lines, s)
		}
	}

	return lines, nil
}

func main() {
	lines, err := readLines2("/etc/passwd")
	if err != nil {
		panic(err)
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
