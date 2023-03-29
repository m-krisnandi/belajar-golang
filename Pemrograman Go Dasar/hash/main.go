package main

import "crypto/sha1"
import "fmt"
import "time"

func main() {
	var text = "this is text to be hashed"
	fmt.Printf("original : %s\n\n", text)

	var sha = sha1.New()
	sha.Write([]byte(text))

	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("Tidak menggunakan salt : %x", encrypted)

	fmt.Println(encryptedString)
	// 90a7eb3b30949cc4ec864c5a102b603477fb408a

	// using salt
	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("Hash 1 : %s\n\n", hashed1)

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("Hash 2 : %s\n\n", hashed2)

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("Hash 3 : %s\n\n", hashed3)

	_, _, _ = salt1, salt2, salt3
}

// salt untuk mengacak data hasil hash
func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))

	var encrypted = sha.Sum(nil)
	return fmt.Sprintf("%x", encrypted), salt
}