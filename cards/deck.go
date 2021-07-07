package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck which is a slice of strings
//deck extends string -deck has the same behaviour as slice of string
//treat this as class
type deck []string

//creating deck of cards of type deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//replace i and j with underscore as we not use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}


//This is a receiver function (fun that accepts receiver)
// any var of type deck has access to this function
// d is instance of deck
//usually receiver is one letter
func (d deck) print()  {
	for _, card := range d {
		fmt.Println(card)
	}
}

//func (d deck) hand() deck {
//	rand.Seed(time.Now().UnixNano())
//	min := 0
//	max := 52
//	var oneHand deck
//	for i := 1; i <= 5; i++ {
//		shuffledCard := rand.Intn(max - min) + min
//		oneHand = append(oneHand, d[shuffledCard])
//	}
//
//	return oneHand
//}

//function returns 2 deck values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//takesa receiver of deck and converts it to string
func (d deck) toString() string {
	//converting d to type string slice and passing to join function
	return strings.Join([]string(d),",")
}

//saves receiver of type deck and saves it to file
func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//getting deck from a file
func newDeckFromFile(filename string) deck{
	//readfile returns 2 types []byte and error - they need to assigned equally to vars
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//split bs to slice of strings
	s:= strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle()  {
	//used to generate different int64 everytime we need a source
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//rand.Seed(time.Now().UnixNano())
	for i:= range d {
 	newPosition := r.Intn(len(d)-1)
 	//swapping position of the both indexes
 	d[i], d[newPosition] = d[newPosition], d[i]
 	}
}