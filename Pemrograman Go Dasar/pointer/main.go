package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) \t:", numberA)
	fmt.Println("numberA (address) \t:", &numberA)

	fmt.Println("numberB (value) \t:", *numberB)
	fmt.Println("numberB (address) \t:", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value) \t:", numberA)
	fmt.Println("numberA (address) \t:", &numberA)
	fmt.Println("numberB (value) \t:", *numberB)
	fmt.Println("numberB (address) \t:", numberB)

	fmt.Println("")

	// parameter pointer
	var number = 4
	fmt.Println("before :", number)

	change(&number, 10)
	fmt.Println("after :", number)
}

func change(original *int, value int) {
	*original = value
}