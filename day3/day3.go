package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

func getInput() []string {
	fname := "inputDay3.txt"

	dir, err := filepath.Abs(filepath.Dir("day3/"))
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

func priority(item rune) uint64 {
	if item >= 'a' && item <= 'z' {
		return uint64(item-'a') + 1
	}
	return uint64(item-'A') + 27
}

func part1(input []string) {
	var sum uint64 = 0
	found := false

	for _, rawLine := range input {
		found = false
		compartmentA := []rune(rawLine[0:(len(rawLine) / 2)])
		compartmentB := []rune(rawLine[(len(rawLine) / 2):len(rawLine)])

		sort.Slice(compartmentA, func(i, j int) bool {
			return compartmentA[i] > compartmentA[j]
		})

		sort.Slice(compartmentB, func(i, j int) bool {
			return compartmentB[i] > compartmentB[j]
		})

		for _, item := range compartmentA {
			for _, other := range compartmentB {
				if item == other {
					sum += priority(item)
					found = true
					break
				}
			}
			if found {
				break
			}
		}

	}

	fmt.Println(sum)
}

func part2(input []string) {
	var sum uint64

	for i := 2; i < len(input); i += 3 {
		found := false
		rs1 := []rune(input[i-2])
		rs2 := []rune(input[i-1])
		rs3 := []rune(input[i])

		sort.Slice(rs1, func(j, k int) bool {
			return rs1[j] > rs1[k]
		})

		sort.Slice(rs2, func(j, k int) bool {
			return rs2[j] > rs2[k]
		})

		sort.Slice(rs3, func(j, k int) bool {
			return rs3[j] > rs3[k]
		})

		for _, item := range rs1 {
			if found {
				break
			}
			for _, other := range rs2 {
				if found {
					break
				}
				if item == other {
					for _, otherother := range rs3 {
						if item == otherother {
							sum += priority(item)
							found = true
							break
						}
					}
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
