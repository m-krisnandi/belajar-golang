package main

import "fmt"

func main() {
	var fruits = []string {"Apple", "Grape", "Banana", "Melon"}
	fmt.Println(fruits[0])

	//  hubungan slice dengan array
	var fruits1 = []string {"Apple", "Grape", "Banana", "Melon"}
	var newFruits = fruits1[0:2]

	fmt.Println(newFruits)

	// slice merupakan tipe data reference
	var fruits2 = []string {"Apple", "Grape", "Banana", "Melon"}

	var newFruits1 = fruits2[0:2]
	var newFruits2 = fruits2[1:3]

	var aFruits = newFruits1[0:1]
	var bFruits = newFruits2[0:3]

	fmt.Println(newFruits1)
	fmt.Println(newFruits2)
	fmt.Println(aFruits)
	fmt.Println(bFruits)

	// buah grape diubah menjadi pear
	bFruits[0] = "Pear"

	fmt.Println(newFruits1)
	fmt.Println(newFruits2)
	fmt.Println(aFruits)
	fmt.Println(bFruits)

	// slice menggunakan fungsi len dan cap
	var fruits3 = []string {"Apple", "Grape", "Banana", "Melon"}

	fmt.Println(len(fruits3))
	fmt.Println(cap(fruits3))

	var newFruits3 = fruits3[0:3]

	fmt.Println(len(newFruits3))
	fmt.Println(cap(newFruits3))

	// slice menggunakan fungsi append
	var fruits4 = []string {"Apple", "Grape", "Banana", "Melon"}
	var cFruits = append(fruits4, "Watermelon")

	fmt.Println(fruits4)
	fmt.Println(cFruits)

	var fruits5 = []string {"Apple", "Grape", "Banana"}
	var dFruits = fruits5[0:2]

	fmt.Println(cap(dFruits))
	fmt.Println(len(dFruits))

	fmt.Println(fruits5)
	fmt.Println(dFruits)

	var eFruits = append(dFruits, "Melon")

	fmt.Println(fruits)
	fmt.Println(dFruits)
	fmt.Println(eFruits)

	// slice menggunakan fungsi copy
	// copy(dest, src)
	dst := make([]string, 3)
	src := []string {"Watermelon", "Pinnaple", "Apple", "Grape"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	dst2 := []string {"potato", "potato", "potato"}
	src2 := []string {"Watermelon", "Pinnaple"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2)
	fmt.Println(src2)
	fmt.Println(n2)

	// akses element slice dengan 3 index
	var fruits6 = []string {"Apple", "Grape", "Banana", "Melon"}
	var newFruits4 = fruits6[0:2]
	var newFruits5 = fruits6[0:2:2]

	fmt.Println(fruits6)
	fmt.Println(len(fruits6))
	fmt.Println(cap(fruits6))

	fmt.Println(newFruits4)
	fmt.Println(len(newFruits4))
	fmt.Println(cap(newFruits4))

	fmt.Println(newFruits5)
	fmt.Println(len(newFruits5))
	fmt.Println(cap(newFruits5))

}