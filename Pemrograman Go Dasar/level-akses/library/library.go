package library

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("My name is", name)
}

type Student struct {
	Name string
	Grade int
}

var Student2 = struct {
	Name string
	Grade int
}{}

func init() {
	Student2.Name = "ethan"
	Student2.Grade = 21

	fmt.Println("init di library")
}