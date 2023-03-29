package main

import "fmt"

func main() {
	// for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// for seperti while
	var counter = 1

	for counter <= 5 {
		fmt.Println("Angka", counter)

		counter++
	}

	// for tanpa argumen
	var counter2 = 1

	for {
		fmt.Println("Angka", counter2)

		counter2++

		if counter2 == 5 {
			break
		}
	}

	// break dan continue
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		if i > 6 {
			break
		}

		fmt.Println("Angka", i)
	}

	// perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println("")
	}

	// perulangan bersarang dengan label
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}

			fmt.Print("Matriks [", i, "][", j, "]", "\n")
		}
	}
}