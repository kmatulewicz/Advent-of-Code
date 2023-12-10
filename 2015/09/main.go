// Description of the task: https://adventofcode.com/2015/day/9

package main

import (
	"fmt"
	"helpers"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := helpers.LoadLines("input")

	shortest := calcDistance(lines, func(saved, current int) int { return min(saved, current) })
	fmt.Printf("Shortest distance is: %d\n", shortest)

	longest := calcDistance(lines, func(saved, current int) int { return max(saved, current) })
	fmt.Printf("Longest distance is: %d\n", longest)

}

type trip struct {
	src  string // starting point
	dst  string // destination
	dist int    // distance between starting point and destination
}

// calcDistance calculates distance in best solution based on comparison function
func calcDistance(routes []string, f func(saved, current int) int) int {
	left := []trip{}
	for _, route := range routes {
		left = append(left, routeToTrip(route)...)
	}

	// channel for sending solution after last move
	solution := make(chan []trip, 1)
	// channel for counting running goroutines
	wait := make(chan struct{}, 100000)

	//first move
	for i := 0; i < len(left); i++ {
		// add starting point to done slice
		d := []trip{{"", left[i].src, 0}}

		// remove starting point from left slice
		l := make([]trip, len(left))
		copy(l, left)
		l = removeFromTrips(l, left[i].src)

		wait <- struct{}{}
		go func(i int) {
			move(l, d, left[i], wait, solution)
			<-wait
		}(i)
	}

	distance := 0
	for {
		select {
		case tmp := <-solution:
			d := 0
			for _, e := range tmp {
				d += e.dist
			}

			// first solution or better than saved one
			if distance == 0 {
				distance = d
			} else {
				distance = f(distance, d)
			}
		}

		// all goroutines ended
		if len(wait) == 0 {
			break
		}
	}

	return distance
}

// move makes next moves recursively
func move(
	left []trip, // moves left
	done []trip, // done moves
	now trip, // current move
	wait chan struct{}, // channel for counting active goroutines
	solution chan []trip, // channel for sending solution on last move
) {

	//new done
	d := make([]trip, len(done))
	copy(d, done)
	d = append(d, now)

	//new left
	l := make([]trip, len(left))
	copy(l, left)
	l = removeFromTrips(l, now.dst)

	//no moves left
	if len(l) == 0 {
		solution <- d
		return
	}

	for i := 0; i < len(l); i++ {
		if l[i].src != now.dst {
			continue
		}
		wait <- struct{}{}
		go func(i int) {
			move(l, d, l[i], wait, solution)
			<-wait
		}(i)
	}
}

// routeToTrip translates strings with description to structs
func routeToTrip(r string) []trip {
	s := strings.Split(r, " ")
	dist, err := strconv.Atoi(s[4])
	if err != nil {
		panic(err.Error())
	}

	return []trip{
		{s[0], s[2], dist},
		{s[2], s[0], dist},
	}
}

// removeFromTrips returns trips without entries with specified destination
func removeFromTrips(trips []trip, dst string) []trip {
	t := make([]trip, len(trips))
	copy(t, trips)
	for {
		index := slices.IndexFunc(t,
			func(t trip) bool {
				if t.dst == dst {
					return true
				}
				return false
			})
		if index == -1 {
			break
		}
		t = append(t[:index], t[index+1:]...)
	}
	return t
}
