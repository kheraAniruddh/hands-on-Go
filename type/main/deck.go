package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

// receiver function
// d-> instance of type deck (by convention we use the initials of the "type")
// deck -> type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

// we add receiver on an instance, here we are creating one
func newDeck() deck {
	cards := deck{}

	suites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range suites {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suite)
		}
	}
	return cards
}

func (d deck) deal(handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFromFile(fileName string) deck {
	bs, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	arr := strings.Split(string(bs), ",")
	return deck(arr)
}

func (d deck) shuffle() {
	for i := range d {
		j := rand.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}
