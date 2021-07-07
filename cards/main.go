package main

//can initialize before function
//var c string

func main()  {

	//declaring slice of type string
	//cards := deck{newCard(), "Ace of Diamonds"}
	cards := newDeck()
	//cards := newDeckFromFile("my_cars")
	//adding card to existing slice - this will return new slice
	//cards = append(cards, "Six of Spades")
	//cards = append(cards, "Six of Spades")

	//assignig to values of function return to new variables
	//hand, remainingDeck := deal(cards, 5)
	//fmt.Println("==================")
	//hand.print()
	//fmt.Println("==================")
	//remainingDeck.print()
	// for index i in cards range print index and element
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}
	//cards.print()
	//fmt.Println("==================")

	//greeting := "Hello"
	//fmt.Println([]byte(greeting))
	//hand := newDeck().hand()
	//
	//hand.print()

	//fmt.Println(cards.toString())
	cards.shuffle()
	cards.print()
	//cards.saveToFile("my_cards")
	//initialisation of value (by calling function which returns string)
	//c = newCard()
	//assign of new value to existing var
	//card = "Five of Diamonds"
	//fmt.Println(cards)

}

//no argument function with string return
//func newCard() string {
//	return "Five of Diamonds"
//}
