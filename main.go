package main

import (
	"fmt"
	"os"
)

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 5)
	hand.print()
	fmt.Println("---")
	hand.shuffle()
	hand.print()
	hand.saveToFile("hand.txt")
	deckFromFile, error := newDeckFromFile("hand.txt")
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	fmt.Println("---")
	deckFromFile.print()
}
