// Description of the task: https://adventofcode.com/2015/day/7
package main

import (
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

type command struct {
	operation   string // type of operation
	p1          string // parameter 1
	p2          string // parameter 2
	destination string // destination of result
}

func main() {
	lines := helpers.LoadLines("input")

	commands := []command{}
	for _, line := range lines {
		if line == "" {
			break
		}
		commands = append(commands, parseCommand(line))
	}

	storedValues := make(map[string]uint16)

	aVal := do(commands, storedValues, "a")
	fmt.Printf("Value of wire 'a' is: %d\n", aVal)

	bIndex := indexOfDest(commands, "b")
	commands[bIndex] = command{"MOVE", strconv.Itoa(int(aVal)), "", "b"}

	storedValues = make(map[string]uint16)
	newAVal := do(commands, storedValues, "a")
	fmt.Printf("New value of wire 'a' after forcing value of 'b' is: %d\n", newAVal)
}

// parseCommand fills and return command struct
func parseCommand(s string) command {
	parts := strings.Split(s, " -> ")
	tmp := strings.Split(parts[0], " ")
	c := command{}
	switch {
	case len(tmp) == 1:
		c = command{"MOVE", tmp[0], "", parts[1]}
	case len(tmp) == 2:
		c = command{"NOT", tmp[1], "", parts[1]}
	default:
		c = command{tmp[1], tmp[0], tmp[2], parts[1]}
	}

	return c
}

// do execute command of specific destination, recursive if necessary
func do(commands []command, storedValues map[string]uint16, dest string) uint16 {
	i := indexOfDest(commands, dest)
	v1 := loadValue(commands, storedValues, commands[i].p1)
	var result uint16
	switch {
	case commands[i].operation == "NOT":
		result = ^v1
	case commands[i].operation == "MOVE":
		result = v1
	case commands[i].operation == "AND":
		v2 := loadValue(commands, storedValues, commands[i].p2)
		result = v1 & v2
	case commands[i].operation == "OR":
		v2 := loadValue(commands, storedValues, commands[i].p2)
		result = v1 | v2
	case commands[i].operation == "RSHIFT":
		v2 := loadValue(commands, storedValues, commands[i].p2)
		result = v1 >> v2
	case commands[i].operation == "LSHIFT":
		v2 := loadValue(commands, storedValues, commands[i].p2)
		result = v1 << v2
	default:
		panic("wrong command")
	}
	storedValues[dest] = result

	return result
}

// loadValue loads value of parameter
func loadValue(commands []command, storedValues map[string]uint16, s string) uint16 {
	tmp, err := strconv.Atoi(s)
	v := uint16(tmp)
	if err != nil {
		var ok bool
		v, ok = storedValues[s]
		if !ok {
			v = do(commands, storedValues, s)
		}
	}
	return v
}

// indexOfDest returns position of specified destination in slice of commands
func indexOfDest(commands []command, dest string) int {
	for i := range commands {
		if commands[i].destination == dest {
			return i
		}
	}
	return -1
}
