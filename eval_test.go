package main

import "testing"

var (
	strflush   = []card{card{rank: 12, suit: 0}, card{rank: 11, suit: 0}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}}
	quads      = []card{card{rank: 12, suit: 0}, card{rank: 12, suit: 1}, card{rank: 12, suit: 2}, card{rank: 12, suit: 3}, card{rank: 11, suit: 0}}
	boat       = []card{card{rank: 12, suit: 0}, card{rank: 12, suit: 1}, card{rank: 12, suit: 2}, card{rank: 11, suit: 0}, card{rank: 11, suit: 1}}
	flush      = []card{card{rank: 12, suit: 0}, card{rank: 11, suit: 0}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 7, suit: 0}}
	straight   = []card{card{rank: 12, suit: 1}, card{rank: 11, suit: 0}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}}
	trips      = []card{card{rank: 12, suit: 0}, card{rank: 12, suit: 1}, card{rank: 12, suit: 2}, card{rank: 9, suit: 0}, card{rank: 5, suit: 0}}
	twopair    = []card{card{rank: 12, suit: 0}, card{rank: 12, suit: 1}, card{rank: 11, suit: 0}, card{rank: 11, suit: 1}, card{rank: 10, suit: 0}}
	onepair    = []card{card{rank: 12, suit: 0}, card{rank: 12, suit: 1}, card{rank: 11, suit: 0}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}}
	hicard     = []card{card{rank: 12, suit: 1}, card{rank: 11, suit: 0}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 7, suit: 0}}
	wheel      = []card{card{rank: 12, suit: 1}, card{rank: 0, suit: 0}, card{rank: 1, suit: 0}, card{rank: 2, suit: 0}, card{rank: 3, suit: 0}}
	steelwheel = []card{card{rank: 12, suit: 0}, card{rank: 0, suit: 0}, card{rank: 1, suit: 0}, card{rank: 2, suit: 0}, card{rank: 3, suit: 0}}
)

func TestRankStraight(t *testing.T) {
	t.Logf("Testing rank of %v", straight)
	if res := rankHand(straight); res != 1600 {
		t.Errorf("Expected value of 1600, but was %d instead.", res)
	}
}

