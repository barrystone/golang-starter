package main

import "fmt"

type numbersArray []int

func main() {
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers := numbersArray{}

	for i := range 11 {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {
		if number%2 == 1 {
			fmt.Println(number, "is Odd")
		} else {
			fmt.Println(number, "is Even")
		}
	}
}
