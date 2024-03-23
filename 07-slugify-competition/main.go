package main

import (
	"fmt"
	"strings"
)

func slugify(src string) string {
	var r strings.Builder
	r.Grow(len(src))
	e := false
	for i := 0; i < len(src); i++ {
		if (src[i] >= '0' && src[i] <= '9') ||
			(src[i] >= 'a' && src[i] <= 'z') || src[i] == '-' {
			if e && r.Len() > 0 {
				r.WriteByte('-')
			}
			r.WriteByte(src[i])
			e = false
		} else if src[i] >= 'A' && src[i] <= 'Z' {
			if e && r.Len() > 0 {
				r.WriteByte('-')
			}
			r.WriteByte(src[i] + ' ')
			e = false
		} else {
			e = true
		}
	}

	return r.String()
}

func main() {
	phrase := "Go Is Awesome!"
	fmt.Println(slugify(phrase))

	phrase = "Tabs are all we've got"
	fmt.Println(slugify(phrase))
}
