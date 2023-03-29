package main

import "fmt"

func main() {
	var nilai int = 85
	// fmt.Print("Masukkan nilai: ")
	// fmt.Scan(&nilai)

	if nilai >= 80 {
		fmt.Println("Nilai A")
	} else if nilai >= 70 {
		fmt.Println("Nilai B")
	} else if nilai >= 60 {
		fmt.Println("Nilai C")
	} else if nilai >= 50 {
		fmt.Println("Nilai D")
	} else {
		fmt.Println("Nilai E")
	}

	// variable temporary pada kondisi
	var point = 8840.0
	
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// switch case
	var point2 = 8

	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better")
		}
	}

	// switch case dengan kondisi
	var point3 = 7

	switch {
	case point3 == 8:
		fmt.Println("perfect")
	case (point3 < 8) && (point3 > 3):
		fmt.Println("awesome")
		// fallthrough digunakan untuk mengeksekusi case selanjutnya
		fallthrough
	case point3 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better")
		}
	}

	// Seleksi Kondisi bersarang
	var point4 = 10

	if point4 > 7 {
		switch point4 {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if point4 == 5 {
			fmt.Println("not bad")
		} else if point4 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point4 == 0 {
				fmt.Println("try harder")
			}
		}
	}
}