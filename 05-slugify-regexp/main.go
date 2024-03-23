package main

import (
	"fmt"
	"regexp"
	"strings"
)

func slugify(src string) string {
	re := regexp.MustCompile(`[a-z0-9-]+`)
	words := re.FindAllString(strings.ToLower(src), -1)

	return strings.Join(words, "-")
}

func main() {
	phrase := "Go Is Awesome!"
	fmt.Println(slugify(phrase))

	phrase = "Tabs are all we've got"
	fmt.Println(slugify(phrase))
}
