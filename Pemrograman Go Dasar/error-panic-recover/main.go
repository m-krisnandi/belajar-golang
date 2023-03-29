package main

import (
	"errors"
	"fmt"
	"strings"
	// "strconv"
)

// custom error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	}
	return true, nil
}

// recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Aplikasi berjalan dengan baik")
	}
}

func main() {
	// var input string
	// fmt.Print("Type some number: ")
	// fmt.Scanln(&input)

	// var number int
	// var err error
	// number, err = strconv.Atoi(input)

	// if err == nil {
	// 	fmt.Println(number, "is a number")
	// } else {
	// 	fmt.Println(input, "is not a number")
	// 	fmt.Println(err.Error())
	// }

	// defer catch() // recover
	defer catch()

	// custom Error
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		fmt.Println(err.Error())

		// menggunakan panic
		panic(err.Error())
		// fmt.Println("end") // tidak akan dieksekusi
	}

	// pemanfaatan recover pada IIFE
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured:", r)
		} else {
			fmt.Println("Aplikasi berjalan dengan baik")
		}
	}()

	panic("Something went wrong")

	// penerapan recover pada IIFE
	// data := []string{"superman", "batman", "spiderman"}

	// for _, each := range data {
	// 	func() {
	// 		// recover untuk IIFE dalam perulangan
	// 		defer func() {
	// 			if r := recover(); r != nil {
	// 				fmt.Println("Panic occured:", r)
	// 			} else {
	// 				fmt.Println("Panic occured on looping", each, "| message: ", r)
	// 			}
	// 		}()
			
	// 		panic("some error happen")
	// 	}()
	// }
}