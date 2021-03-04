package main

import "strings"

type deck []string

func newDeck() deck {
	return deck{"a", "b", "c", "d"}
}

func deal(d deck, size int32) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}
