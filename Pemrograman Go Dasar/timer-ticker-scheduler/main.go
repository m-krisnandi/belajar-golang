package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// fmt.Println("start time.Sleep()")
	// time.Sleep(time.Second * 4)
	// fmt.Println("after 4 seconds time.Sleep()")

	// // scheduler menggunakan time.Sleep()
	// // scheduler ini akan berjalan setiap 1 detik
	// // dan akan berhenti jika program dihentikan

	// // for true {
	// // 	fmt.Println("Hello !!")
	// // 	time.Sleep(1 * time.Second)
	// // }

	// // fungsi time.NewTimer()
	// // fungsi ini akan mengembalikan 2 nilai
	// // 1. channel
	// // 2. timer
	// var timer3 = time.NewTimer(4 * time.Second)
	// fmt.Println("start time.NewTimer()")
	// <-timer3.C
	// fmt.Println("finish time.NewTimer()")

	// // fungsi time.AfterFunc()
	// // fungsi ini akan mengembalikan 1 nilai
	// // 1. channel
	// // fungsi ini akan menjalankan fungsi yang kita inginkan
	// // setelah waktu yang kita tentukan

	// var ch = make(chan bool)
	// time.AfterFunc(4*time.Second, func() {
	// 	fmt.Println("expired time.AfterFunc()")
	// 	ch <- true
	// })

	// fmt.Println("start time.AfterFunc()")
	// <-ch
	// fmt.Println("finish time.AfterFunc()")

	// // fungsi time.After()
	// <-time.After(4 * time.Second)
	// fmt.Println("expired time.After()")

	// // fungsi time.NewTicker()
	// // fungsi ini akan mengembalikan 2 nilai
	// // 1. channel
	// // 2. ticker

	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second)

	// go func ()  {
	// 	time.Sleep(10 * time.Second) // setelah 10 detik
	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop()
	// 		return
	// 	case t := <-ticker.C:
	// 		fmt.Println("Hello !!", t)
	// 	}
	// }

	// kombinasi timer & goroutine
	var timeout = 5
	var ch2 = make(chan bool)

	go timer(timeout, ch2)
	go watcher(timeout, ch2)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is correct")
	} else {
		fmt.Println("the answer is wrong")
	}

}

// kombinasi timer & goroutine
func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}


func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}
