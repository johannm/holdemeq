package deck

import (
	"testing"
	"reflect"
	"time"
	"math/rand"
)

func TestCreatdDeckSize(t *testing.T) {
	deck := CreateDeck()
	if res := len(deck.cards); res != 52 {
		t.Errorf("Expected value of %v, but was %v instead.", 52, res)
	} 
}

func TestShuffle(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	deck1 := CreateDeck()
	deck2 := CreateDeck()
	deck2.Shuffle(r1)
	if res := len(deck2.cards); res != 52 {
		t.Errorf("Expected value of %v, but was %v instead.", 52, res)
	}
	if reflect.DeepEqual(deck1, deck2) {
		t.Errorf("Expected deck to be not equal after shuffle.")
	} 
}

func TestDealOne(t *testing.T) {
	deck := CreateDeck()
	card := deck.DealOne()
	twoOfSpades := Card{Rank: 0, Suit:0}
	if res := len(deck.cards); res != 51 {
		t.Errorf("Expected value of %v, but was %v instead.", 51, res)
	}
	if card != twoOfSpades{
		t.Errorf("Expected first card of unshuffled deck to be 2s")
	}
}

func TestRemove(t *testing.T) {
	deck := CreateDeck()
	card := Card{Rank: 0, Suit: 0}
	deck.Remove(card)
	if res := len(deck.cards); res != 51 {
		t.Errorf("Expected value of %v, but was %v instead.", 51, res)
	}
	for _, c := range deck.cards {
		if c == card {
			t.Errorf("Expected removed card to equal 2s")
		}
	}
}