// Description of the task: https://adventofcode.com/2015/day/4

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

func main() {
	content := helpers.Load("input")
	input := strings.TrimSpace(string(content))

	answer := findLowestSuffix(input, 5)
	fmt.Printf("Lowest positive number to mine AdventCoins with 5 zeroes is: %d\n", answer)

	answer = findLowestSuffix(input, 6)
	fmt.Printf("Lowest positive number to mine AdventCoins with 6 zeroes is: %d\n", answer)
}

func findLowestSuffix(key string, numberOfZeroes int) int {
	checkVal := ""
	for i := 0; i < numberOfZeroes; i++ {
		checkVal += "0"
	}

	for i := 1; ; i++ {
		combined := key + strconv.Itoa(i)
		sum := md5.Sum([]byte(combined))
		hash := hex.EncodeToString(sum[:])
		if hash[:numberOfZeroes] == checkVal {
			return i
		}
	}
}
