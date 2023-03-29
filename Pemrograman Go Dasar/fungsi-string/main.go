package main

import (
	"fmt"
	"strings"
)

func main() {
	// deteksi apakah string(params ke 2) ada di string(params ke 1)
	var isExist = strings.Contains("Hello World", "Hello")
	fmt.Println(isExist) // true

	// deteksi string(params 1) diawali string tertentu(params 2)
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	// deteksi string(params 1) diakhiri string tertentu(params 2)
	var isSuffix1 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix1) // true

	var isSuffix2 = strings.HasSuffix("john wick", "wi")
	fmt.Println(isSuffix2) // false

	// menghitung jumlah string tertentu(params 2) di string(params 1)
	var howMany = strings.Count("hello world", "l")
	fmt.Println(howMany) // 3

	// mencari index string tertentu(params 2) di string(params 1)
	var index = strings.Index("hello world", "l")
	fmt.Println(index) // 2

	// replace bagian string dengan string tertentu
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // bonana

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // bonona

	var newText3 = strings.Replace(text, find, replaceWith, -1) // -1 untuk mengganti semua
	fmt.Println(newText3) // bonono

	// mengulang string tertentu(params 1) sebanyak params ke 2
	var str = strings.Repeat("na", 4)
	fmt.Println(str) // nananana

	// memsiahkan string(params 1) dengan tanda pemisah(params 2)
	var string1 = strings.Split("the cat in the hat", " ")
	fmt.Println(string1) // [the cat in the hat]

	var string2 = strings.Split("golang#python#javascript", "#")
	fmt.Println(string2) // [golang python javascript]

	var string3 = strings.Split("batman", "")
	fmt.Println(string3) // [b a t m a n]

	// menggabungkan string dengan pemisah tertentu
	var data = []string{"banana", "papaya", "apple"}
	var str1 = strings.Join(data, "-")
	fmt.Println(str1) // banana-papaya-apple

	// mengubah string menjadi huruf kecil
	var str2 = strings.ToLower("GOLANG")
	fmt.Println(str2) // golang

	// mengubah string menjadi huruf besar
	var str3 = strings.ToUpper("golang")
	fmt.Println(str3) // GOLANG
}