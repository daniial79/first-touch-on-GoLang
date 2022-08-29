package main

func main() {
	deck := newDeck()

	deck.shuffle()

	hand, remainingDeck := deal(deck, 10)

	hand.print()
	remainingDeck.print()

	hand.saveToFile("hand.txt")
	remainingDeck.saveToFile("remaining-deck.txt")
}
