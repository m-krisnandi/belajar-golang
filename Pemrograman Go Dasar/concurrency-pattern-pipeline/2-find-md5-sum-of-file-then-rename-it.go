package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// var tempPath = filepath.Join(os.Getenv("TEMP"), "currency-pattern-pipeline-temp")

func proceed() {
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {

		// if there is an error, return immediately
		if err != nil {
			return err
		}

		// if it is a sub directory, return immediately
		if info.IsDir() {
			return nil
		}

		counterTotal++

		// read the file content
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// calculate the md5 sum
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		// rename the file
		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destinationPath)

		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("ERROR", err.Error())
	}

	log.Printf("Total %d files, renamed %d files", counterTotal, counterRenamed)
}

// func main() {
// 	log.Println("start")
// 	start := time.Now()

// 	proceed()

// 	duration := time.Since(start)
// 	log.Println("done in", duration.Seconds(), "seconds")
// }