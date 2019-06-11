package eval

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestCreateDeck(t *testing.T) {
	deck1 := createDeck()
	deck2 := createDeck()
	if !reflect.DeepEqual(deck1, deck2) {
		t.Errorf("Expected decks to be equal after creation.")
	}
}

func TestCreateDeckSize(t *testing.T) {
	deck := createDeck()
	if res := len(deck.cards); res != 52 {
		t.Errorf("Expected value of %v, but was %v instead.", 52, res)
	}
}

func TestShuffle(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	deck1 := createDeck()
	deck2 := createDeck()
	deck1.shuffle(r1)
	if res := len(deck1.cards); res != 52 {
		t.Errorf("Expected value of %v, but was %v instead.", 52, res)
	}
	if reflect.DeepEqual(deck1, deck2) {
		t.Errorf("Expected deck to be not equal after shuffle.")
	}
}

func TestDealOne(t *testing.T) {
	deck := createDeck()
	card := deck.dealOne()
	twoOfSpades := Card{Rank: 0, Suit: 0}
	if res := len(deck.cards); res != 51 {
		t.Errorf("Expected value of %v, but was %v instead.", 51, res)
	}
	if card != twoOfSpades {
		t.Errorf("Expected first card of unshuffled deck to be 2s")
	}
}

func TestRemove(t *testing.T) {
	deck := createDeck()
	card := Card{Rank: 0, Suit: 0}
	deck.remove(card)
	if res := len(deck.cards); res != 51 {
		t.Errorf("Expected value of %v, but was %v instead.", 51, res)
	}
	for _, c := range deck.cards {
		if c == card {
			t.Errorf("Expected removed card to equal 2s")
		}
	}
}

func TestParseStr(t *testing.T) {
	s := "8h9h2c"
	cards := []Card{Card{Rank: 6, Suit: 2}, Card{Rank: 7, Suit: 2}, Card{Rank: 0, Suit: 0}}
	if res := ParseStr(s); !reflect.DeepEqual(res, cards) {
		t.Errorf("Expected value of %v, but was %v instead.", cards, res)
	}
}

func TestCardToStr(t *testing.T) {
	card := Card{Rank: 0, Suit: 0}
	if res := card.toStr(); res != "2c" {
		t.Errorf("Expected value of %v, but was %v instead.", "2s", res)
	}
}

func TestDeckToStr(t *testing.T) {
	deckStr := "2c3d4h5s6c7d8h9sTcJdQhKsAc2d3h4s5c6d7h8s9cTdJhQsKcAd2h3s4c5d6h7s8c9dThJsQcKdAh2s3c4d5h6s7c8d9hTsJcQdKhAs"
	d := createDeck()
	if res := d.toStr(); res != deckStr {
		t.Errorf("Expected value of %v, but was %v instead.", deckStr, res)
	}
}
