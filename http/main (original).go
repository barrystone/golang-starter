package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	resp, err := http.Get("http://google.com")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	// fmt.Println(resp)

// 	// fmt.Printf("\n=== Response (1) ===\n\n")
// 	// bs := make([]byte, 99999)
// 	// resp.Body.Read(bs)
// 	// fmt.Println(string(bs))

// 	fmt.Printf("\n=== Response (2) ===\n\n")
// 	_, err = io.Copy(os.Stdout, resp.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// }
