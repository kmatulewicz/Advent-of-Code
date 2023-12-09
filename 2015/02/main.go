// Description of the task: https://adventofcode.com/2015/day/2

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	presents := loadInput("input")

	totalPaper, totalRibbon := 0, 0
	for _, present := range presents {
		totalPaper += present.CountSurface() + present.CountMinSurface()
		totalRibbon += present.CountRibbon()
	}

	fmt.Printf("Elves need to order %d square feet of paper.\n", totalPaper)
	fmt.Printf("There is also %d feet of ribbon needed.\n", totalRibbon)
}

func loadInput(name string) []present {
	input, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	var presents []present

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		// last line is empty
		if len(line) == 0 {
			continue
		}
		numbers := strings.Split(line, "x")

		var present present
		present.l, err = strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		present.w, err = strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		present.h, err = strconv.Atoi(numbers[2])
		if err != nil {
			panic(err)
		}
		presents = append(presents, present)
	}
	return presents
}
