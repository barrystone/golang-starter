package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// // fmt.Println(os.Args)
	// content := readFromFile(os.Args[1])
	// fmt.Println(content)

	// file, err := os.Open("files/" + os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

}

// func readFromFile(fileName string) string {
// 	filePath := filepath.Join("./files", fileName)
// 	bs, err := os.ReadFile(filePath)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		os.Exit(1)
// 	}

// 	return string(bs)
// }
