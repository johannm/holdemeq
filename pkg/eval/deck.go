package eval

import (
	"math/rand"
)

const numCards = 52

type Card struct {
	Rank int // 0, 1, 2, ..., 10, 11, 12 <=> deuce, trey, four, ..., queen, king, ace
	Suit int // 0 == club, 1 == diamond, 2 == heart, 3 == spade
}

func ParseStr(s string) []Card {
	var b []Card
	for len(s) > 0 {
		b = append(b, parseCard(s[0:2]))
		s = s[2:]
	}
	return b
}

func parseCard(s string) Card {
	rank, suit := s[0], s[1]
	var c Card
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

func (c *Card) ToStr() string {
	var rank, suit rune
	switch c.Rank {
	case 12:
		rank = 'A'
	case 11:
		rank = 'K'
	case 10:
		rank = 'Q'
	case 9:
		rank = 'J'
	case 8:
		rank = 'T'
	case 7:
		rank = '9'
	case 6:
		rank = '8'
	case 5:
		rank = '7'
	case 4:
		rank = '6'
	case 3:
		rank = '5'
	case 2:
		rank = '4'
	case 1:
		rank = '3'
	case 0:
		rank = '2'
	}
	switch c.Suit {
	case 3:
		suit = 's'
	case 2:
		suit = 'h'
	case 1:
		suit = 'd'
	case 0:
		suit = 'c'
	}
	return string(rank) + string(suit)
}

type deck struct {
	cards []Card
}

func createDeck() deck {
	var deck deck
	deck.cards = make([]Card, numCards)
	for i, _ := range deck.cards {
		rank := i % 13
		suit := i % 4
		deck.cards[i] = Card{Rank: rank, Suit: suit}
	}
	return deck
}

func (d *deck) shuffle(r *rand.Rand) {
	for i := 0; i < len(d.cards)-1; i++ {
		j := r.Intn(len(d.cards)-i) + i
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *deck) dealOne() Card {
	var card Card
	card, d.cards = d.cards[0], d.cards[1:]
	return card
}

func (d *deck) remove(card Card) {
	for i, c := range d.cards {
		if card == c {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
		}
	}
}

func (d deck) len() int {
	return len(d.cards)
}

func (d *deck) toStr() string {
	var s string
	for _, c := range d.cards {
		s += c.ToStr()
	}
	return s
}
