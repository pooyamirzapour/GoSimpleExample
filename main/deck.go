package main

type deck []string

func newDeck() deck {
	return deck{"a", "b", "c", "d"}
}

func deal(d deck, size int32) (deck, deck) {
	return d[:size], d[size:]
}
