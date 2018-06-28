package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/johannm/pokereq/deck"
	"github.com/johannm/pokereq/eval"
)

func findMaxhand(cards []deck.Card) []deck.Card {
	hand := make([]deck.Card, 5)
	maxHand := make([]deck.Card, 5)
	bestRank := 100000
	for _, combo := range eval.Perm7 {
		for i, cardIndex := range combo {
			hand[i] = cards[cardIndex]
		}
		r := eval.RankHand(hand)
		if r < bestRank {
			bestRank = r
			copy(maxHand, hand)
		}
	}
	return maxHand
}

func main() {
	fmt.Printf("Hold-em Hand Equity Calculator\n")
	for i, arg := range os.Args[1:] {
		fmt.Printf("arg %d: %v\n", i, arg)
	}
	hand1 := deck.ParseStr(os.Args[1])
	hand2 := deck.ParseStr(os.Args[2])
	var board []deck.Card
	if len(os.Args) > 3 {
		board = deck.ParseStr(os.Args[3])
		fmt.Printf("board: %v\n", board)
	}
	fmt.Printf("hand1: %v, hand2: %v\n", hand1, hand2)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	n := 1000000
	win, lose, draw := 0, 0, 0

	for i := 0; i < n; i++ {
		d := deck.CreateDeck()
		d.Shuffle(r1)

		// Remove holecards and board from deck
		for _, c := range append(append(hand1, hand2...), board...) {
			d.Remove(c)
		}

		// Deal out the rest of the cards
		dealtBoard := make([]deck.Card, len(board))
		copy(dealtBoard, board)
		for i := 0; i < 5-len(board); i++ {
			dealtBoard = append(dealtBoard, d.DealOne())
		}

		maxhand1 := findMaxhand(append(hand1, dealtBoard...))
		maxhand2 := findMaxhand(append(hand2, dealtBoard...))

		won := eval.Compare(maxhand1, maxhand2)
		if won > 0 {
			win++
		} else if won < 0 {
			lose++
		} else {
			draw++
		}
	}

	fmt.Printf("n: %v, win: %v, lose: %v, draw: %v\n", n, float64(win)/float64(n), float64(lose)/float64(n), float64(draw)/float64(n))
	fmt.Printf("equity: %v\n", (float64(win)+0.5*float64(draw))/float64(n))
}