func TestRankHand(t *testing.T) {
	// todo: verify all ranks
	t.Logf("Testing rank of %v", strflush)
	if res := rankHand(strflush); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", quads)
	if res := rankHand(quads); res != 11 {
		t.Errorf("Expected value of 11, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", boat)
	if res := rankHand(boat); res != 167 {
		t.Errorf("Expected value of 167, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", flush)
	if res := rankHand(flush); res != 323 {
		t.Errorf("Expected value of 323, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", straight)
	if res := rankHand(straight); res != 1600 {
		t.Errorf("Expected value of 1600, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", trips)
	if res := rankHand(trips); res != 1634 {
		t.Errorf("Expected value of 1634, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", twopair)
	if res := rankHand(twopair); res != 2468 {
		t.Errorf("Expected value of 2468, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", onepair)
	if res := rankHand(onepair); res != 3326 {
		t.Errorf("Expected value of 3326, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", hicard)
	if res := rankHand(hicard); res != 6186 {
		t.Errorf("Expected value of 6186, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", wheel)
	if res := rankHand(wheel); res != 1609 {
		t.Errorf("Expected value of 1609, but was %d instead.", res)
	}
	t.Logf("Testing rank of %v", steelwheel)
	if res := rankHand(steelwheel); res != 10 {
		t.Errorf("Expected value of 10, but was %d instead.", res)
	}
}

func TestStrflushBeatsQuads(t *testing.T) {
	if res := compare(strflush, quads); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(quads, strflush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestQuadsBeatsBoat(t *testing.T) {
	if res := compare(quads, boat); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(boat, quads); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestBoatBeatsFlush(t *testing.T) {
	if res := compare(boat, flush); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(flush, boat); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestFlushBeatsStraight(t *testing.T) {
	if res := compare(flush, straight); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(straight, flush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestStraightBeatsTrips(t *testing.T) {
	if res := compare(straight, trips); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(trips, straight); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestTripsBeatsTwopair(t *testing.T) {
	if res := compare(trips, twopair); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(twopair, trips); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestTwopairBeatsOnepair(t *testing.T) {
	if res := compare(twopair, onepair); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(onepair, twopair); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestOnepairBeatsHicard(t *testing.T) {
	if res := compare(onepair, hicard); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(hicard, onepair); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestStraightBeatsWheel(t *testing.T) {
	straight := []card{card{rank: 11, suit: 1}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}, card{rank: 7, suit: 0}}
	if res := compare(straight, wheel); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(wheel, straight); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestWheelBeatsTrips(t *testing.T) {
	if res := compare(wheel, trips); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(trips, wheel); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestStraightflushBeatsSteelwheel(t *testing.T) {
	strflush := []card{card{rank: 11, suit: 0}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}, card{rank: 7, suit: 0}}
	if res := compare(strflush, steelwheel); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(steelwheel, strflush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestSteelwheelBeatsQuads(t *testing.T) {
	if res := compare(steelwheel, quads); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(quads, steelwheel); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
}

func TestStraightFlushVsStraightFlush(t *testing.T) {
	strflush2 := []card{card{rank: 11, suit: 0}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}, card{rank: 7, suit: 0}}

	if res := compare(strflush, strflush2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(strflush2, strflush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(strflush2, strflush2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestQuadsVsQuads(t *testing.T) {
	quads2 := []card{card{rank: 11, suit: 0}, card{rank: 11, suit: 1}, card{rank: 11, suit: 2}, card{rank: 11, suit: 3}, card{rank: 10, suit: 0}}

	if res := compare(quads, quads2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(quads2, quads); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(quads2, quads2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestBoatVsBoat(t *testing.T) {
	boat2 := []card{card{rank: 11, suit: 0}, card{rank: 11, suit: 1}, card{rank: 11, suit: 2}, card{rank: 10, suit: 0}, card{rank: 10, suit: 1}}

	if res := compare(boat, boat2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(boat2, boat); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(boat2, boat2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestFlushVsFlush(t *testing.T) {
	flush2 := []card{card{rank: 11, suit: 0}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}, card{rank: 6, suit: 0}}

	if res := compare(flush, flush2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(flush2, flush); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(flush2, flush2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestStraightVsStraight(t *testing.T) {
	straight2 := []card{card{rank: 11, suit: 1}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}, card{rank: 7, suit: 0}}

	if res := compare(straight, straight2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(straight2, straight); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(straight2, straight2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestTripsVsTrips(t *testing.T) {
	trips2 := []card{card{rank: 11, suit: 0}, card{rank: 11, suit: 1}, card{rank: 11, suit: 2}, card{rank: 9, suit: 0}, card{rank: 5, suit: 0}}

	if res := compare(trips, trips2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(trips2, trips); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(trips2, trips2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestTwopairVsTwopair(t *testing.T) {
	twopair2 := []card{card{rank: 11, suit: 0}, card{rank: 11, suit: 1}, card{rank: 10, suit: 0}, card{rank: 10, suit: 1}, card{rank: 9, suit: 0}}

	if res := compare(twopair, twopair2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(twopair2, twopair); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(twopair2, twopair2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestOnepairVsOnepair(t *testing.T) {
	onepair2 := []card{card{rank: 11, suit: 0}, card{rank: 11, suit: 1}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}}

	if res := compare(onepair, onepair2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(onepair2, onepair); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(onepair2, onepair2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestHicardVsHicard(t *testing.T) {
	hicard2 := []card{card{rank: 11, suit: 1}, card{rank: 10, suit: 0}, card{rank: 9, suit: 0}, card{rank: 8, suit: 0}, card{rank: 6, suit: 0}}

	if res := compare(hicard, hicard2); res != 1 {
		t.Errorf("Expected value of 1, but was %d instead.", res)
	}
	if res := compare(hicard2, hicard); res != -1 {
		t.Errorf("Expected value of -1, but was %d instead.", res)
	}
	if res := compare(hicard2, hicard2); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestWheelVsWheel(t *testing.T) {
	if res := compare(wheel, wheel); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}

func TestSteelwheelVsSteelwheel(t *testing.T) {
	if res := compare(steelwheel, steelwheel); res != 0 {
		t.Errorf("Expected value of 0, but was %d instead.", res)
	}
}