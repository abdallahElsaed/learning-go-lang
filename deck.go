package main

import (
	"fmt"
	"os"
	"strings"
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

/* I want to save Deck slice in Hard Disk 
	1. Slice of string must convert to Slice of byte []byte
	2. to convert it you must convert []string to string before
*/

// 1. convert to string
func (d deck) toString () string{
	return strings.Join(d , ",\n ")
}

// 2. save to hard disk 

func (d deck) saveToFile(filename string) error {
	return os.WriteFile( filename , []byte(d.toString()),0777)
}

