package main

import "fmt"

func main() {
	fmt.Printf("=== Maps ===\n\n")
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"
	fmt.Printf("%v\n\n", colors)

	fmt.Print("= delete(colors, white) =\n")
	delete(colors, "white")
	fmt.Printf("%v\n\n", colors)

	printMap(colors)

	fmt.Printf("======\n")
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
	fmt.Println()
}
