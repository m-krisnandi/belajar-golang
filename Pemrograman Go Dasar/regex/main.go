package main

import "fmt"
import "regexp"

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	// mencari 2 kata pertama
	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1) // []string{"banana", "burger"}

	// mencari semua kata
	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2) // []string{"banana", "burger", "soup"}

	// mendeteksi apakah memenuhi pola regex
	var text1 = "banana burger soup"
	var regex1, _ = regexp.Compile(`[a-z]+`)

	var isMatch = regex1.MatchString(text1)
	fmt.Println(isMatch) // true

	// mencari string yang memenuhi pola regex
	var text2 = "banana burger soup"
	var regex2, _ = regexp.Compile(`[a-z]+`)

	var res3 = regex2.FindString(text2)
	fmt.Println(res3) // banana

	// mencari index string yang memenuhi pola regex
	var text3 = "banana burger soup"
	var regex3, _ = regexp.Compile(`[a-z]+`)

	var idx = regex3.FindStringIndex(text3)
	fmt.Println(idx) // [0 6]

	// mencari semua index string yang memenuhi pola regex
	var text4 = "banana burger soup"
	var regex4, _ = regexp.Compile(`[a-z]+`)

	var idx2 = regex4.FindAllStringIndex(text4, -1)
	fmt.Println(idx2) // [[0 6] [7 12] [13 17]]

	var str2 = text4[0:6]
	fmt.Println(str2) // banana

	// replace string yang memenuhi pola regex
	var text5 = "banana burger soup"
	var regex5, _ = regexp.Compile(`[a-z]+`)

	var res4 = regex5.ReplaceAllString(text5, "potato")
	fmt.Println(res4) // potato potato potato

	// replace string yang memenuhi pola regex dengan fungsi
	var text6 = "banana burger soup"
	var regex6, _ = regexp.Compile(`[a-z]+`)

	var res5 = regex6.ReplaceAllStringFunc(text6, func(each string) string {
		if each == "banana" {
			return "potato"
		}
		return each
	})

	fmt.Println(res5) // potato burger soup

	// split string yang memenuhi pola regex
	var text7 = "banana burger soup"
	var regex7, _ = regexp.Compile(`[a-b]+`) // mencari huruf a sampai b

	var res6 = regex7.Split(text7, -1)
	fmt.Printf("%#v \n", res6) // []string{"", "n", "na", " ", "urg", "r ", "oup"}
}