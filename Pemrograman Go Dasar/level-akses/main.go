package main

import (
	. "level-akses/library"
	f "fmt"
)

func main() {
	SayHello("wick")

	var s1 = Student{Name:"wick", Grade: 21}
	f.Println("name:", s1.Name)
	f.Println("grade:", s1.Grade)

	// akses partial.go
	sayHello("ethan")

	// akses library.go function init()
	f.Printf("Name : %s\n", Student2.Name)
	f.Printf("Grade : %d\n", Student2.Grade)

}