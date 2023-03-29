package main

import "fmt"
import "strconv"

func main() {
	// konversi string ke integer
	var str = "1234"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
	}

	// konversi integer ke string
	var num1 = 124
	var str1 = strconv.Itoa(num1)

	fmt.Println(str1)

	// konversi string ke integer dengan base 10 dan 2 (biner)
	// dengan tipe data int64 dan int8
	// var str2 = "124"
	var str2 = "1010"
	// var num2, err2 = strconv.ParseInt(str2, 10, 64)
	var num2, err2 = strconv.ParseInt(str2, 2, 8)

	if err2 == nil {
		fmt.Println(num2)
	}

	// konversi int64 ke string dengan base 8
	var num3 = int64(24)
	var str3 = strconv.FormatInt(num3, 8)

	fmt.Println(str3)

	// konversi string ke tipe data float32
	var str4 = "24.12"
	var num4, err3 = strconv.ParseFloat(str4, 32)

	if err3 == nil {
		fmt.Println(num4)
	}

	// konversi float64 ke string dengan 6 digit di belakang koma
	var num5 = float64(24.12)
	var str5 = strconv.FormatFloat(num5, 'f', 6, 64)

	fmt.Println(str5)

	// konversi string ke boolean
	var str6 = "true"
	var bul, err4 = strconv.ParseBool(str6)

	if err4 == nil {
		fmt.Println(bul)
	}

	// konversi boolean ke string
	var bul1 = true
	var str7 = strconv.FormatBool(bul1)

	fmt.Println(str7)

	// konversi menggunakan teknik casting
	var a float64 = float64(24)
	fmt.Println(a)

	var b int32 = int32(24.00)
	fmt.Println(b)

	// casting string ke byte
	var text = "halo"
	var bytes = []byte(text)

	fmt.Printf("%d %d %d %d \n", bytes[0], bytes[1], bytes[2], bytes[3])

	// casting byte ke string
	var bytes2 = []byte{104, 97, 108, 111}
	var text2 = string(bytes2)

	fmt.Println(text2)

	// casting string ke int64
	var c int64 = int64('h')
	fmt.Println(c)

	// casting int ke string
	var d string = string(104)
	fmt.Println(d)

	// type assertion pada interface kosong
	var data = map[string]interface{} {
		"name": "John Wick",
		"grade": 2,
		"height": 156.5,
		"isMale": true,
		"hobbies": []string{"reading", "swimming", "coding"},
	}

	// cara 1
	// casting interface kosong ke tipe data yang sesuai
	// dengan menggunakan switch case
	// untuk mengetahui tipe data dari interface kosong

	for _, val := range data{
		switch val.(type) {
			case string:
				fmt.Println(val.(string))
			case int:
				fmt.Println(val.(int))
			case float64:
				fmt.Println(val.(float64))
			case bool:
				fmt.Println(val.(bool))
			case []string:
				fmt.Println(val.([]string))
			default:
				fmt.Println("unknown type")
		}
	}

	// cara 2

	fmt.Println(data["name"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	
}