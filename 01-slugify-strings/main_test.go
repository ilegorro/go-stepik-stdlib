package main

import "testing"

func TestSlugify(t *testing.T) {
	const data = "Go Is Awesome"
	const want = "go-is-awesome"
	got := slugify(data)
	if got != want {
		t.Errorf("%s: got: %#v, want: %#v", data, got, want)
	}
}
