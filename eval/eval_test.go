package eval

import  (
	"testing"
	"github.com/johannm/pokereq/deck"
)

var (
	strflush   = []deck.Card{deck.Card{Rank: 12, Suit: 0}, deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}}
	quads      = []deck.Card{deck.Card{Rank: 12, Suit: 0}, deck.Card{Rank: 12, Suit: 1}, deck.Card{Rank: 12, Suit: 2}, deck.Card{Rank: 12, Suit: 3}, deck.Card{Rank: 11, Suit: 0}}
	boat       = []deck.Card{deck.Card{Rank: 12, Suit: 0}, deck.Card{Rank: 12, Suit: 1}, deck.Card{Rank: 12, Suit: 2}, deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 11, Suit: 1}}
	flush      = []deck.Card{deck.Card{Rank: 12, Suit: 0}, deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 7, Suit: 0}}
	straight   = []deck.Card{deck.Card{Rank: 12, Suit: 1}, deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}}
	trips      = []deck.Card{deck.Card{Rank: 12, Suit: 0}, deck.Card{Rank: 12, Suit: 1}, deck.Card{Rank: 12, Suit: 2}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 5, Suit: 0}}
	twopair    = []deck.Card{deck.Card{Rank: 12, Suit: 0}, deck.Card{Rank: 12, Suit: 1}, deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 10, Suit: 0}}
	onepair    = []deck.Card{deck.Card{Rank: 12, Suit: 0}, deck.Card{Rank: 12, Suit: 1}, deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}}
	hicard     = []deck.Card{deck.Card{Rank: 12, Suit: 1}, deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 7, Suit: 0}}
	wheel      = []deck.Card{deck.Card{Rank: 12, Suit: 1}, deck.Card{Rank: 0, Suit: 0}, deck.Card{Rank: 1, Suit: 0}, deck.Card{Rank: 2, Suit: 0}, deck.Card{Rank: 3, Suit: 0}}
	steelwheel = []deck.Card{deck.Card{Rank: 12, Suit: 0}, deck.Card{Rank: 0, Suit: 0}, deck.Card{Rank: 1, Suit: 0}, deck.Card{Rank: 2, Suit: 0}, deck.Card{Rank: 3, Suit: 0}}
)

func TestRankStraight(t *testing.T) {
	t.Logf("Testing rank of %v", straight)
	if res := RankHand(straight); res != 1600 {
		t.Errorf("Expected value of 1600, but was %d instead.", res)
	}
}

