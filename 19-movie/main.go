package main

import (
	"encoding/json"
	"fmt"
)

type Genre string

func (g *Genre) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var m map[string]string
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	*g = Genre(m["name"])

	return nil
}

type Movie struct {
	Title  string  `json:"name"`
	Year   int     `json:"released_at"`
	Genres []Genre `json:"tags"`
}

func main() {
	const src = `{
		"name": "Interstellar",
		"released_at": 2014,
		"director": "Christopher Nolan",
		"tags": [
			{ "name": "Adventure" },
			{ "name": "Drama" },
			{ "name": "Science Fiction" }
		],
		"duration": "2h49m",
		"rating": "★★★★★"
	}`

	var m Movie
	err := json.Unmarshal([]byte(src), &m)
	fmt.Println(err)
	// nil
	fmt.Println(m)
	// {Interstellar 2014 [Adventure Drama Science Fiction]}
}
