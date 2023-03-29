package main

import "fmt"

func Sum[V int | float32 | float64](numbers []V) V {
	var total V
	for _, number := range numbers {
		total += number
	}
	return total
}

// non-generic function
func SumNumbers1(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}


// generic function 
func SumNumbers2[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// generic function with constraints
type Number interface {
	int64 | float64
}

func SumNumbers3[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// struct generic
type UserModel[T int | float64] struct {
	Name string
	Scores []T
}

func (m *UserModel[int]) SetScoresA(scores []int) {
	m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}

func main() {
	total1 := Sum([]int {1, 2, 3, 4, 5})
	fmt.Println("total: ", total1)

	total2 := Sum([]float32 {1.1, 2.2, 3.3, 4.4, 5.5})
	fmt.Println("total: ", total2)

	total3 := Sum([]float64 {1.1, 2.7, 3.8, 8.4, 9.5})
	fmt.Println("total: ", total3)

	ints := map[string]int64{"first": 32, "second": 64, "third": 128}
	floats := map[string]float64{"first": 32.1, "second": 64.2, "third": 128.3}

	// generic function
	fmt.Printf("Generic Sums with constraints: %v and %v\n", SumNumbers2(ints), SumNumbers2(floats))
	
	// generic function with constraints
	fmt.Printf("Generic Sums without constraints: %v and %v\n", SumNumbers3(ints), SumNumbers3(floats))

	// struct generic
	var m1 UserModel[int]
	m1.Name = "John"
	m1.Scores = []int{1, 2, 3, 4, 5}
	m1.SetScoresA([]int{1, 2, 3, 4, 5})
	fmt.Println("scores: ", m1.Scores)

	var m2 UserModel[float64]
	m2.Name = "Wick"
	m2.Scores = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	m2.SetScoresB([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
	fmt.Println("scores: ", m2.Scores)
}