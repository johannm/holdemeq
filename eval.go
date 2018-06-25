package main

const STRAIGHT_FLUSH = 1
const FOUR_OF_A_KIND = 2
const FULL_HOUSE = 3
const FLUSH = 4
const STRAIGHT = 5
const THREE_OF_A_KIND = 6
const TWO_PAIR = 7
const ONE_PAIR = 8
const HIGH_CARD = 9

var value_str = [...]string{
	"",
	"Straight Flush",
	"Four of a Kind",
	"Full House",
	"Flush",
	"Straight",
	"Three of a Kind",
	"Two Pair",
	"One Pair",
	"High Card"}

const Deuce = 0
const Trey = 1
const Four = 2
const Five = 3
const Six = 4
const Seven = 5
const Eight = 6
const Nine = 7
const Ten = 8
const Jack = 9
const Queen = 10
const King = 11
const Ace = 12

func hand_rank(val uint16) int {
	if val > 6185 {
		return HIGH_CARD // 1277 high card
	}
	if val > 3325 {
		return ONE_PAIR // 2860 one pair
	}
	if val > 2467 {
		return TWO_PAIR // 858 two pair
	}
	if val > 1609 {
		return THREE_OF_A_KIND // 858 three-kind
	}
	if val > 1599 {
		return STRAIGHT // 10 straights
	}
	if val > 322 {
		return FLUSH // 1277 flushes
	}
	if val > 166 {
		return FULL_HOUSE // 156 full house
	}
	if val > 10 {
		return FOUR_OF_A_KIND // 156 four-kind
	}
	return STRAIGHT_FLUSH // 10 straight-flushes
}

func findit(key uint32) uint32 {
	low, high, mid := uint32(0), uint32(4887), uint32(0)

	for low <= high {
		mid = (high + low) >> 1 // divide by two
		if key < products[mid] {
			high = mid - 1
		} else if key > products[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return 0
}

func eval_5cards(c1, c2, c3, c4, c5 uint32) uint16 {
	var q uint32
	var s uint16

	q = (c1 | c2 | c3 | c4 | c5) >> 16

	// check for Flushes and StraightFlushes
	if (c1 & c2 & c3 & c4 & c5 & 0xF000) != 0 {
		return flushes[q]
	}

	// check for Straights and HighCard hands
	s = unique5[q]
	if s != 0 {
		return s
	}

	// let's do it the hard way
	q = (c1 & 0xFF) * (c2 & 0xFF) * (c3 & 0xFF) * (c4 & 0xFF) * (c5 & 0xFF)
	q = findit(q)

	return values[q]
}

func eval_5hand(hand []uint32) uint16 {
	var c1, c2, c3, c4, c5 uint32

	c1 = hand[0]
	c2 = hand[1]
	c3 = hand[2]
	c4 = hand[3]
	c5 = hand[4]

	return eval_5cards(c1, c2, c3, c4, c5)
}

func rankHand(cards []card) int {
	var hand [5]uint32
	for i, c := range cards {
		var cardBin uint32
		cardBin = uint32(primes[c.rank]) | (uint32(c.rank << 8)) | (1 << (15 - uint32(c.suit))) | (1 << (16 + uint32(c.rank)))
		hand[i] = cardBin
	}
	return int(eval_5hand(hand[:]))
}

func compare(hand1 []card, hand2 []card) int {
	if rankHand(hand1) < rankHand(hand2) {
		return 1
	} else if rankHand(hand1) > rankHand(hand2) {
		return -1
	} else {
		return 0
	}
}
