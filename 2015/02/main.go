// Description of the task: https://adventofcode.com/2015/day/2

package main

import (
	"fmt"
	"helpers"
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

	var presents []present

	lines := helpers.LoadLines(name)
	for _, line := range lines {
		numbers := strings.Split(line, "x")

		var present present
		var err error
		present.l, err = strconv.Atoi(numbers[0])
		if err != nil {
			panic(err.Error())
		}
		present.w, err = strconv.Atoi(numbers[1])
		if err != nil {
			panic(err.Error())
		}
		present.h, err = strconv.Atoi(numbers[2])
		if err != nil {
			panic(err.Error())
		}
		presents = append(presents, present)
	}
	return presents
}
