package main

import (
	"errors"
	"fmt"
	"time"
)

type TimeOfDay struct {
	hour   int
	minute int
	second int
	tz     *time.Location
}

func (t *TimeOfDay) Hour() int {
	return t.hour
}

func (t *TimeOfDay) Minute() int {
	return t.minute
}

func (t *TimeOfDay) Second() int {
	return t.second
}

func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d %v",
		t.hour, t.minute, t.second, t.tz.String())
}

func (t *TimeOfDay) Equal(other TimeOfDay) bool {
	return t.tz.String() == other.tz.String() &&
		t.hour == other.hour &&
		t.minute == other.minute &&
		t.second == other.second
}

func (t *TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.tz.String() != other.tz.String() {
		return false, errors.New("wrong timezone")
	}
	tsec := t.second + t.minute*60 + t.hour*60*60
	osec := other.second + other.minute*60 + other.hour*60*60

	return tsec < osec, nil
}

func (t *TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.tz.String() != other.tz.String() {
		return false, errors.New("wrong timezone")
	}
	tsec := t.second + t.minute*60 + t.hour*60*60
	osec := other.second + other.minute*60 + other.hour*60*60

	return tsec > osec, nil
}

func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	return TimeOfDay{hour, min, sec, loc}
}

func main() {
	t1 := MakeTimeOfDay(17, 45, 22, time.UTC)
	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

	fmt.Println(t1.Hour(), t1.Minute(), t1.Second())
	// 17 45 22

	fmt.Println(t1)
	// 17:45:22 UTC

	fmt.Println(t1.Equal(t2))
	// false

	before, err := t1.Before(t2)
	fmt.Println(before, err)
	// true <nil>

	after, err := t1.After(t2)
	fmt.Println(after, err)
	// false <nil>
}
