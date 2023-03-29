package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string {"apple", "grape", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var data map[string]interface{}

	data = map[string]interface{} {
		"name": "Ethan Hunt",
		"grade": 2,
		"breakfast": []string {"apple", "grape", "banana"},
	}

	fmt.Println(data)

	// tipe alias any
	var data2 map[string]any

	data2 = map[string]any {
		"name": "Ethan Hunt",
		"grade": 2,
		"breakfast": []string {"apple", "grape", "banana"},
	}

	fmt.Println(data2)

	// casting variable interface kosong
	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is: ", number)

	secret = []string {"apple", "grape", "banana"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")

	// casting variable interface kosong ke objek pointer
	type person struct {
		name string
		age int
	}

	var secret2 interface{} = &person{name: "Ethan Hunt", age: 40}
	var name = secret2.(*person).name
	fmt.Println(name)

	// kombinasi slice, map, dan interface kosong
	var person2 = []map[string]interface{} {
		{"name": "wick", "age": 21},
		{"name": "ethan", "age": 40},
		{"name": "bourne", "age": 37},
	}

	for _, each := range person2 {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits2 = []interface{} {
		map[string]interface{} {
			"name": "apple",
			"price": 10000,
		},
		[]string {"apple", "grape", "banana"},
	}

	for _, each := range fruits2 {
		fmt.Println(each)
	}
}