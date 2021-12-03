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

type command struct {
	direction string
	value     int
}

func parseInput() []*command {
	dat, err := ioutil.ReadFile("input")
	check(err)
	input := strings.Split(string(dat), "\n")

	lines := make([]*command, len(input)-1)
	for i := range lines {
		line := input[i]
		c := &command{}
		c.direction = strings.Split(line, " ")[0]
		c.value, _ = strconv.Atoi(strings.Split(line, " ")[1])
		lines[i] = c
	}

	return lines
}

func processData(directions []*command) (int, int) {
	horPos := 0
	depthPos := 0
	aim := 0

	for i := range directions {
		d := directions[i]
		switch d.direction {
		case "forward":
			fmt.Println("Forward")
			horPos += d.value
			depthPos += (aim * d.value)
		case "down":
			fmt.Println("Down")
			aim += d.value
		case "up":
			fmt.Println("Up")
			aim -= d.value
		}
	}

	return horPos, depthPos
}

func main() {
	directions := parseInput()

	horPos, depthPos := processData(directions)
	answer := horPos * depthPos
	fmt.Println("Final numbers are " + strconv.Itoa(horPos) + " " + strconv.Itoa(depthPos))
	fmt.Println("Multipled is " + strconv.Itoa(answer))

}
