package main

import (
	"fmt"
	"sort"
	"strings"
)

func prettify(m map[string]int) string {
	if len(m) == 0 {
		return "{}"
	}

	str := make([]string, 0)
	for k, v := range m {
		str = append(str, fmt.Sprintf("%v: %d", k, v))
	}
	if len(str) == 1 {
		return fmt.Sprintf("{ %v }", str[0])
	}

	sort.Strings(str)
	var r strings.Builder
	r.WriteString("{\n")
	for _, s := range str {
		r.WriteString("    ")
		r.WriteString(s)
		r.WriteString(",\n")
	}
	r.WriteString("}")

	return r.String()
}

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(prettify(m))
}
