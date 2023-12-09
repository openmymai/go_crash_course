package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender	string
	age																int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i'm " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// strconv.Itoa string converter, nothing too complex

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}
func main() {
	// Init person with struct
	// person1 := Person{firstName.....}

	person1 := Person{"Samantha", "Rico", "NYC", "f", 23}
	person2 := Person{"Bob", "Johnson", "LA", "m", 42}

	fmt.Println(person1)

	// Get single field
	// person1.age++
	// fmt.Println(person1.age)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Jackson")
	fmt.Println(person2.greet())
}