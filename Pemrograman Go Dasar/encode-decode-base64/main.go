package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "hello world"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString) // aGVsbG8gd29ybGQ=

	var decodedBytes, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedBytes)
	fmt.Println(decodedString) // hello world

	// penerapan fungsi encode dan decode
	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString2 = string(encoded)
	fmt.Println(encodedString2) // aGVsbG8gd29ybGQ=

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)

	if err != nil {
		fmt.Println(err.Error())
	}

	var decodedString2 = string(decoded)
	fmt.Println(decodedString2) // hello world

	// penerapan fungsi encode dan decode dengan url
	var dataUrl = "https://mengcoding.com/"

	var encodedUrl = base64.URLEncoding.EncodeToString([]byte(dataUrl))
	fmt.Println(encodedUrl) // aHR0cHM6Ly9tZW5nY29kaW5nLmNvbS8=

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedUrl)
	var decodedUrl = string(decodedByte)
	fmt.Println(decodedUrl) // https://mengcoding.com/
}