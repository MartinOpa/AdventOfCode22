package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func getInput() [][]int {
	fname := "inputDay1.txt"
	dir := "/"
	/*dir, err := filepath.Abs(filepath.Dir("inputs/"))
	if err != nil {
		panic(err)
	}*/

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		panic(err)
	}

	inputString := string(b)

	rawLines := strings.Split(inputString, "\n\n")
	lines := make([][]int, len(rawLines))
	for i, rawLine := range rawLines {
		rawNums := strings.Split(rawLine, "\n")
		nums := []int{}
		for _, rawNum := range rawNums {
			if rawNum == "" {
				continue
			}
			num, err := strconv.Atoi(rawNum)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		lines[i] = nums
	}

	return lines
}

func main() {
	input := getInput()
	part1(input)
	part2(input)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func part1(input [][]int) {
	m := -1
	for _, carrying := range input {
		sum := 0
		for _, calories := range carrying {
			sum += calories
		}
		m = max(m, sum)
	}
	fmt.Println(m)
}

func part2(input [][]int) {
	res := 0
	sums := []int{}
	for _, carrying := range input {
		sum := 0
		for _, calories := range carrying {
			sum += calories
		}
		sums = append(sums, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	for i := 0; i < 3; i++ {
		res += sums[i]
	}

	fmt.Println(res)
}
