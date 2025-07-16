package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"https://www.google.com/", "https://stackoverflow.com", "https://amazon.com"}
	fmt.Println(links)
	c := make(chan string) //channel is used to receive message whether go routine executes
	for _, link := range links {
		go checkLink(link, c) // go routine (parallel execution splits the execution) it will be created (kind of thread) after each iteration and look website valid/not

	}
	// for l := range c {
	// 	go func(link string) { // func literal
	// 		time.Sleep(5 * time.Second) // pause endless loop by adding 5 sec delay
	// 		checkLink(link, c)
	// 	}(l)
	// } // can be used argument(pass-by-value) but in go 1.22 it has been fixed where we can use main and go routine

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}() // this has been fixed where we can use main routine variable
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
