package day04

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day04_1(filename string) (result int) {
	game := NewGame(filename)
	return game.FindLuckyValue()
}

func Day04_2(filename string) (result int) {
	return result
}

type Game struct {
	Cards []*Card
}

type Card struct {
	No      int
	Winning []int
	Random  []int
	Lucky   []int
}

func NewGame(filename string) *Game {
	game := &Game{}
	for line := range utils.InputCh(filename) {
		game.Cards = append(game.Cards, NewCard(line))
	}
	return game
}

func (g *Game) FindLuckyValue() int {
	result := 0
	for _, card := range g.Cards {
		result += card.FindLuckyValue()
	}
	return result
}

func NewCard(line string) *Card {
	card := &Card{}
	cardReg := regexp.MustCompile(`Card.*(\d+): ([\d ]+) \| ([\d ]+)`)
	numsReg := regexp.MustCompile(`(\d+)`)

	cardMatches := cardReg.FindStringSubmatch(line)
	card.No, _ = strconv.Atoi(cardMatches[1])
	numsMatches := numsReg.FindAllStringSubmatch(cardMatches[2], -1)
	for _, match := range numsMatches {
		no, _ := strconv.Atoi(match[1])
		card.Winning = append(card.Winning, no)
	}
	numsMatches = numsReg.FindAllStringSubmatch(cardMatches[3], -1)
	for _, match := range numsMatches {
		no, _ := strconv.Atoi(match[1])
		card.Random = append(card.Random, no)
	}
	card.FindLuckyNums()
	return card
}

func (g *Card) FindLuckyNums() {
	for _, winning := range g.Winning {
		for _, random := range g.Random {
			if winning == random {
				g.Lucky = append(g.Lucky, winning)
			}
		}
	}
}

func (g *Card) FindLuckyValue() int {
	if len(g.Lucky) == 0 {
		return 0
	}
	return int(math.Pow(float64(2), float64(len(g.Lucky)-1)))
}

func (g *Game) String() string {
	result := "Game{\n"
	for _, card := range g.Cards {
		result += fmt.Sprintf("  %v\n", card)
	}
	result += "}"
	return result
}

func (c *Card) String() string {
	return fmt.Sprintf("Card{No: %d, Winning: %v, Random: %v, Lucky: %v => %d}", c.No, c.Winning, c.Random, c.Lucky, c.FindLuckyValue())
}
