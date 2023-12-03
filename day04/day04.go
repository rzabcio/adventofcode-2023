package day04

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day04_1(filename string) (result int) {
	game := NewGame(filename)
	fmt.Printf("game: %+v\n", game)
	return 13
}

func Day04_2(filename string) (result int) {
	return result
}

type Game struct {
	Cards []*Card
}

type Card struct {
	No         int
	WinningNos []int
	CardNos    []int
}

func NewGame(filename string) *Game {
	game := &Game{}
	for line := range utils.InputCh(filename) {
		game.Cards = append(game.Cards, NewCard(line))
	}
	return game
}

func NewCard(line string) *Card {
	card := &Card{}
	regex := regexp.MustCompile(`Game (\d+): ([ \d]+) | ([ \d]+)`)
	matches := regex.FindStringSubmatch(line)
	card.No, _ = strconv.Atoi(matches[1])
	nos1String := matches[2]
	nos2String := matches[3]
	fmt.Printf("card.No: %d, winning: '%s', card: '%s'\n", card.No, nos1String, nos2String)
	return card
}
