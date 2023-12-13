// Description of the task: https://adventofcode.com/2015/day/13

package main

import (
	"fmt"
	"helpers"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := helpers.LoadLines("input")
	guests := parseInput(input)
	fmt.Printf("Happiness for optimal arrangement is: %d\n", optimalArrangementValue(guests))

	guests = addMyself(guests)
	fmt.Printf("Happiness for optimal arrangement after adding myself is: %d\n", optimalArrangementValue(guests))
}

type Guest struct {
	name  string
	rules []Rule
}

type Rule struct {
	name   string
	change int
}

// optimalArrangementValue returns best happiness value from given guests
func optimalArrangementValue(guests []Guest) int {
	sat := []Guest{guests[0]}
	rest := guests[1:]
	ch := make(chan []Guest)
	b := block{}
	b.Add()

	sc := make([]Guest, len(sat))
	copy(sc, sat)
	rc := make([]Guest, len(rest))
	copy(rc, rest)
	go sitDown(sc, rc, ch, &b)

	// flag for defining first result
	first := false
	// best result
	best := 0

loop:
	for {
		select {
		case v := <-ch: // read solutions from goroutines
			if first {
				best = count(v)
				continue
			}
			current := count(v)
			if current > best {
				best = current
			}
		default: // check if all goroutines has been completed
			if b.Check() {
				break loop
			}
		}
	}

	return best
}

// sitDown recursively sits all guests
func sitDown(sat []Guest, rest []Guest, ch chan []Guest, b *block) {
	for i := 0; i < len(rest); i++ {
		// sit
		s := make([]Guest, 0, len(sat)+1)
		s = append(s, sat...)
		s = append(s, rest[i])

		// last guest sat
		if len(rest) == 1 {

			// return slice with ordered guests
			ch <- s
			continue
		}

		// rest
		r := make([]Guest, 0, len(rest)-1)
		r = append(r, rest[:i]...)
		r = append(r, rest[i+1:]...)

		b.Add()

		go sitDown(s, r, ch, b)

	}
	b.Sub()
}

// count counts sum of happiness
func count(guests []Guest) int {
	sum := 0
	for i := range guests {
		prev := ""
		next := ""
		switch {
		case i == 0: // first guest
			prev = guests[len(guests)-1].name
			next = guests[i+1].name
		case i == len(guests)-1: // last guest
			prev = guests[i-1].name
			next = guests[0].name
		default:
			prev = guests[i-1].name
			next = guests[i+1].name
		}

		prevRule := findRule(guests[i].rules, prev)
		nextRule := findRule(guests[i].rules, next)

		sum += prevRule.change + nextRule.change
	}

	return sum
}

// findRule return rule for given name
func findRule(r []Rule, n string) Rule {
	i := slices.IndexFunc(r, func(r Rule) bool {
		if r.name == n {
			return true
		}
		return false
	})

	return r[i]
}

// parseInput parses input into slice of guests
func parseInput(input []string) []Guest {
	guests := []Guest{}
	for _, line := range input {
		a := strings.Split(line, " ")

		// prepare a rule
		rule := Rule{name: a[10][:len(a[10])-1]}
		amount, err := strconv.Atoi(a[3])
		if err != nil {
			panic(err.Error())
		}
		if a[2] == "gain" {
			rule.change = amount
		} else {
			rule.change = -amount
		}

		// add a rule to the guest or create a new guest if it does not exist
		index := slices.IndexFunc(guests,
			func(g Guest) bool {
				if g.name == a[0] {
					return true
				}
				return false
			})
		if index != -1 {
			guests[index].rules = append(guests[index].rules, rule)
		} else {
			guests = append(guests, Guest{a[0], []Rule{rule}})
		}
	}

	return guests
}

// addMyself add me to the list of guests
func addMyself(guests []Guest) []Guest {
	newRule := Rule{name: "Me", change: 0}
	names := []string{}

	for i := range guests {
		guests[i].rules = append(guests[i].rules, newRule)
		names = append(names, guests[i].name)
	}

	rulesForMyself := []Rule{}
	for _, n := range names {
		rulesForMyself = append(rulesForMyself, Rule{n, 0})
	}
	myself := Guest{"Me", rulesForMyself}

	guests = append(guests, myself)

	return guests
}
