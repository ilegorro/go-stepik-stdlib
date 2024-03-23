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
	}

	for _, tt := range tests {
		got := slugify(tt.s)
		if tt.want != got {
			t.Errorf("%v: got %v, want %v", tt.s, got, tt.want)
		}
	}
}
