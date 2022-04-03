package ungar

func Evaluate(cards [7]uint16) uint16 {
	// establish some variables
	var totalValue uint16
	var suitProduct uint32 = 1
	var valProduct uint64 = 1

	for i := range cards {
		/*
		   First 7 MSB of a card is a prime number of its value
		   [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41]
		   and the last 3 LSB are primes of the cards suit
		   [2, 3, 5, 7]

		   So first thing we do is calculate product of primes to get
		   unique keys for lookup map for both suit and value product
		*/
		suitProduct *= uint32(cards[i] & 0x7)
		valProduct *= uint64(cards[i] >> 3)
	}

	flush, ok := Suited[suitProduct]
	if ok {
		/*
		   Now it makes sense to check if there's
		   a straight among suited cards
		*/
		suitedCards := []uint16{}
		suitedCardsProduct := uint64(1)
		for _, card := range cards {
			if byte(card&0x7) == flush {
				suitedCards = append(suitedCards, card)
				suitedCardsProduct *= uint64(card >> 3)
			}
		}
		straight, ok := Values[suitedCardsProduct]
		if ok {
			/*
			   We have straight flush so we get just regular staights value
			   and put it on top of the highest quads value. This way our
			   straight flushes will still all be in order but be valued higher
			   then any other combination
			*/
			totalValue = straight - 1599
		} else {
			/*
			   There's no straight flush on board and it makes no sense
			   to check for quads or full houses as at this point its not
			   possible that these combinations might be on board. So we
			   just return same value as in high card combination put on
			   top of highest straight value
			*/
			totalValue = Flushes[suitedCardsProduct] - 5863
		}
	} else {
		/*
		   At this point we just return precomputed value from a lookup map
		*/
		totalValue = Values[valProduct]
	}
	return totalValue
}
