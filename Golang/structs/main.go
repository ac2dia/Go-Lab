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
	contactInfo // contactInfo contactInfo 와 동일
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
	(*p).firstName = newFirstName
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


# &variable
- value의 메모리 주소를 지칭

# *pointer
- 해당 메모리 주소가 가르키는 값을 지칭

func (pointerToPerson *person) updateName() {
	*pointerToPerson
}

- *person = type description
- *pointerToPerson = operator

# address vs value
address | value
00001   | person{firstName: "Jim"...}

- Turn address into value with *address
- Turn value into address with &value

# Reference vs Value Types
## Value Types
- int, float, string, bool, structs
  - Use pointers to change these things in a function

## Reference Types
  - slices, maps, channels, pointers, functions
  - Don't worry about pointers with these

# slice vs array

slice
- ptr to head (array[0] 을 가리킴)
- capacity
- length

array
"Hi", "There", "how", "are", "you?"

*/