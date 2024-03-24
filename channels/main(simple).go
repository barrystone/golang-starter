package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	links := []string{
// 		"https://www.google.com",
// 		"https://www.facebook.com",
// 		"https://www.youtube.com",
// 		"https://www.amazon.com",
// 		"https://www.golang.org",
// 	}

// 	c := make(chan string)

// 	for _, link := range links {
// 		go checkLink(link, c)
// 	}

// 	for i := 0; i < len(links); i++ {
// 		fmt.Println(<-c)
// 	}

// 	// fmt.Println(<-c)
// }

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		c <- "Might be dowm I think!"
// 		return

// 	}

// 	fmt.Println(link, "is up!")
// 	c <- "Yes it is up!"
// }
