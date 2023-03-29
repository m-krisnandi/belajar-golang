package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// method biasa
func (s student) changeName(name string) {
	fmt.Println("before :", s.name)
	s.name = name
}

// method pointer
func (s *student) changeName2(name string) {
	fmt.Println("before :", s.name)
	s.name = name
}

func main() {
	// akses method dari variable object pointer
	var s1 = student{"john wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan: ", name)

	fmt.Println("s1 before :", s1.name)

	s1.changeName("jason bourne")
	fmt.Println("s1 after changeName :", s1.name)

	s1.changeName2("ethan hunt")
	fmt.Println("s1 after changeName2 :", s1.name)

	// akses method dari variable object pointer
	var s2 = &student{"ethan hunt", 21}
	s2.sayHello()
}