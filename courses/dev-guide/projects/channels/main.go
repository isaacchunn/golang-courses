package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//Serial link problem
	// for _, link := range links {
	// 	checkLink(link)
	// }

	//Create a channel of type strign
	c := make(chan string)
	//Using Go routines
	for _, link := range links {
		go checkLink(link, c)
	}
	//Wait for value to be sent to channel and then log it
	//for i := 0; i < len(links); i++ {
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //This is actually a blocking call
	if err != nil {
		fmt.Println(link, "might be down!")
		//Push back link name into channel
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
