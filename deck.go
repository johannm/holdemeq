package main

import (
	"math/rand"
)

type card struct {
	rank int // 0, 1, 2, ..., 10, 11, 12 <=> deuce, trey, four, ..., queen, king, ace
	suit int // 0 == club, 1 == diamond, 2 == heart, 3 == spade
}

type hand struct {
	cards []card
}

func createDeck() []card {
	var deck [52]card
	for i := 0; i < 52; i++ {
		rank := i % 13
		suit := i % 4
		deck[i] = card{rank: rank, suit: suit}
	}
	return deck[:]
}

func shuffle(deck []card, r *rand.Rand) {
	for i := 0; i < 51; i++ {
		j := r.Intn(52-i) + i
		tmp := deck[j]
		deck[j] = deck[i]
		deck[i] = tmp
	}
}

func remove(removeCard card, deck []card) []card {
	var newDeck []card
	for i := 0; i < 52; i++ {
		if removeCard == deck[i] {
			newDeck = append(deck[:i], deck[i+1:]...)
			return newDeck
		}
	}
	return deck
}
