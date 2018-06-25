package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func parseHand(s string) []card {
	return []card{parseCard(s[0:2]), parseCard(s[2:4])}
}

func parseBoard(s string) []card {
	var b []card
	for len(s) > 0 {
		b = append(b, parseCard(s[0:2]))
		s = s[2:]
	}
	return b
}

func parseCard(s string) card {
	rank, suit := s[0], s[1]
	var c card
	switch rank {
	case 'A':
		c.rank = 12
	case 'K':
		c.rank = 11
	case 'Q':
		c.rank = 10
	case 'J':
		c.rank = 9
	case 'T':
		c.rank = 8
	case '9':
		c.rank = 7
	case '8':
		c.rank = 6
	case '7':
		c.rank = 5
	case '6':
		c.rank = 4
	case '5':
		c.rank = 3
	case '4':
		c.rank = 2
	case '3':
		c.rank = 1
	case '2':
		c.rank = 0
	}
	switch suit {
	case 's':
		c.suit = 3
	case 'h':
		c.suit = 2
	case 'd':
		c.suit = 1
	case 'c':
		c.suit = 0
	}
	return c
}

func findMaxhand(cards []card) []card {
	hand := make([]card, 5)
	maxHand := make([]card, 5)
	bestRank := 100000
	for _, combo := range perm7 {
		for i, cardIndex := range combo {
			hand[i] = cards[cardIndex]
		}
		r := rankHand(hand)
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
	var board []card
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
		deck := createDeck()
		shuffle(deck, r1)
		
		// Remove holecards from deck
		for _, c := range append(hand1, hand2...) {
			deck = remove(c, deck)
		}

		// Deal out the rest of the cards
		dealtBoard := append(board, deck[0:5-len(board)]...)

		maxhand1 := findMaxhand(append(hand1, dealtBoard...))
		maxhand2 := findMaxhand(append(hand2, dealtBoard...))

		won := compare(maxhand1, maxhand2)
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
