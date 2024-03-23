package main

import "testing"

func TestPrettify(t *testing.T) {
	tests := []struct {
		m    map[string]int
		want string
	}{
		{
			m:    map[string]int{"one": 1, "two": 2, "three": 3},
			want: "{\n    one: 1,\n    three: 3,\n    two: 2,\n}",
		},
		{
			m:    map[string]int{},
			want: "{}",
		},
		{
			m:    map[string]int{"one": 1},
			want: "{ one: 1 }",
		},
	}
	for _, tt := range tests {
		got := prettify(tt.m)
		if got != tt.want {
			t.Errorf("%#v: got %v, want %v", tt.m, got, tt.want)
		}

	}
}
