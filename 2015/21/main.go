// Description of the task: https://adventofcode.com/2015/day/21

package main

import (
	"fmt"
	"helpers"
	"math"
	"strconv"
	"strings"
)

type Player struct {
	name   string
	damage int
	armor  int
	health int
}

type Shop struct {
	weapons []Item
	armor   []Item
	rings   []Item
}

type Item struct {
	name   string
	cost   int
	damage int
	armor  int
}

func main() {
	fmt.Println("Part 1: ", calcMinInvPrice())
	fmt.Println("Part 2: ", calcMaxInvPrice())
}

// calcMinInvPrice calculates the price of the minimum equipment needed to beat the boss
func calcMinInvPrice() int {

	boss := getBoss("input.boss")

	sum := math.MaxInt

	combinations := getCombinations()

	for _, c := range combinations {
		b := boss
		if fight(c.p, b).name == "me" {
			sum = min(sum, c.sum)
		}
	}

	return sum
}

// calcMaxInvPrice calculates the maximum price of the equipment that still is not enough to win
func calcMaxInvPrice() int {

	boss := getBoss("input.boss")

	sum := 0

	combinations := getCombinations()

	for _, c := range combinations {
		b := boss
		if fight(c.p, b).name != "me" {
			sum = max(sum, c.sum)
		}
	}

	return sum
}

// C is a equipment set and its cost
type C struct {
	p   Player
	sum int
}

// getCombinations returns all combinations of equipment
func getCombinations() []C {
	shop := getShop("input.shop")

	combinations := []C{}

	// weapon
	for _, a := range shop.weapons {
		player := Player{name: "me", health: 100, damage: a.damage}
		sum := a.cost
		combinations = append(combinations, C{player, sum})
	}

	//armor
	newCombinations := []C{}
	for _, comb := range combinations {
		for _, a := range shop.armor {
			player := comb.p
			sum := comb.sum
			player.armor += a.armor
			sum += a.cost
			newCombinations = append(newCombinations, C{player, sum})
		}
	}
	combinations = append(combinations, newCombinations...)

	//rings
	newCombinations = []C{}
	for _, comb := range combinations {
		for _, a := range shop.rings {
			player := comb.p
			sum := comb.sum
			player.damage += a.damage
			player.armor += a.armor
			sum += a.cost
			//first ring
			newCombinations = append(newCombinations, C{player, sum})
			for _, b := range shop.rings {
				if b.name == a.name {
					continue
				}
				newPlayer := player
				newSum := sum
				newPlayer.damage += b.damage
				newPlayer.armor += b.armor
				newSum += b.cost
				//second ring
				newCombinations = append(newCombinations, C{newPlayer, newSum})
			}
		}
	}
	combinations = append(combinations, newCombinations...)

	return combinations
}

// fight simulates a fight between p1 and p2 and returns the player who won, p1 always starts
func fight(p1, p2 Player) Player {

	for p1.health > 0 && p2.health > 0 {
		p2.health -= max(1, p1.damage-p2.armor)
		//fmt.Printf("%s deals %d-%d = %d damage; %s goes down to %d hit points\n", p1.name, p1.damage, p2.armor, p1.damage-p2.armor, p2.name, p2.health)
		if p2.health < 1 {
			return p1
		}
		p1.health -= max(1, p2.damage-p1.armor)
		//fmt.Printf("%s deals %d-%d = %d damage; %s goes down to %d hit points\n", p2.name, p2.damage, p1.armor, p2.damage-p1.armor, p1.name, p1.health)
		if p1.health < 1 {
			return p2
		}

	}

	if p1.health > 0 {
		return p1
	}
	return p2

}

// getShop returns a shop with all equipments types in separate slices
func getShop(name string) Shop {
	shop := Shop{}
	lines := helpers.LoadLines(name)
	t := '0'
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		line := removeExcessSpaces(line)
		parts := strings.Split(line, " ")

		// line with a name of a equipment group
		if parts[0][len(parts[0])-1] == ':' {
			parts[0] = parts[0][:len(parts[0])-1]
			switch parts[0] {
			case "Weapons":
				t = 'w'
			case "Armor":
				t = 'a'
			case "Rings":
				t = 'r'
			default:
				panic("wrong type of inventory")
			}
			continue
		}

		// parse stats
		name := ""
		cost, damage, armor := 0, 0, 0
		var err error
		switch t {
		case 'w', 'a':
			name = parts[0]
			cost, err = strconv.Atoi(parts[1])
			if err != nil {
				panic(err.Error())
			}
			damage, err = strconv.Atoi(parts[2])
			if err != nil {
				panic(err.Error())
			}
			armor, err = strconv.Atoi(parts[3])
			if err != nil {
				panic(err.Error())
			}
		case 'r':
			name = parts[0] + " " + parts[1]
			cost, err = strconv.Atoi(parts[2])
			if err != nil {
				panic(err.Error())
			}
			damage, err = strconv.Atoi(parts[3])
			if err != nil {
				panic(err.Error())
			}
			armor, err = strconv.Atoi(parts[4])
			if err != nil {
				panic(err.Error())
			}
		default:
			panic("wrong t")
		}

		// append to correct group
		switch t {
		case 'w':
			shop.weapons = append(shop.weapons, Item{name, cost, damage, armor})

		case 'a':
			shop.armor = append(shop.armor, Item{name, cost, damage, armor})

		case 'r':
			shop.rings = append(shop.rings, Item{name, cost, damage, armor})
		}
	}

	return shop
}

// removeExcessSpaces returns a string containing only one space in a row
func removeExcessSpaces(in string) string {
	last := '0'
	new := ""
	for _, c := range in {
		if last == ' ' && c == ' ' {
			continue
		}
		new += string(c)
		last = c
	}
	return new
}

// getBoss parse the boss stats and returns them in a boss struct
func getBoss(name string) Player {
	boss := Player{name: "Boss"}
	var err error
	lines := helpers.LoadLines(name)
	//health
	parts := strings.Split(lines[0], ":")
	boss.health, err = strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic(err.Error())
	}
	//damage
	parts = strings.Split(lines[1], ":")
	boss.damage, err = strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic(err.Error())
	}
	//armor
	parts = strings.Split(lines[2], ":")
	boss.armor, err = strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic(err.Error())
	}

	return boss
}
