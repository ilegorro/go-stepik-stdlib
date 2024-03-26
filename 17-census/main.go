package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct {
	files map[string]*os.File
	n     int
}

func (c *Census) Count() int {
	return c.n
}

func (c *Census) Add(name string) {
	var err error
	path := filepath.Join("census", fmt.Sprintf("%v.txt", string(name[0])))
	f, ok := c.files[path]
	if !ok {
		f, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			panic(err)
		}
		c.files[path] = f
	}
	_, err = fmt.Fprintf(f, "%v\n", name)
	if err != nil {
		panic(err)
	}
	c.n++
}

func (c *Census) Close() {
	for _, v := range c.files {
		v.Close()
	}
}

func NewCensus() *Census {
	census := Census{make(map[string]*os.File, 0), 0}
	err := os.Mkdir("census", 0o755)
	if err != nil {
		panic(err)
	}
	for _, v := range alphabet {
		path := filepath.Join("census", fmt.Sprintf("%v.txt", string(v)))
		err = os.WriteFile(path, []byte{}, 0o644)
		if err != nil {
			panic(err)
		}
	}
	return &census
}

func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	rand.Seed(0)
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}
