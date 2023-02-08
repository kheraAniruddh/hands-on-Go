package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println()
	hand1, remainingDeck := cards.deal(5)

	hand1.print()
	fmt.Println()
	remainingDeck.print()

	fmt.Println(hand1.toString())
	fmt.Println([]byte(hand1.toString()))
	hand1.saveToFile("saveFromGo.txt")
	readFromFile("saveFromGo.txt").print()

	cards.print()
	cards.shuffle()
	cards.print()
}
