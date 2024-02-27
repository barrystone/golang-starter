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
	fmt.Printf("=== Define struct ===\n\n")
	var barry person
	// barry := person{firstName: "Barry", lastName: "Stone"}
	// // It also works by following the order of the struct
	// barry := person{"Barry", "Stone"}
	barry.firstName = "Barry"
	barry.lastName = "Stone"

	// "%+v" prints all the fields and Values of the struct
	fmt.Printf("= barry.printPerson() =\n")
	barry.printPerson()
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

	fmt.Printf("= paul.printPerson() =\n")
	paul.printPerson()

	fmt.Printf("======\n")
	fmt.Printf("=== Struct with Pointers ===\n\n")

	fmt.Printf("= paul.printPerson(): paul.updateNameWithNoPointer(\"James\") =\n")
	// This updates the firstName of the "copy" of paul
	paul.updateNameWithNoPointer("James")
	paul.printPerson()

	fmt.Printf("= paul.printPerson(): paulPointer.updateName(\"James\") =\n")
	paulPointer := &paul
	// This updates the firstName of the original paul
	paulPointer.updateName("James")
	paul.printPerson()

	fmt.Printf("= paul.printPerson(): paul.updateName(\"Russell\") =\n")
	// Easy way to use pointer with golang shortcut
	paul.updateName("Russell")
	paul.printPerson()

	fmt.Printf("======\n")
	fmt.Printf("=== Gotchas with Pointers ===\n\n")

	fmt.Printf("= updateSlice =\n")
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Printf("%+v\n\n", mySlice)

	fmt.Printf("= updateValue =\n")
	name := "Barry"
	updateValue(name)
	fmt.Printf("%+v\n\n", name)

	fmt.Printf("======\n")
}

func (p person) printPerson() {
	fmt.Printf("%+v\n\n", p)
}

func (p person) updateNameWithNoPointer(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func updateValue(n string) {
	n = "Rock"
}
