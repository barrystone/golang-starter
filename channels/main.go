package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// This is an infinite loop, same as the for loop below
	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		// go checkLink(l, c)
		go func(link string) {
			// Notice it, Don't use sleep in the main thread, because it will block the main thread, and the program will not be able to receive the value from the channel
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return

	}

	fmt.Println(link, "is up!")
	c <- link
}
