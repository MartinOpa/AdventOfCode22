package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func Split(r rune) bool {
	return r == '\n' || r == '-' || r == ','
}

func getInput() []string {
	fname := "inputDay4.txt"

	dir, err := filepath.Abs(filepath.Dir("day4/"))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		panic(err)
	}

	inputString := string(b)

	rawLines := strings.FieldsFunc(inputString, Split)

	return rawLines
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func part1(input []string) {
	sum := 0

	for i := 0; i < len(input); i += 4 {
		p1a := i
		p1b := i + 1
		p2a := i + 2
		p2b := i + 3
		if (parseInt(input[p1a]) <= parseInt(input[p2a]) && parseInt(input[p2b]) <= parseInt(input[p1b])) ||
			(parseInt(input[p1a]) >= parseInt(input[p2a]) && parseInt(input[p2b]) >= parseInt(input[p1b])) {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0

	for i := 0; i < len(input); i += 4 {
		p1a := i
		p1b := i + 1
		p2a := i + 2
		p2b := i + 3
		if (parseInt(input[p2a]) <= parseInt(input[p1a]) && parseInt(input[p1a]) <= parseInt(input[p2b])) ||
			(parseInt(input[p2a]) <= parseInt(input[p1b]) && parseInt(input[p1b]) <= parseInt(input[p2b])) ||
			(parseInt(input[p1a]) <= parseInt(input[p2a]) && parseInt(input[p2a]) <= parseInt(input[p1b])) ||
			(parseInt(input[p1a]) <= parseInt(input[p2b]) && parseInt(input[p2b]) <= parseInt(input[p1b])) {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func main() {
	input := getInput()
	part1(input)
	part2(input)
}
