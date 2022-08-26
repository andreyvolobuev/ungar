package ungar

import "strings"

func Convert(input string) [7]uint16 {
	var cards [7]uint16
	hand := strings.Split(input, " ")
	for i, card := range hand {
		// converting string representation of cards to prime ints
		cards[i] = Deck[card]
	}
	return cards
}
