package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
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
	fmt.Printf("= fmt.Printf(\"%%+v\", barry) =\n")
	fmt.Printf("%+v", barry)

	fmt.Println("\n= fmt.Println(barry) =")
	fmt.Println(barry)

	james := person{
		firstName: "Paul",
		lastName:  "George",
		contact: contactInfo{
			email:   "paul@example.com",
			zipCode: 13,
		},
	}

	fmt.Printf("= fmt.Printf(\"%%+v\", james) =\n")
	fmt.Printf("%+v", james)

	fmt.Printf("\n======\n")

}
