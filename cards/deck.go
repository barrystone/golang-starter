package main

import "fmt"

// Creat a new type of 'deck'
// which is a slice of strings
type deck []string

// Receiver function
// 'd' is a variable(like 'this' or 'self'),
// excetly like cards put in there
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}