package main

import "fmt"

func main() {
	card := newDeck()
	d1, d2 := deal(card, 2)
	for _, c := range d1 {
		fmt.Println(c)
	}
	fmt.Println("-------------------")
	for _, c := range d2 {
		fmt.Println(c)
	}

	fmt.Println(card.toString())

	card.saveToFile("myFile")

	d := newDeckFromFile("myFile")
	fmt.Println(d)

}
