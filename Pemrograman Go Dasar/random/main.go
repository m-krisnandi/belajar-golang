package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// random tipe data integer
	fmt.Println("random ke-1", randomizer.Int())
	fmt.Println("random ke-2", randomizer.Int())
	fmt.Println("random ke-3", randomizer.Int())

	// random tipe data string
	randomStr := randomString(10)
	fmt.Println("random string:", randomStr)
}

// random tipe data string
var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}