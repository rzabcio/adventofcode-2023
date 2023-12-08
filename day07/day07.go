package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day07_1(filename string) (result int) {
	deals := readDeals(filename)
	// sort deals by rank
	sort.Slice(deals, func(i, j int) bool {
		return compareDeals(*deals[i], *deals[j])
	})
	// calculate result
	for i, deal := range deals {
		deals[i].Rank = i + 1
		result += deal.Bet * deal.Rank
	}

	return result
}

func Day07_2(filename string) (result int) {
	deals := readDeals2(filename)
	// sort deals by rank
	sort.Slice(deals, func(i, j int) bool {
		return compareDeals(*deals[i], *deals[j])
	})
	// calculate result
	for i, deal := range deals {
		deals[i].Rank = i + 1
		result += deal.Bet * deal.Rank
	}
	return result
}

func readDeals(filename string) (deals []*Deal) {
	for line := range utils.InputCh(filename) {
		deal := NewDeal(line)
		deals = append(deals, deal)
	}
	return deals
}

func readDeals2(filename string) (deals []*Deal) {
	for line := range utils.InputCh(filename) {
		deal := NewDeal2(line)
		deals = append(deals, deal)
	}
	return deals
}

type Deal struct {
	// read values
	Bet   int
	Cards string // original: "T88QQ"
	// technical values
	Values    []int       // int values: [10 8 8 12 12]
	ValCounts map[int]int // counted: [10:1, 8:2, 12:2]
	ValSet    []int       // only counts, sorted: [2 2 1]
	SetSymbol string      // above but in string: "221" (for sorting)
	// rank
	Rank int
}

// --- part 1
var cardValues = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5,
	"6": 6, "7": 7, "8": 8, "9": 9,
	"T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
}

func NewDeal(s string) *Deal {
	deal := new(Deal)
	fields := strings.Fields(s)
	deal.ValCounts = make(map[int]int)
	deal.Cards = fields[0]
	for _, card := range deal.Cards { // change symbols to int values, count them
		if cardInt, ok := cardValues[string(card)]; ok {
			deal.Values = append(deal.Values, cardInt)
			deal.ValCounts[cardInt]++
		} else {
			panic("unknown card " + string(card))
		}
	}
	// get only counts, sort them, convert to string for ranking
	for _, cardInt := range deal.ValCounts {
		deal.ValSet = append(deal.ValSet, cardInt)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(deal.ValSet)))
	deal.SetSymbol = strings.Trim(strings.Replace(fmt.Sprint(deal.ValSet), " ", "", -1), "[]")

	// read bet
	if len(fields) > 1 {
		deal.Bet, _ = strconv.Atoi(fields[1])
	}
	return deal
}

// --- part 2
var cardValues2 = map[string]int{
	"J": 1, "2": 2, "3": 3, "4": 4, "5": 5,
	"6": 6, "7": 7, "8": 8, "9": 9,
	"T": 10, "Q": 12, "K": 13, "A": 14,
}

func NewDeal2(s string) *Deal {
	deal := new(Deal)
	fields := strings.Fields(s)
	deal.ValCounts = make(map[int]int)
	deal.Cards = fields[0]
	for _, card := range deal.Cards { // change symbols to int values, count them
		if cardInt, ok := cardValues2[string(card)]; ok {
			deal.Values = append(deal.Values, cardInt)
			deal.ValCounts[cardInt]++
		} else {
			panic("unknown card " + string(card))
		}
	}
	// get only counts, sort them, convert to string for ranking
	jokerCount := 0
	for cardVal, cardInt := range deal.ValCounts {
		if cardVal == 1 {
			jokerCount = cardInt
			continue
		}
		deal.ValSet = append(deal.ValSet, cardInt)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(deal.ValSet)))
	if len(deal.ValSet) == 0 {
		deal.ValSet = append(deal.ValSet, 0)
	}
	deal.ValSet[0] += jokerCount
	deal.SetSymbol = strings.Trim(strings.Replace(fmt.Sprint(deal.ValSet), " ", "", -1), "[]")

	// read bet
	if len(fields) > 1 {
		deal.Bet, _ = strconv.Atoi(fields[1])
	}
	return deal
}

// --- general
func compareDeals(deal1, deal2 Deal) bool {
	if deal1.SetSymbol == deal2.SetSymbol {
		// if set symbols (counter card, ordered, changed to string, i.e "221") are the same
		for i, card1 := range deal1.Values {
			card2 := deal2.Values[i]
			if card1 == card2 {
				continue
			}
			return card1 < card2
		}
	} else {
		// if not the same, compare set symbols
		return deal1.SetSymbol < deal2.SetSymbol
	}
	return false
}
