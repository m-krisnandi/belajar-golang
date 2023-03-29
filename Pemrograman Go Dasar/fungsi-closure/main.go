package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	var numbers = []int{2, 3, 2, 5, 6, 2, 8, 7, 6}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max)

	// immediately invoked function expression (IIFE)
	var numbers1 = []int {2, 3, 2, 5, 6, 2, 8, 7, 6}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers1 {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number:", numbers1)
	fmt.Println("filtered number:", newNumbers)

	// find max
	var max1 = 3
    var numbers2 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    var howMany, getNumbers = findMax(numbers2, max1)
    var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers2)
	fmt.Printf("max \t: %d\n\n", max1)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)

}

// closure sebagai nilai kembalian
func findMax (numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e < max {
			res = append(res, e)
		}
	}

	// return len(res), func() []int {
	// 	return res
	// }

	// atau simpan di variable
	var getNumbers = func() []int {
		return res
	}

	return len(res), getNumbers
}