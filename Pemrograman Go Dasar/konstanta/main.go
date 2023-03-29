package main

import "fmt"

func main() {
	const firstName string = "John"
	const lastName string = "Doe"

	fmt.Println(firstName, lastName)

	// Multiple Constant
	const (
		firstName2 string = "John"
		lastName2 string = "Doe"
		square int = 10 * 10
		isToday bool = true
		numeric uint8 = 1
		floatNumber float32 = 1.2
	)

	const (
		today string = "Senin"
		sekarang
		isToday2 bool = true
	)

	// multiple constant dalam satu baris
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
}
	
