package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func getInput() []string {
	fname := "inputDay2.txt"

	dir, err := filepath.Abs(filepath.Dir("day2/"))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		panic(err)
	}

	inputString := string(b)

	rawLines := strings.Split(inputString, "\n")

	return rawLines
}

func part1(input []string) {
	sum := 0

	for _, rawLine := range input {
		match := strings.Split(rawLine, " ")
		oMove := match[0]
		myMove := match[1]
		switch oMove {
		case "A":
			{
				switch myMove {
				case "X":
					sum += 3 + 1
				case "Y":
					sum += 6 + 2
				case "Z":
					sum += 0 + 3
				}
			}
		case "B":
			{
				switch myMove {
				case "X":
					sum += 0 + 1
				case "Y":
					sum += 3 + 2
				case "Z":
					sum += 6 + 3
				}
			}
		case "C":
			{
				switch myMove {
				case "X":
					sum += 6 + 1
				case "Y":
					sum += 0 + 2
				case "Z":
					sum += 3 + 3
				}
			}
		}
	}

	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0

	for _, rawLine := range input {
		match := strings.Split(rawLine, " ")
		oMove := match[0]
		myMove := match[1]
		switch oMove {
		case "A":
			{
				switch myMove {
				case "Z":
					sum += 6 + 2
				case "Y":
					sum += 3 + 1
				case "X":
					sum += 0 + 3
				}
			}
		case "B":
			{
				switch myMove {
				case "Z":
					sum += 6 + 3
				case "Y":
					sum += 3 + 2
				case "X":
					sum += 0 + 1
				}
			}
		case "C":
			{
				switch myMove {
				case "Z":
					sum += 6 + 1
				case "Y":
					sum += 3 + 3
				case "X":
					sum += 0 + 2
				}
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	input := getInput()
	part1(input)
	part2(input)
}
