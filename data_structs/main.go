package main

import "fmt"

// Structs are like a dictionary in Python.

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName:   "Jim",
		lastName:    "Party",
		contactInfo: contactInfo{email: "jim_party@testing.com", zipCode: 94000},
	}

	jim.updateFirstName("Jimmy")
	jim.updateLastName("Pants")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) updateLastName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}
