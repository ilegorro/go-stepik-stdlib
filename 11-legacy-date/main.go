package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func asLegacyDate(t time.Time) string {
	s := []byte(fmt.Sprintf("%010d", t.UnixNano()))
	head := string(s[:len(s)-9])
	tail := string(s[len(s)-9:])
	tail = strings.TrimRight(tail, "0")
	if tail == "" {
		tail = "0"
	}

	return head + "." + tail
}

func parseLegacyDate(d string) (time.Time, error) {
	re := regexp.MustCompile(`(\d+)\.(\d+)`)
	parts := re.FindStringSubmatch(d)
	if len(parts) != 3 {
		return time.Time{}, errors.New("wrong date")
	}
	sec, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, errors.New("wrong date")
	}
	nsecStr := parts[2]
	nsecStr += strings.Repeat("0", 9-len(nsecStr))
	nsec, err := strconv.Atoi(nsecStr)
	if err != nil {
		return time.Time{}, errors.New("wrong date")
	}

	return time.Unix(int64(sec), int64(nsec)), nil
}

func main() {
	leg := "3600.123456789"
	fmt.Println(parseLegacyDate(leg))

	date := time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC)
	fmt.Println(asLegacyDate(date))
}
