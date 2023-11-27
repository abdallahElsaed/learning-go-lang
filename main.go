package main

import "fmt"

func main() {
	cards := newCard()
	// cards = partition(cards, 0,3)
	// cards.print()
	fmt.Println(cards.toString())
}
