package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput() []int {
	dat, err := ioutil.ReadFile("day1.input")
	check(err)
	input := strings.Split(string(dat), "\n")

	lines := make([]int, len(input)-1)
	for i := range lines {
		lines[i], _ = strconv.Atoi(input[i])
	}

	return lines
}

func main() {
	depths := parseInput()

	increases := 0
	previousWindow := depths[0] + depths[1] + depths[2]
	for i := 2; i < len(depths)-1; i++ {
		currentWindow := depths[i-1] + depths[i] + depths[i+1]
		if currentWindow > previousWindow {
			increases++
		}
		previousWindow = currentWindow
	}
	fmt.Println(len(depths))
	fmt.Println(increases)
}
