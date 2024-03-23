package main

import (
	"fmt"
	"time"
)

func isLeapYear(year int) bool {
	t := time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC)

	return t.YearDay() == 366
}

func main() {
	fmt.Println(isLeapYear(2020))
	fmt.Println(isLeapYear(2022))
}
