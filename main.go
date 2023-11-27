package main 

func main(){
    cards := newCard()
    cards = partition(cards, 0,3)
    cards.print()
}



