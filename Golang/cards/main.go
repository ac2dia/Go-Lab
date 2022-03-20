package main

func main() {

	// 15. Variable Declarations
	// // var card string = "Ace of spades"
	// card := "Ace of spades"
	// card = "Five of Diamonds"

	// 16. Functions and Return Types
	// card := newCard()
	// fmt.Println(card)

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

/*
Variable Structure
- var name string = "ac2dia"

Dynamic Types vs Static Types
- JavaScript, Python ...
- C++, Java, Go ...

Basic Go Types
- bool, string, int, float64 ...

------------
Function Structure
- func newCard() string

------------
data structure for Handling list
1. Array
- Fixed length list of things

2. Slice
- An array that can grow or shrink

for index, card := range cards {
	fmt.Println(i, card)
}

------------


*/
