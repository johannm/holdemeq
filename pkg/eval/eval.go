package eval

import (
	"math/rand"
	"time"
)

const StraightFlush = 1
const FourOfAKind = 2
const FullHouse = 3
const Flush = 4
const Straight = 5
const ThreeOfAKind = 6
const TwoPair = 7
const OnePair = 8
const HighCard = 9

var value_str = []string{
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
		return HighCard // 1277 high card
	}
	if val > 3325 {
		return OnePair // 2860 one pair
	}
	if val > 2467 {
		return TwoPair // 858 two pair
	}
	if val > 1609 {
		return ThreeOfAKind // 858 three-kind
	}
	if val > 1599 {
		return Straight // 10 straights
	}
	if val > 322 {
		return Flush // 1277 flushes
	}
	if val > 166 {
		return FullHouse // 156 full house
	}
	if val > 10 {
		return FourOfAKind // 156 four-kind
	}
	return StraightFlush // 10 straight-flushes
}

func CalculateHoldemEquity(hand1, hand2, board []Card, n int) (int, int, int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	win, lose, draw := 0, 0, 0

	for i := 0; i < n; i++ {
		d := createDeck()
		d.shuffle(r1)

		// remove hole cards and board from deck
		for _, c := range append(append(hand1, hand2...), board...) {
			d.remove(c)
		}

		// Deal out the rest of the cards
		dealtBoard := make([]Card, len(board))
		copy(dealtBoard, board)
		for i := 0; i < 5-len(board); i++ {
			dealtBoard = append(dealtBoard, d.dealOne())
		}

		maxhand1 := findMaxhand(append(hand1, dealtBoard...))
		maxhand2 := findMaxhand(append(hand2, dealtBoard...))

		won := Compare(maxhand1, maxhand2)
		if won > 0 {
			win++
		} else if won < 0 {
			lose++
		} else {
			draw++
		}
	}
	return win, lose, draw
}


func findMaxhand(cards []Card) []Card {
	hand := make([]Card, 5)
	maxHand := make([]Card, 5)
	bestRank := 100000
	for _, combo := range Perm7 {
		for i, cardIndex := range combo {
			hand[i] = cards[cardIndex]
		}
		r := RankHand(hand)
		if r < bestRank {
			bestRank = r
			copy(maxHand, hand)
		}
	}
	return maxHand
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

func RankHand(cards []Card) int {
	var hand [5]uint32
	for i, c := range cards {
		var cardBin uint32
		cardBin = uint32(primes[c.Rank]) | (uint32(c.Rank << 8)) | (1 << (15 - uint32(c.Suit))) | (1 << (16 + uint32(c.Rank)))
		hand[i] = cardBin
	}
	return int(eval_5hand(hand[:]))
}

func Compare(hand1 []Card, hand2 []Card) int {
	if RankHand(hand1) < RankHand(hand2) {
		return 1
	} else if RankHand(hand1) > RankHand(hand2) {
		return -1
	} else {
		return 0
	}
}
