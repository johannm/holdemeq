package deck

import (
	"math/rand"
)

const numCards = 52

type Deck struct {
	cards []Card
}

type Card struct {
	Rank int // 0, 1, 2, ..., 10, 11, 12 <=> deuce, trey, four, ..., queen, king, ace
	Suit int // 0 == club, 1 == diamond, 2 == heart, 3 == spade
}

func CreateDeck() Deck {
	var deck Deck
	deck.cards = make([]Card, numCards)
	for i, _ := range deck.cards {
		rank := i % 13
		suit := i % 4
		deck.cards[i] = Card{Rank: rank, Suit: suit}
	}
	return deck
}

func (d *Deck) Shuffle(r *rand.Rand) {
	for i := 0; i < len(d.cards)-1; i++ {
		j := r.Intn(len(d.cards)-i) + i
		tmp := d.cards[j]
		d.cards[j] = d.cards[i]
		d.cards[i] = tmp
	}
}

func (d *Deck) DealOne() Card {
	var card Card
	card, d.cards = d.cards[0], d.cards[1:]
	return card
}

func (d *Deck) Remove(card Card) {
	for i, c := range d.cards {
		if card == c {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
		}
	}
}
