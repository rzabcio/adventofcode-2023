package main

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	redMax   = 12
	greenMax = 13
	blueMax  = 14
)

func Day02_1(filename string) (result int) {
	games := ReadGames(filename)
	for _, game := range games {
		if game.RedCont <= redMax && game.GreenCount <= greenMax && game.BlueCount <= blueMax {
			result += game.No
		}
	}
	return result
}

func Day02_2(filename string) (result int) {
	return result
}

type Game struct {
	No         int
	RedCont    int
	GreenCount int
	BlueCount  int
	Rounds     []GameRound
}

type GameRound struct {
	RedCount   int
	GreenCount int
	BlueCount  int
}

func ReadGames(filename string) (games []Game) {
	for line := range inputCh(filename) {
		games = append(games, ReadGame(line))
	}
	return games
}

func ReadGame(line string) (game Game) {
	reg := regexp.MustCompile("Game (\\d+): (.*)")
	matches := reg.FindStringSubmatch(line)
	var err error
	game.No, err = strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	game.Rounds = make([]GameRound, 0)

	redreg := regexp.MustCompile(`(\d+) red`)
	greenreg := regexp.MustCompile(`(\d+) green`)
	bluereg := regexp.MustCompile(`(\d+) blue`)

	for _, roundString := range strings.Split(matches[2], ";") {
		round := GameRound{}

		redmatches := redreg.FindStringSubmatch(roundString)
		if len(redmatches) > 0 {
			round.RedCount, _ = strconv.Atoi(redmatches[1])
			if round.RedCount > game.RedCont {
				game.RedCont = round.RedCount
			}
		}

		greenmatches := greenreg.FindStringSubmatch(roundString)
		if len(greenmatches) > 0 {
			round.GreenCount, _ = strconv.Atoi(greenmatches[1])
			if round.GreenCount > game.GreenCount {
				game.GreenCount = round.GreenCount
			}
		}

		bluematches := bluereg.FindStringSubmatch(roundString)
		if len(bluematches) > 0 {
			round.BlueCount, _ = strconv.Atoi(bluematches[1])
			if round.BlueCount > game.BlueCount {
				game.BlueCount = round.BlueCount
			}
		}
		game.Rounds = append(game.Rounds, round)
	}
	return game
}
