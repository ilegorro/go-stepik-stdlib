package main

import "testing"

var result string // prevents compile optimization

var items = []string{
	"go is awesome",
	"cats are cute",
	"rain is wet",
	"fire is hot",
	"air is transparent",
	"violets are blue",
	"apple is fruit",
	"tennis is life",
	"flowers are beautiful",
	"coding is passion",
}

func BenchmarkFormatList1(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = formatList1(items)
	}

	result = s
}

func BenchmarkFormatList2(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		formatList2(items)
	}

	result = s
}

func BenchmarkFormatList3(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		formatList3(items)
	}

	result = s
}
