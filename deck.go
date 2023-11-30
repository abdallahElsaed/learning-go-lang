package main

import (
	"math/rand"
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

func newDeckFromFile( filename string) deck {
	bs , err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:" , err)
		os.Exit(1)
	}
	// convert to slice 
	s := strings.Split(string(bs) , ",\n ")
	return deck(s) 
}

/*
	I want to make shuffling for slice or cheng position for all slice element
		1. I will loop for slice 
		2. make random number 
		3. cheng element of this slice by his index and random number
*/

func (d deck) shuffle(){
	for  i := range d{
		newPosition := rand.Intn(len(d) -1)
		d[i] , d[newPosition] = d[newPosition] , d[i]
	}
}
