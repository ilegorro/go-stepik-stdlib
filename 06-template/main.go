package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

var templateText = `{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽.` +
	` {{if ge .Balance 100 -}} Все в порядке.` +
	`{{- else if gt .Balance 0 -}} Пора пополнить.` +
	`{{- else -}} Доступ заблокирован.{{- end}}`

type User struct {
	Name    string
	Balance int
}

func Say(u User) string {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))
	var res bytes.Buffer
	err := tpl.Execute(&res, u)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	return res.String()
}

func main() {
	user := User{"Алиса", 500}
	fmt.Println(Say(user))
}