func TestRankHand(t *testing.T) {
	// todo: verify all ranks
	t.Logf("Testing rank of %v", strflush)
	if res := RankHand(strflush); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", quads)
	if res := RankHand(quads); res != 11 {
		t.Errorf("Expected value of 11, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", boat)
	if res := RankHand(boat); res != 167 {
		t.Errorf("Expected value of 167, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", flush)
	if res := RankHand(flush); res != 323 {
		t.Errorf("Expected value of 323, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", straight)
	if res := RankHand(straight); res != 1600 {
		t.Errorf("Expected value of 1600, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", trips)
	if res := RankHand(trips); res != 1634 {
		t.Errorf("Expected value of 1634, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", twopair)
	if res := RankHand(twopair); res != 2468 {
		t.Errorf("Expected value of 2468, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", onepair)
	if res := RankHand(onepair); res != 3326 {
		t.Errorf("Expected value of 3326, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", hicard)
	if res := RankHand(hicard); res != 6186 {
		t.Errorf("Expected value of 6186, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", wheel)
	if res := RankHand(wheel); res != 1609 {
		t.Errorf("Expected value of 1609, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", steelwheel)
	if res := RankHand(steelwheel); res != 10 {
		t.Errorf("Expected value of 10, but was %d instead.", res)
	}
}

func TestStrflushBeatsQuads(t *testing.T) {
	if res := Compare(strflush, quads); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(quads, strflush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestQuadsBeatsBoat(t *testing.T) {
	if res := Compare(quads, boat); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(boat, quads); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestBoatBeatsFlush(t *testing.T) {
	if res := Compare(boat, flush); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(flush, boat); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestFlushBeatsStraight(t *testing.T) {
	if res := Compare(flush, straight); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(straight, flush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestStraightBeatsTrips(t *testing.T) {
	if res := Compare(straight, trips); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(trips, straight); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestTripsBeatsTwopair(t *testing.T) {
	if res := Compare(trips, twopair); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(twopair, trips); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestTwopairBeatsOnepair(t *testing.T) {
	if res := Compare(twopair, onepair); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(onepair, twopair); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestOnepairBeatsHicard(t *testing.T) {
	if res := Compare(onepair, hicard); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(hicard, onepair); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestStraightBeatsWheel(t *testing.T) {
	straight := []deck.Card{deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}, deck.Card{Rank: 7, Suit: 0}}
	if res := Compare(straight, wheel); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(wheel, straight); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestWheelBeatsTrips(t *testing.T) {
	if res := Compare(wheel, trips); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(trips, wheel); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestStraightflushBeatsSteelwheel(t *testing.T) {
	strflush := []deck.Card{deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}, deck.Card{Rank: 7, Suit: 0}}
	if res := Compare(strflush, steelwheel); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(steelwheel, strflush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestSteelwheelBeatsQuads(t *testing.T) {
	if res := Compare(steelwheel, quads); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(quads, steelwheel); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestStraightFlushVsStraightFlush(t *testing.T) {
	strflush2 := []deck.Card{deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}, deck.Card{Rank: 7, Suit: 0}}

	if res := Compare(strflush, strflush2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(strflush2, strflush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(strflush2, strflush2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestQuadsVsQuads(t *testing.T) {
	quads2 := []deck.Card{deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 11, Suit: 2}, deck.Card{Rank: 11, Suit: 3}, deck.Card{Rank: 10, Suit: 0}}

	if res := Compare(quads, quads2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(quads2, quads); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(quads2, quads2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestBoatVsBoat(t *testing.T) {
	boat2 := []deck.Card{deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 11, Suit: 2}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 10, Suit: 1}}

	if res := Compare(boat, boat2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(boat2, boat); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(boat2, boat2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestFlushVsFlush(t *testing.T) {
	flush2 := []deck.Card{deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}, deck.Card{Rank: 6, Suit: 0}}

	if res := Compare(flush, flush2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(flush2, flush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(flush2, flush2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestStraightVsStraight(t *testing.T) {
	straight2 := []deck.Card{deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}, deck.Card{Rank: 7, Suit: 0}}

	if res := Compare(straight, straight2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(straight2, straight); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(straight2, straight2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestTripsVsTrips(t *testing.T) {
	trips2 := []deck.Card{deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 11, Suit: 2}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 5, Suit: 0}}

	if res := Compare(trips, trips2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(trips2, trips); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(trips2, trips2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestTwopairVsTwopair(t *testing.T) {
	twopair2 := []deck.Card{deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 10, Suit: 1}, deck.Card{Rank: 9, Suit: 0}}

	if res := Compare(twopair, twopair2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(twopair2, twopair); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(twopair2, twopair2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestOnepairVsOnepair(t *testing.T) {
	onepair2 := []deck.Card{deck.Card{Rank: 11, Suit: 0}, deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}}

	if res := Compare(onepair, onepair2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(onepair2, onepair); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(onepair2, onepair2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestHicardVsHicard(t *testing.T) {
	hicard2 := []deck.Card{deck.Card{Rank: 11, Suit: 1}, deck.Card{Rank: 10, Suit: 0}, deck.Card{Rank: 9, Suit: 0}, deck.Card{Rank: 8, Suit: 0}, deck.Card{Rank: 6, Suit: 0}}

	if res := Compare(hicard, hicard2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := Compare(hicard2, hicard); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := Compare(hicard2, hicard2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestWheelVsWheel(t *testing.T) {
	if res := Compare(wheel, wheel); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestSteelwheelVsSteelwheel(t *testing.T) {
	if res := Compare(steelwheel, steelwheel); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}