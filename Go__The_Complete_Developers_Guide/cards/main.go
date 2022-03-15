package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

	hand, deck := deal(cards, 5)
	fmt.Println(hand.toString())
	fmt.Println(deck.toString())

	deck.saveToFile("my_cards")
	cards = newDeckFromFile("my_cards")
	cards.print()

	cards.shuffle()
	cards.print()
}
