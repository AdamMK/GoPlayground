package main

import "fmt"

//any concrete type that has a function getGreeting which returns string can also use type bot
//in exchange you can also call a function printGreeting or any function accepting bot
//to qualify your type has to implement all of the functions listed below
type bot interface {
	getGreeting() string
	//funName(list of allowed argument types) (list of allowed return types comma separated)

}
type englishBot struct {}
type spanishBot struct {}

//dont need value if not used but type of receiver still stays
func (englishBot) getGreeting() string  {
	//custom logic for English bot
	return "Hi There"
}

func (spanishBot) getGreeting() string  {
	//custom logic for Spanish bot
	return "Hola"
}

func main()  {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}
