package deck

import (
	"math/rand"
)

type Card struct {
	Rank int // 0, 1, 2, ..., 10, 11, 12 <=> deuce, trey, four, ..., queen, king, ace
	Suit int // 0 == club, 1 == diamond, 2 == heart, 3 == spade
}

type hand struct {
	cards []Card
}

func CreateDeck() []Card {
	var deck [52]Card
	for i := 0; i < 52; i++ {
		rank := i % 13
		suit := i % 4
		deck[i] = Card{Rank: rank, Suit: suit}
	}
	return deck[:]
}

func Shuffle(deck []Card, r *rand.Rand) {
	for i := 0; i < 51; i++ {
		j := r.Intn(52-i) + i
		tmp := deck[j]
		deck[j] = deck[i]
		deck[i] = tmp
	}
}

func Remove(removeCard Card, deck []Card) []Card {
	var newDeck []Card
	for i := 0; i < 52; i++ {
		if removeCard == deck[i] {
			newDeck = append(deck[:i], deck[i+1:]...)
			return newDeck
		}
	}
	return deck
}
