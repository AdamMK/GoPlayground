package main

import "fmt"

func main()  {
	//declaring map with key and value of type string
	colors := map[string]string{
		"red":"#ff0000",
		"green":"#4bf745",
	}

	//second way of declaring map - no instatiation
	var colors2 map[string]string

	//third way
	colors3 := make(map[string]string)

	//add element
	colors["white"] = "#ff0033"

	//checking elelent
	fmt.Println(colors["red"])

	//delete element of the map
	delete(colors, "green")

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	printMap(colors)
}

func printMap(c map[string]string)  {
	//for key,value in the c map
	for color,hex:= range c {
		fmt.Println("Hex code for", color,"is",  hex)
	}
}
