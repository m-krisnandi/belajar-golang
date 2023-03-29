package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var argsRaw = os.Args
	fmt.Printf("-> %v\n", argsRaw)
	
	var args = argsRaw[1:]
	fmt.Printf("-> %v\n", args)

	// penggunaan flag
	// cara 1
	var name = flag.String("name", "Anonymous", "Your name")
	var age = flag.Int64("age", 20, "Your age")

	// cara 2
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")

	flag.Parse()
	// cara 1
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
	// cara 2
	fmt.Printf("gender\t: %s\n", data2)
	// go run main.go --name="John Doe" --age=30 --gender="female"
}