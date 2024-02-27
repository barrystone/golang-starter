package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	// // Same as above
	// contactInfo contactInfo
}

func main() {
	fmt.Println("=== Define struct ===")
	var barry person
	// barry := person{firstName: "Barry", lastName: "Stone"}
	// // It also works by following the order of the struct
	// barry := person{"Barry", "Stone"}
	barry.firstName = "Barry"
	barry.lastName = "Stone"

	// "%+v" prints all the fields and Values of the struct
	fmt.Printf("= barry.print() =\n")
	barry.print()
	fmt.Println("= fmt.Println(barry) =")
	fmt.Println(barry)

	paul := person{
		firstName: "Paul",
		lastName:  "George",
		contactInfo: contactInfo{
			email:   "paul@example.com",
			zipCode: 13,
		},
	}

	fmt.Printf("= paul.print() =\n")
	paul.print()

	fmt.Println("=== Struct with Pointers ===")
	fmt.Printf("= paul.print(): updateNameWithNoPointer =\n")
	// This updates the firstName of the "copy" of paul
	paul.updateNameWithNoPointer("James")
	paul.print()

	fmt.Printf("= paul.print(): updateName =\n")
	paulPointer := &paul
	// This updates the firstName of the original paul
	paulPointer.updateName("James")
	paul.print()

	fmt.Printf("======\n")
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateNameWithNoPointer(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
