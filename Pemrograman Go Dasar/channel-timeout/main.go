package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan int) {
	loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`Received data: `, data, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout. No activity for 5 second")
			break loop
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	
	var messages = make(chan int)

	go sendData(messages)
	retreiveData(messages)
}