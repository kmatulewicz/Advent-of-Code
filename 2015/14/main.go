// Description of the task: https://adventofcode.com/2015/day/14

package main

import (
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

type Reindeer struct {
	name             string
	fly, rest, speed int
}

type score struct {
	name   string
	points int
}

func main() {
	lines := helpers.LoadLines("input")
	reindeers := parseInput(lines)

	time := 2503
	max := 0
	for _, r := range reindeers {
		c := calcDistance(time, r)
		if c > max {
			max = c
		}
	}
	fmt.Printf("Distance traveled by the fastest reindeer in %d seconds is: %d\n", time, max)

	scores := assignPoints(time, reindeers)
	max = 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	fmt.Printf("Points gained by the best reindeer in %d seconds is: %d\n", time, max)

}

// calcDistance calculates distance flown by a specified reindeer in specified time
func calcDistance(time int, r Reindeer) int {
	fullPeriods := int(time / (r.fly + r.rest))
	lastPeriodFly := min(time%(r.fly+r.rest), r.fly)
	flyTime := fullPeriods*r.fly + lastPeriodFly

	return flyTime * r.speed
}

func assignPoints(time int, reindeers []Reindeer) map[string]int {
	// prepare a score table
	scores := make(map[string]int)

	// calculate for each second
	for ; time > 0; time-- {
		// current distances
		c := make(map[string]int)

		max := 0
		// calculate for each reindeer
		for _, r := range reindeers {
			d := calcDistance(time, r)
			if d >= max {
				c[r.name] = d
				max = d
			}
		}

		// which reindeer traveled max distance in the round
		for k, v := range c {
			if v == max {
				// assign point
				scores[k]++
			}
		}
	}

	return scores
}

// parseInput parses the description into a slice of reindeers
func parseInput(lines []string) []Reindeer {
	reindeers := []Reindeer{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		rest, err := strconv.Atoi(parts[13])
		if err != nil {
			panic(err.Error())
		}
		fly, err := strconv.Atoi(parts[6])
		if err != nil {
			panic(err.Error())
		}
		speed, err := strconv.Atoi(parts[3])
		if err != nil {
			panic(err.Error())
		}
		r := Reindeer{parts[0], fly, rest, speed}

		reindeers = append(reindeers, r)
	}

	return reindeers
}
