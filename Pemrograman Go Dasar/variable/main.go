package main

import "fmt"

func main() {
	var firstName string = "John"
	var lastName string
	lastName = "Doe"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Printf("halo" + " " + firstName + " " + lastName + "!\n")
	fmt.Println("halo", firstName, lastName, "!")

	// Deklarasi tanpa tipe data, hanya bisa di dalam function
	var firstWord = "Hello"
	secondWord := "World"

	fmt.Println(firstWord, secondWord)

	// Deklarasi multiple variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	// bisa juga langsung di assign
	var fourth, fifth, sixth = "empat", "lima", "enam"

	// bisa juga langsung di assign tanpa var
	seventh, eighth, ninth := "tujuh", "delapan", "sembilan"

	// bisa juga memiliki tipe data yang berbeda
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)
	fmt.Println(seventh, eighth, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)

	// Variable Underscore
	// Variable yang tidak digunakan
	// Biasanya digunakan untuk mengabaikan nilai yang di return
	// oleh function
	_ = "Hello"
	_ = "World"

	// Deklarasi variable keyword new
	// Mengembalikan pointer
	name := new(string)

	fmt.Println(name)  // 0xc00000e018
	fmt.Println(*name) // ""

	// Deklarasi variable keyword make
	// Mengembalikan tipe data slice, map, dan channel
}