package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"math/rand"
	"time"
	"log"
)

const totalFile = 3000
const contentLength = 5000

// var tempPath = filepath.Join(os.Getenv("TEMP"), "currency-pattern-pipeline-temp")

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		fileName := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomString(contentLength)
		err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", fileName)
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "files created")
		}
	}

	log.Printf("Created %d files", totalFile)
}

// func main() {
// 	log.Println("start")
// 	start := time.Now()

// 	generateFiles()

// 	duration := time.Since(start)
// 	log.Println("done in", duration.Seconds(), "seconds")
// }