package main

import "fmt"

// Define Embedding Struct
type contactInfo struct {
	email string
	zipCode int
}

// Define Person Struct
type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	// Declare Person Struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	
	jim.updateName("jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*
# Think struct vs collection

# Tell go what field the person struct has
- firstName string
- lastName string

# Create a new value of type person
- firstname = "Alex"
- lastName = "Anderson"

# Pass by Value?
jim := person{}

RAM
Address | Value
0000
0001    | person{firstName: "Jim"....} <--- jim

*/