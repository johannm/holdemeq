package main

import (
	"testing"
	"reflect"
)

func TestMaxHand(t *testing.T) {
	seven := []card{
		card{rank: 12, suit: 0}, 
		card{rank: 11, suit: 0}, 
		card{rank: 10, suit: 0}, 
		card{rank: 10, suit: 1}, 
		card{rank: 10, suit: 2}, 
		card{rank: 9, suit: 0}, 
		card{rank: 8, suit: 0}}
	best5 := []card{
		card{rank: 12, suit: 0}, 
		card{rank: 11, suit: 0}, 
		card{rank: 10, suit: 0}, 
		card{rank: 9, suit: 0}, 
		card{rank: 8, suit: 0}}
	if res := findMaxhand(seven); !reflect.DeepEqual(res, best5) {
		t.Errorf("Expected value of %v, but was %v instead.", best5, res)
	} 
}