package main

import (
	"fmt"
	"os"
)

func main() {
	// defer selalu dieksekusi di akhir fungsi
	defer fmt.Println("Halo")
	fmt.Println("Selamat Datang")

	orderSomeFood("pizza")
	orderSomeFood("burger")

	// kombinasi defer dan IIFE
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		defer fmt.Println("halo 3")

		// supaya defer dieksekusi di akhir blok if
		// maka harus menggunakan IIFE
		func() {
			defer fmt.Println("halo 4")
		}()
	}

	fmt.Println("halo 2")

	// fungsi os.Exit() akan menghentikan eksekusi program
	// defer tidak akan dieksekusi kecuali menggunakan IIFE
	defer fmt.Println("halo 5")
	os.Exit(1)
	fmt.Println("Selamat Tinggal")
}

// defer akan dieksekusi meskipun ada return di fungsi
// defer akan dieksekusi secara berurutan
func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, silahkan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda adalah", menu)
}