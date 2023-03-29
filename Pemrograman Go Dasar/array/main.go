package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "John"
	names[1] = "Wick"
	names[2] = "Ethan"
	names[3] = "Hunt"

	fmt.Println(names[0], names[1], names[2], names[3])

	// inisialasi nilai awal array
	var fruit = [4]string {"Apple", "Grape", "Banana", "Melon"}

	fmt.Println("Jumlah element \t\t", len(fruit))
	fmt.Println("Isi semua element \t", fruit)

	// inisialisasi array tanpa jumlah element
	var numbers = [...]int {2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah element \t:", len(numbers))

	// array multidimensi
	var numbers1 = [2][3]int { [3]int {3, 2, 3}, [3]int {2, 4, 3} }
	var numbers2 = [2][3]int { {3, 2, 3}, {2, 4, 3} } // tanpa deklarasi tipe data

	fmt.Println("numbers1 \t:", numbers1)
	fmt.Println("numbers2 \t:", numbers2)

	// perulangan array dengan for
	var fruits = [4]string {"Apple", "Grape", "Banana", "Melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	// perulangan array dengan for range
	var fruits1 = [4]string {"Apple", "Grape", "Banana", "Melon"}

	for i, fruit := range fruits1 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// variable underscore dalam for range
	var fruits2 = [4]string {"Apple", "Grape", "Banana", "Melon"}

	for _, fruit := range fruits2 {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// alokasi element array menggunakan make
	var fruits3 = make([]string, 2)
	fruits3[0] = "Apple"
	fruits3[1] = "Grape"

	fmt.Println(fruits3)
}