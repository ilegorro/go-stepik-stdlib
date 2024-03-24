package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

func incDuration(parts []string, m map[string]time.Duration) error {
	if len(parts) != 4 {
		return errors.New("неверный формат данных")
	}
	t1, err := time.Parse("15:04", parts[1])
	if err != nil {
		return err
	}
	t2, err := time.Parse("15:04", parts[2])
	if err != nil {
		return err
	}
	if t1.Compare(t2) != -1 {
		return errors.New("неверный формат данных")
	}
	title := parts[3]
	m[title] += t2.Sub(t1)

	return nil
}

func parseDate(src string) (time.Time, error) {
	return time.Parse("02.01.2006", src)
}

func parseTasks(date time.Time, lines []string) ([]Task, error) {
	m := make(map[string]time.Duration, 0)
	re := regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) (.+)`)
	for i := 0; i < len(lines); i++ {
		parts := re.FindStringSubmatch(lines[i])
		err := incDuration(parts, m)
		if err != nil {
			return nil, err
		}
	}

	tasks := make([]Task, 0)
	for k, v := range m {
		tasks = append(tasks, Task{date, v, k})
	}

	return tasks, nil
}

func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Dur > tasks[j].Dur })
}

func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	if len(lines) == 0 {
		return nil, errors.New("неверный формат данных")
	}

	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}

	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}

	sortTasks(tasks)

	return tasks, nil
}

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}
}
