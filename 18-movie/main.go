package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Duration time.Duration

type Rating int

func (d Duration) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, '"')
	hours := int(time.Duration(d).Hours())
	minutes := int(time.Duration(d).Minutes()) - 60*hours
	if hours > 0 {
		b = append(b, []byte(fmt.Sprintf("%dh", hours))...)
	}
	if minutes > 0 {
		b = append(b, []byte(fmt.Sprintf("%dm", minutes))...)
	}
	b = append(b, '"')
	return b, nil
}

func (r Rating) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, '"')
	s := strings.Repeat("★", int(r)) + strings.Repeat("☆", 5-int(r))
	b = append(b, []byte(s)...)
	b = append(b, '"')
	return b, nil
}

type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

func MarshalMovies(indent int, movies ...Movie) (string, error) {
	if indent == 0 {
		r, err := json.Marshal(movies)
		return string(r), err
	} else {
		r, err := json.MarshalIndent(movies, "", strings.Repeat(" ", indent))
		return string(r), err
	}
}

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	fmt.Println(s)
}
