package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactinfo
}

type contactinfo struct {
	email string
	postcode int
}

func main()  {
	alex := person{firstName: "Alex",lastName: "Anderson"}
	fmt.Println(alex)
	//default to 2x empty strings
	var adam person
	//name of variable + value
	fmt.Printf("%+v", adam)
	fmt.Println()
	adam.firstName = "Adam"
	adam.lastName = "Smith"
	fmt.Println(adam)

	jim := person{
		firstName: "Jim",
		lastName: "Par",
		contactinfo: contactinfo{
			email: "jim@gmail.com",
			//each line has to have comma even the last one
			postcode: 59220,
		},
	}
	jim.print()
	fmt.Println()
	//gives the memory address of the value this variable is pointing at
	//jimPointer is a memory address
	jimPionter := &jim
	jimPionter.updateName("Dan")
	jim.print()
	fmt.Println()

	//as long as function takes a pointer go will know that it want to pull the pointer rather then value of jim
	jim.updateName("New guy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

//* in front on the type means we are working with a pointer(memory) to a type person
// *person will only take a memory address (not the actuall value) e.g jimPointer
func (pointerToPerson *person) updateName(newFirstName string)  {
	//* in front of variable means we want to manipulate the value(in the memory) the pointer is referencing
	//now we have pointer to memory and we pull a value from it - then we can change it
	(*pointerToPerson).firstName = newFirstName
}

//& before value - pull the pointer from where the value is
//* before pointer - pull the value from where the pointer is


//Slice works differently then stuckt - no need to use pointers for slice
// value Types - needs pointers to change value:
//int,float, string, bool, structs
//reference types - no need to worry about pointers:
//slices, maps, channels, pointers, functions

