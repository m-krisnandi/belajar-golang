package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int {}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	// inisialisasi nilai map
	// var data map[string]int
	// data["one"] = 1
	// akan muncul error karena map belum diinisialisasi

	// data = map[string]int {}
	// data["one"] = 1
	// tidak akan muncul error karena map sudah diinisialisasi

	// cara horizontal
	var chicken1 = map[string]int {"januari": 50, "februari": 40}
	fmt.Println(chicken1)

	// cara vertical
	var chicken2 = map[string]int {
		"januari": 50,
		"februari": 40,
	}
	fmt.Println(chicken2)

	// inisialisasi map dengan make dan new
	// var chicken3 = map[string]int {}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	// Iterasi item map menggunakan for range
	var chicken6 = map[string]int {
		"januari": 50,
		"februari": 40,
		"maret": 34,
		"april": 67,
	}

	for key, value := range chicken6 {
		fmt.Println(key, "  \t:", value)
	}

	// delete item map
	var chicken7 = map[string]int {"januari": 50, "februari": 40}

	fmt.Println(len(chicken7))
	fmt.Println(chicken7)

	delete(chicken7, "januari")

	fmt.Println(len(chicken7))
	fmt.Println(chicken7)

	// deteksi item map
	var chicken8 = map[string]int {"januari": 50, "februari": 40}
	var value, isExist = chicken8["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	// deteksi item map dengan if short statement
	var chicken9 = map[string]int {"januari": 50, "februari": 40}
	if value, isExist := chicken9["mei"]; isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	// kombinasi map dan slice
	var chicken10 = []map[string]string {
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
		{"name": "chicken red", "gender": "male"},
	}

	for _, chicken := range chicken10 {
		fmt.Println(chicken["gender"])
		fmt.Println(chicken["name"])
	}

	// kombinasi map dan slice dengan key yang berbeda
	var chicken11 = []map[string]string {
		{"name": "chicken blue", "gender": "male"},
		{"address": "mangga street"},
	}

	for _, chicken := range chicken11 {
		fmt.Println(chicken["address"])
	}

}