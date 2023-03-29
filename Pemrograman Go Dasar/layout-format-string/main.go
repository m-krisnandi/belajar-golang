package main

import "fmt"

type student struct {
	name string
	height int
	age int32
	isGraduated bool
	hobbies []string
}

func main() {
	var data = student{
		name: "John Wick",
		height: 170,
		age: 21,
		isGraduated: false,
		hobbies: []string{"Fishing", "Hunting"},
	}
	
	// layout format %b
	fmt.Println("Binary: ", fmt.Sprintf("%b", data.age))

	// layout format %c
	fmt.Println("Character: ", fmt.Sprintf("%c", 1400))

	// layout format %d
	fmt.Println("Decimal: ", fmt.Sprintf("%d", data.age))

	// layout format %o
	fmt.Println("Octal: ", fmt.Sprintf("%o", data.age))

	// layout format %x
	fmt.Println("Hexadecimal: ", fmt.Sprintf("%x", data.age))

	// layout format %U
	fmt.Println("Unicode: ", fmt.Sprintf("%U", data.age))

	// layout format %T
	fmt.Println("Type: ", fmt.Sprintf("%T", data.age))

	// layout format %v
	fmt.Println("Value: ", fmt.Sprintf("%v", data))

	// layout format %+v
	fmt.Println("Value with field name: ", fmt.Sprintf("%+v", data))

	// layout format %#v
	fmt.Println("Go Syntax: ", fmt.Sprintf("%#v", data))

	// layout format %q
	fmt.Println("Quoted: ", fmt.Sprintf("%q", data.name))

	// layout format %s
	fmt.Println("String: ", fmt.Sprintf("%s", data.name))

	// layout format %t
	fmt.Println("Boolean: ", fmt.Sprintf("%t", data.isGraduated))

	// layout format %e
	fmt.Println("Exponential: ", fmt.Sprintf("%e", 0.123123123123))

	// layout format %E
	fmt.Println("Exponential: ", fmt.Sprintf("%E", 0.123123123123))

	// layout format %f
	fmt.Println("Float: ", fmt.Sprintf("%f", 0.123123123123))

	// layout format %F
	fmt.Println("Float: ", fmt.Sprintf("%F", 0.123123123123))

	// layout format %g
	fmt.Println("General: ", fmt.Sprintf("%g", 0.123123123123))

	// layout format %G
	fmt.Println("General: ", fmt.Sprintf("%G", 0.123123123123))

	// layout format %p
	fmt.Println("Pointer: ", fmt.Sprintf("%p", &data))

	// layout format %%
	fmt.Println("Percent: ", fmt.Sprintf("%%"))
}