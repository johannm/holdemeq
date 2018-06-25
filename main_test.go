package main

import (
	"testing"
	"reflect"

	"github.com/johannm/pokereq/deck"
)

func TestMaxHand(t *testing.T) {
	seven := []deck.Card{
		deck.Card{Rank: 12, Suit: 0}, 
		deck.Card{Rank: 11, Suit: 0}, 
		deck.Card{Rank: 10, Suit: 0}, 
		deck.Card{Rank: 10, Suit: 1}, 
		deck.Card{Rank: 10, Suit: 2}, 
		deck.Card{Rank: 9, Suit: 0}, 
		deck.Card{Rank: 8, Suit: 0}}
	best5 := []deck.Card{
		deck.Card{Rank: 12, Suit: 0}, 
		deck.Card{Rank: 11, Suit: 0}, 
		deck.Card{Rank: 10, Suit: 0}, 
		deck.Card{Rank: 9, Suit: 0}, 
		deck.Card{Rank: 8, Suit: 0}}
	if res := findMaxhand(seven); !reflect.DeepEqual(res, best5) {
		t.Errorf("Expected value of %v, but was %v instead.", best5, res)
	} 
}