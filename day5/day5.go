package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func getInput() string {
	fname := "inputDay5.txt"

	dir, err := filepath.Abs(filepath.Dir("day5/"))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		panic(err)
	}

	return string(b)
}

type step struct {
	amount, from, to int
}

func (s step) String() string {
	return fmt.Sprintf("move %d from %d to %d", s.amount, s.from, s.to)
}

func parseInput(input string) ([][]string, []step) {
	parts := strings.Split(input, "\n\n")

	boxes := parts[0]
	rawBoxes := [][]string{}

	for _, row := range strings.Split(boxes, "\n") {
		rawBoxes = append(rawBoxes, strings.Split(row, ""))
	}

	rawRows, rawCols := len(rawBoxes), len(rawBoxes[0])

	stacks := [][]string{}

	for i := 0; i < rawCols-1; i++ {
		if rawBoxes[rawRows-1][i] != " " {
			stack := []string{}
			for j := rawRows - 2; j >= 0; j-- {
				char := rawBoxes[j][i]
				if char != " " {
					stack = append(stack, char)
				}
			}
			stacks = append(stacks, stack)
		}
	}

	stepsRaw := parts[1]
	steps := []step{}
	for _, row := range strings.Split(stepsRaw, "\n") {
		curStep := step{}
		_, err := fmt.Sscanf(row, "move %d from %d to %d", &curStep.amount, &curStep.from, &curStep.to)
		if err != nil {
			panic(err)
		}
		curStep.from--
		curStep.to--
		steps = append(steps, curStep)
	}

	return stacks, steps
}

func part1(input string) {
	stacks, steps := parseInput(input)

	for _, step := range steps {
		for i := 0; i < step.amount; i++ {
			top := stacks[step.from][len(stacks[step.from])-1]
			stacks[step.to] = append(stacks[step.to], top)
			stacks[step.from] = stacks[step.from][:len(stacks[step.from])-1]
		}
	}

	ans := ""
	for _, stack := range stacks {
		ans += stack[len(stack)-1]
	}

	fmt.Println(ans)
}

func part2(input string) {
	stacks, steps := parseInput(input)

	for _, step := range steps {
		boxes := len(stacks[step.from]) - step.amount
		stacks[step.to] = append(stacks[step.to], stacks[step.from][boxes:]...)
		stacks[step.from] = stacks[step.from][:boxes]
	}

	ans := ""
	for _, stack := range stacks {
		ans += stack[len(stack)-1]
	}

	fmt.Println(ans)
}

func main() {
	input := getInput()
	part1(input)
	part2(input)
}
