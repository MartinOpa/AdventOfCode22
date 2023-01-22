package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func getInput() string {
	fname := "inputDay6.txt"

	dir, err := filepath.Abs(filepath.Dir("day6/"))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		panic(err)
	}

	return string(b)
}

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}

func part1(input string) {
	runes := []rune(input)

	for i := 4; i < len(runes); i++ {
		if unique(string(runes[i-4 : i])) {
			fmt.Println(i)
			break
		}
	}
}

func part2(input string) {
	runes := []rune(input)

	for i := 14; i < len(runes); i++ {
		if unique(string(runes[i-14 : i])) {
			fmt.Println(i)
			break
		}
	}
}

func main() {
	input := getInput()
	part1(input)
	part2(input)
}
