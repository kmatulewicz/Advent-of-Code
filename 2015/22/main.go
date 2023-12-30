// Description of the task: https://adventofcode.com/2015/day/22

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"helpers"
	"math"
	"strconv"
	"strings"
)

func main() {
	boss := getBoss("input.boss")
	player := Player{
		Health: 50,
		Mana:   500,
		Spells: []string{"Magic Missile", "Drain", "Shield", "Poison", "Recharge"},
		Buffs:  map[string]int{},
	}

	fmt.Println("Part 1:", battle(player, boss, false))
	fmt.Println("Part 2:", battle(player, boss, true))
}

type Player struct {
	Health int
	Mana   int
	Buffs  map[string]int
	Spells []string
}

type Boss struct {
	Health int
	Damage int
	Buffs  map[string]int
}

type Pair struct {
	P        Player
	B        Boss
	ManaUsed uint
}

// battle performs the battle between the player and the boss and returns the minimal amount of used mana that allows the player to win.
func battle(p Player, b Boss, part2 bool) uint {
	pairs := []Pair{
		{p, b, 0},
	}
	win := uint(math.MaxUint)
	for {
		newPairs := []Pair{}
		for _, pair := range pairs {
			for _, spell := range pair.P.Spells {
				cPair := copyAny(pair)

				// Player move

				// Part 2
				if part2 {
					cPair.P.Health--
					if cPair.P.Health < 1 {
						continue
					}
				}

				r, mu := pMove(&cPair.P, &cPair.B, spell)
				// player lost
				if r == -1 {
					continue
				}
				cPair.ManaUsed += mu

				// player win
				if r == 1 {
					win = min(cPair.ManaUsed, win)
					continue
				}

				// Boss move
				r = bMove(&cPair.P, &cPair.B)
				// player lost
				if r == -1 {
					continue
				}
				// player win
				if r == 1 {
					win = min(cPair.ManaUsed, win)
					continue
				}

				if win == math.MaxUint || cPair.ManaUsed < win {
					newPairs = append(newPairs, cPair)
				}
			}
		}

		// no more games
		if len(newPairs) == 0 {
			break
		}

		pairs = newPairs
	}

	return win
}

// pMove performs the player move and returns the result (-1 player lost, 0 inconclusive move, 1 player won) and amount of mana used
func pMove(p *Player, b *Boss, s string) (result int, mu uint) {
	_, ok := b.Buffs["Poison"]
	if ok {
		b.Health -= 3
		debuff(b.Buffs, "Poison")
	}
	if b.Health < 1 {
		// player won - boss poisoned to death
		return 1, 0
	}

	_, ok = p.Buffs["Recharge"]
	if ok {
		p.Mana += 101
		debuff(p.Buffs, "Recharge")
	}

	_, ok = p.Buffs["Shield"]
	if ok {
		debuff(p.Buffs, "Shield")
	}

	mu, err := doSpell(p, b, s)
	if err != nil {
		// player lost - not enough mana
		return -1, 0
	}

	if b.Health < 1 {
		// player won - boss dead
		return 1, mu
	}

	return 0, mu
}

// bMove performs the boss move and returns the result (-1 player lost, 0 inconclusive move, 1 player won)
func bMove(p *Player, b *Boss) (result int) {
	_, ok := b.Buffs["Poison"]
	if ok {
		b.Health -= 3
		debuff(b.Buffs, "Poison")
	}
	if b.Health < 1 {
		// player won - boss poisoned to death
		return 1
	}

	_, ok = p.Buffs["Recharge"]
	if ok {
		p.Mana += 101
		debuff(p.Buffs, "Recharge")
	}

	_, shield := p.Buffs["Shield"]

	if shield {
		p.Health -= max(1, b.Damage-7)
		debuff(p.Buffs, "Shield")
	} else {
		p.Health -= max(1, b.Damage)
	}

	if p.Health < 1 {
		// player lost - dead
		return -1
	}

	return 0
}

// copy copies any variable and returns its copy. Used for copying complex structs
func copyAny[T any](s T) (o T) {
	c, err := json.Marshal(s)
	if err != nil {
		panic("cannot marshal")
	}
	err = json.Unmarshal(c, &o)
	if err != nil {
		panic("cannot unmarshal")
	}

	return
}

// debuff decreases the counter on the selected buff and removes it on 0
func debuff(b map[string]int, n string) {
	b[n] -= 1
	if b[n] == 0 {
		delete(b, n)
	}
}

// getBoss returns a boss struct created from a pointed file
func getBoss(name string) Boss {
	lines := helpers.LoadLines(name)

	// health
	parts := strings.Split(lines[0], ": ")
	h, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err.Error())
	}

	// damage
	parts = strings.Split(lines[1], ": ")
	d, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err.Error())
	}

	return Boss{h, d, map[string]int{}}
}

// doSpell performs the selected spell
func doSpell(p *Player, b *Boss, spell string) (uint, error) {
	switch spell {
	case "Magic Missile":
		if p.Mana < 53 {
			return 0, errors.New("not enough mana")
		}
		p.Mana -= 53
		b.Health -= 4
		return 53, nil
	case "Drain":
		if p.Mana < 73 {
			return 0, errors.New("not enough mana")
		}
		p.Mana -= 73
		p.Health += 2
		b.Health -= 2
		return 73, nil
	case "Shield":
		if p.Mana < 113 {
			return 0, errors.New("not enough mana")
		}
		if _, ok := p.Buffs["Shield"]; ok {
			return 0, errors.New("the buff already exists")
		}
		p.Mana -= 113
		p.Buffs["Shield"] = 6
		return 113, nil
	case "Poison":
		if p.Mana < 173 {
			return 0, errors.New("not enough mana")
		}
		if _, ok := p.Buffs["Poison"]; ok {
			return 0, errors.New("the buff already exists")
		}
		p.Mana -= 173
		b.Buffs["Poison"] = 6
		return 173, nil
	case "Recharge":
		if p.Mana < 229 {
			return 0, errors.New("not enough mana")
		}
		if _, ok := p.Buffs["Recharge"]; ok {
			return 0, errors.New("the buff already exists")
		}
		p.Mana -= 229
		p.Buffs["Recharge"] = 5
		return 229, nil
	default:
		panic("wrong spell")
	}
}
