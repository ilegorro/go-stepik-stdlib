package main

import (
	"fmt"
	"strconv"
	"strings"
)

func formatList1(items []string) string {
	str := ""
	for idx, v := range items {
		str += fmt.Sprintf("%d) %v\n", idx+1, v)
	}

	return str
}

func formatList2(items []string) string {
	str := make([]string, len(items))
	for idx, v := range items {
		str[idx] = fmt.Sprintf("%d) %v", idx+1, v)
	}

	return strings.Join(str, "\n")
}

func formatList3(items []string) string {
	var r strings.Builder
	r.Grow(len(items) * 4)

	for idx, v := range items {
		r.WriteString(strconv.Itoa(idx + 1))
		r.WriteString(") ")
		r.WriteString(v)
		r.WriteRune('\n')
	}

	return r.String()
}

func main() {
	list := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
	}
	s := formatList3(list)
	fmt.Print(s)
}
