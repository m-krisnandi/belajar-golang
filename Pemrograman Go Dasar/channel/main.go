package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("John Wick")
	go sayHelloTo("Ethan Hunt")
	go sayHelloTo("Indiana Jones")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	// channel sebagai parameter
	for _, each := range []string{"wick", "hunt", "indiana"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}

	// Eksekusi Goroutine Pada IIFE
	var messages2 = make(chan string)

	go func(who string) {
    	var data = fmt.Sprintf("hello %s", who)
    	messages2 <- data
	}("wick")

	var message = <-messages2
		fmt.Println(message)
	}

// channel sebagai parameter
func printMessage(what chan string) {
	fmt.Println(<-what)
}


