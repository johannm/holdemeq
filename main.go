package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/johannm/pokereq/deck"
	"github.com/johannm/pokereq/eval"
)

func parseHand(s string) []deck.Card {
	return []deck.Card{parseCard(s[0:2]), parseCard(s[2:4])}
}

func parseBoard(s string) []deck.Card {
	var b []deck.Card
	for len(s) > 0 {
		b = append(b, parseCard(s[0:2]))
		s = s[2:]
	}
	return b
}

func parseCard(s string) deck.Card {
	rank, suit := s[0], s[1]
	var c deck.Card
	switch rank {
	case 'A':
		c.Rank = 12
	case 'K':
		c.Rank = 11
	case 'Q':
		c.Rank = 10
	case 'J':
		c.Rank = 9
	case 'T':
		c.Rank = 8
	case '9':
		c.Rank = 7
	case '8':
		c.Rank = 6
	case '7':
		c.Rank = 5
	case '6':
		c.Rank = 4
	case '5':
		c.Rank = 3
	case '4':
		c.Rank = 2
	case '3':
		c.Rank = 1
	case '2':
		c.Rank = 0
	}
	switch suit {
	case 's':
		c.Suit = 3
	case 'h':
		c.Suit = 2
	case 'd':
		c.Suit = 1
	case 'c':
		c.Suit = 0
	}
	return c
}

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
	hand1 := parseHand(os.Args[1])
	hand2 := parseHand(os.Args[2])
	var board []deck.Card
	if len(os.Args) > 3 {
		board = parseBoard(os.Args[3])
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

		// Remove holecards from deck
		for _, c := range append(hand1, hand2...) {
			d.Remove(c)
		}

		// Deal out the rest of the cards
		var dealtBoard []deck.Card
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
