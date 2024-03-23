package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "Go Is Awesome!",
			want: "go-is-awesome",
		},
		{
			s:    "Tabs are all we've got",
			want: "tabs-are-all-we-ve-got",
		},
		{
			s:    "!Attention, attention!",
			want: "attention-attention",
		},
		{
			s:    "Go - Is - Awesome",
			want: "go---is---awesome",
		},
		{
			s:    "Carbon Language: An experimental successor to C++:",
			want: "carbon-language-an-experimental-successor-to-c",
		},
	}

	for _, tt := range tests {
		got := slugify(tt.s)
		if tt.want != got {
			t.Errorf("%v: got %v, want %v", tt.s, got, tt.want)
		}
	}
}
