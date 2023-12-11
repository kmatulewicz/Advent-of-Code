// Description of the task: https://adventofcode.com/2015/day/10

package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	input := []rune("1321131112")

	out := lookAndSay(input, 40)
	fmt.Printf("The result length of part 1 is: %d\n", len(out))

	// very long execution
	out = lookAndSay([]rune(out), 10)
	fmt.Printf("The result length of part 2 is: %d\n", len(out))
}

func lookAndSay(input []rune, iterations int) string {
	output := ""
	for i := 0; i < iterations; i++ {
		count := 0
		last := input[0]
		output = ""
		for _, c := range input {
			if c == last {
				count++
			} else {
				output += strconv.Itoa(count) + string(last)
				count = 1
				last = c
			}
		}
		// last rune
		output += strconv.Itoa(count) + string(last)

		log.Printf("Iteration %d length: %d\n", i+1, len(output))

		// output is the next input
		input = []rune(output)
	}

	return output
}
