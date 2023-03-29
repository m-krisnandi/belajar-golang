package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// tipe data numerik desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// tipe data boolean
	var exist bool = true

	fmt.Printf("exist? %t \n", exist)

	// tipe data string
	var message string = "Halo"

	fmt.Printf("message: %s \n", message)

	// using backtick
	var message2 string = `Halo nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang"`

	fmt.Println(message2)
}