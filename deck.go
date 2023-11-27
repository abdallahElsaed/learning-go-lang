package main

import (
	"fmt"
)

type deck []string 


func  newCard() deck {
	cardSuits := []string{"Spades","Hearts","Diamonds","Clubs",}
	cardValues := []string{"Ace","Two","Three","Four",}
	cards :=  deck{} // what is difference when use this ([]string {})?

	for  _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards , suit + " " + value)
		}
	}
	return cards
}
func (d deck) print(){
	for  i, card := range d {
		fmt.Println(i , card)
	}
}

func partition(d deck, first int , second int) deck{
	return d[first:second]
}
