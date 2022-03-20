package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
Receiver => (d deck)
- receiver 표기의 경우 주로 한 글자, 두 글자와 같은 약어로 사용

Slice Range Syntax
- slice[startIndexIncluding:endIndexNotIncluding]
- slice[:3]
  - 0, 1, 2
- slice[3:]
  - 3 ...

*/
