package main

import "testing"

func TestSay(t *testing.T) {
	tests := []struct {
		user User
		want string
	}{
		{
			user: User{"Алиса", 500},
			want: "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке.",
		},
		{
			user: User{"Алиса", 100},
			want: "Алиса, добрый день! Ваш баланс - 100₽. Все в порядке.",
		},
		{
			user: User{"Алиса", 75},
			want: "Алиса, добрый день! Ваш баланс - 75₽. Пора пополнить.",
		},
		{
			user: User{"Алиса", 0},
			want: "Алиса, добрый день! Ваш баланс - 0₽. Доступ заблокирован.",
		},
	}

	for _, tt := range tests {
		got := Say(tt.user)
		if got != tt.want {
			t.Errorf("\n%#v:\n got %v,\n want: %v", tt.user, got, tt.want)
		}

	}
}
