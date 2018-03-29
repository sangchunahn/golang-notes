package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// you can add another stuct and use the type
	// you just use the type as the field name
	contactInfo
}

func main() {
	// one of the first way to create and utilize the struct.
	// since we put alex first and the firstName first in the struct it will correspond with order
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// creates variable and assign the type struct to person
	// this is an empty object with empty string
	// var alex person

	// alex.firstName = "alex"
	// alex.lastName = "anderson"

	// fmt.Println(alex)
	// // %+v will print out all the fields that alex holds
	// fmt.Printf("%+v", alex)

	// reusing structs inside the same object
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// the & is an operator (%variable)
	// give us the memory address of the value this variable is pointing at
	// then we assign it to jimPointer
	// there is a shortcut way to do pointers
	// jimPointer := &jim
	// if you have a receiver function with *person it automatically does it for you
	jim.updateName("Jimmy")
	jim.print()
}

// this does not update the name for some reason (pointer)
// it doesnt update because the RAM points the the already created value in the address and copies it to another address so we are not updating the original
// we are using the copy to use the update function
// go is a pass by value language

// *person is a pointer that points at the person type
// we are looking at a pointer of person type of person
func (pointerToPerson *person) updateName(newFirstName string) {
	// the pointer *
	// give me the value this memory address is pointing at
	// the bottom pointer gets changed to a value
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// MAP is a collection of key/value pairs
