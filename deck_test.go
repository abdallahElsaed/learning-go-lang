package main

import "testing"

//When you want to make test you must create file (..._test.go)
/* Test in newDeck() function
	1- check this function return the right number of element of a slice
	2- the first element is (Spades Ace)
	3- the last element is (Clubs Four)
*/

func TestNewDeck( t *testing.T){
	d := newCard()

	// first case
	if len(d) != 16 {
		t.Errorf("The NUmber of Cards expected is 16 %v" , len(d))
	}
	// second case
	if (d[0] != "Spades Ace"){
		t.Errorf("The first element of Cards expected is Spades Ace and the coming is %v" , d[0])
	}
	// third case
	if (d[len(d) -1] != "Clubs Four"){
		t.Errorf("The last element of Cards expected is Clubs Four and the coming is %v" , d[len(d) -1])
	}

}