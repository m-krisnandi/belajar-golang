package main

import "fmt"

func main() {
	// 1. Arithmetic operators
	// + - * / % ++ --
	// 2. Perbandingan operators
	// == != > < >= <=
	// 3. Logical operators
	// && || !

	// penggunaan aritmatika operator
	var value = (((2+6)%3)*4 - 2) / 3

	// penggunaan Perbandingan operator
	var isEqual = value == 2

	// penggunaan logical operator
	var isGreater = value > 5 && value < 10
	var isLess = value < 5 || value > 10

	fmt.Printf("nilai %d (%t) \t", value, isEqual)
	fmt.Printf("nilai lebih besar dari 5 dan kurang dari 10 (%t) \n", isGreater) // false
	fmt.Printf("nilai lebih kecil dari 5 atau lebih besar dari 10 (%t) \n", isLess) // true

	// penggunaan logical operator 2
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t(%t) \n", leftReverse)
}