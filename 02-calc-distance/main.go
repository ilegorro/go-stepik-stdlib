package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getDistValue(word string, unit string) *float64 {
	num, found := strings.CutSuffix(word, unit)
	v, err := strconv.ParseFloat(num, 64)
	if found && err == nil {
		return &v
	}

	return nil
}

func calcDistance(directions []string) int {
	res := 0
	for _, str := range directions {
		for _, word := range strings.Fields(str) {
			v := getDistValue(word, "km")
			if v != nil {
				res += int(1000 * *v)
			}
			v = getDistValue(word, "m")
			if v != nil {
				res += int(*v)
			}
		}
	}

	return res
}

func main() {
	directions := []string{
		"100m to intersection",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
	}
	res := calcDistance(directions)

	fmt.Println(res)
}
