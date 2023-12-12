// Description of the task: https://adventofcode.com/2015/day/12

package main

import (
	"encoding/json"
	"fmt"
	"helpers"
)

func main() {
	input := helpers.Load("input")
	sum := sumDigits(input, false)
	fmt.Printf("Sum of all digits in JSON is: %.0f\n", sum)

	sum = sumDigits(input, true)
	fmt.Printf("Sum of all digits in JSON after removing red is: %.0f\n", sum)
}

// sumDigits returns sum of all digits from JSON values
func sumDigits(input string, removeRed bool) float64 {
	var data any
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err.Error())
	}

	switch v := data.(type) {
	case []any:
		return run(v, removeRed)
	case map[string]any:
		return run(mapToSlice(v, removeRed), removeRed)
	default:
		panic("wrong type")
	}
}

// run runs recursively through slice and returns sum of all digits in it
func run(m []any, removeRed bool) float64 {
	sum := float64(0)
	for _, v := range m {
		switch vv := v.(type) {
		case float64:
			sum += vv
		case []any:
			sum += run(vv, removeRed)
		case map[string]any:
			sum += run(mapToSlice(vv, removeRed), removeRed)
		}
	}

	return sum
}

// mapToSlice converts map to slice
func mapToSlice(m map[string]any, removeRed bool) []any {
	s := make([]any, 0)
	for _, v := range m {
		if removeRed && isRed(v) {
			return []any{}
		}
		s = append(s, v)
	}

	return s
}

// isRed checks if value is red
func isRed(v any) bool {
	vv, ok := v.(string)
	if ok {
		if vv == "red" {
			return true
		}
	}
	return false
}
