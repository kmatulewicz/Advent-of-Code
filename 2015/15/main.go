// Description of the task: https://adventofcode.com/2015/day/15

package main

import (
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

type Ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
	quantity                                        int
}

func main() {
	lines := helpers.LoadLines("input")
	ingredients := parseInput(lines)

	maxAll, max500 := calculateScores(ingredients, 100, 0)

	fmt.Printf("The best cookie reached score: %d\n", maxAll)
	fmt.Printf("For 500 calories cooke the best score is: %d\n", max500)
}

// calculateScores calculates max score for all cookies and max score for cookie with 500 calories
func calculateScores(ingredients []Ingredient, space int, i int) (int, int) {
	ch := make(chan []Ingredient)
	b := &block{} // counter for ongoing goroutines

	// start adding
	b.Add()
	go add(ingredients, space, i, ch, b)

	maxScore := 0
	max500 := 0
loop:
	for {
		select {
		case result := <-ch:
			// calculate score for the given combination of ingredients
			capacity, durability, flavor, texture, calories := 0, 0, 0, 0, 0
			for _, ingredient := range result {
				capacity += ingredient.quantity * ingredient.capacity
				durability += ingredient.quantity * ingredient.durability
				flavor += ingredient.quantity * ingredient.flavor
				texture += ingredient.quantity * ingredient.texture
				calories += ingredient.quantity * ingredient.calories
			}
			capacity = max(capacity, 0)
			durability = max(durability, 0)
			flavor = max(flavor, 0)
			texture = max(texture, 0)
			sc := capacity * durability * flavor * texture

			// if the score is better than saved remember it
			if sc > maxScore {
				maxScore = sc
			}
			if calories == 500 && sc > max500 {
				max500 = sc
			}
		default:
			// break loop if all goroutines are ended
			if b.Check() {
				break loop
			}
		}
	}

	return maxScore, max500
}

// add ads an ingredient pointed by ingredientIndex and if there is no space left it returns the combination to ch channel
func add(ingredients []Ingredient, space int, ingredientIndex int, ch chan []Ingredient, b *block) {

	// fill all the space
	for y := 0; y < space; y++ {
		//add the ingredient
		ingredients[ingredientIndex].quantity++

		// go to the next ingredient
		if ingredientIndex < len(ingredients)-1 {
			ci := make([]Ingredient, len(ingredients))
			copy(ci, ingredients)
			j := ingredientIndex + 1
			b.Add()
			go add(ci, space-y-1, j, ch, b)
		}

		// if there is no space left return the combination to ch
		if y == space-1 {
			ch <- ingredients
		}
	}
	b.Sub()
}

// parseInput parses provided lines of a description into slice of ingredients
func parseInput(lines []string) []Ingredient {

	ingredients := []Ingredient{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		ingredient := Ingredient{name: parts[0]}

		parts = strings.Split(parts[1], ", ")
		var err error
		ingredient.capacity, err = strconv.Atoi(strings.Split(parts[0], " ")[1])
		if err != nil {
			panic(err.Error())
		}
		ingredient.durability, err = strconv.Atoi(strings.Split(parts[1], " ")[1])
		if err != nil {
			panic(err.Error())
		}
		ingredient.flavor, err = strconv.Atoi(strings.Split(parts[2], " ")[1])
		if err != nil {
			panic(err.Error())
		}
		ingredient.texture, err = strconv.Atoi(strings.Split(parts[3], " ")[1])
		if err != nil {
			panic(err.Error())
		}
		ingredient.calories, err = strconv.Atoi(strings.Split(parts[4], " ")[1])
		if err != nil {
			panic(err.Error())
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients
}
