package main

// import "fmt"

// import "fmt"

func main() {
	// cards := newCard()
	// cards = partition(cards, 0,3)
	// cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("Cards file")
	cards := newDeckFromFile("Cards file")
	
	cards.shuffle()
	cards.print()
}
