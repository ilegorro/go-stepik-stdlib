package main

import (
	"fmt"
	"strings"
)

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
	const valid = "0123456789-abcdefghijklmnopqrstuvwxyz"
	res := strings.Map(func(i rune) rune {
		if strings.Contains(valid, string(i)) {
			return i
		}
		return ' '
	}, strings.ToLower(src))
	words := strings.Fields(res)

	return strings.Join(words, "-")
}

func main() {
	phrase := "Go Is Awesome!"
	fmt.Println(slugify(phrase))
	// go-is-awesome

	phrase = "Tabs are all we've got"
	fmt.Println(slugify(phrase))
	// tabs-are-all-we-ve-got
}
